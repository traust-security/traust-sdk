package traust

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	// Transport tuning for the token endpoint, applied when the caller
	// supplies no client or clock of its own.
	tokenHTTPTimeout = 30 * time.Second
	tokenBodyLimit   = 1 << 20
	tokenRefreshLead = 10 * time.Second
)

// ClientCredentialsTokenSource is a TokenSource for the OAuth 2.0
// client-credentials grant: it fetches a bearer token from the issuer's token
// endpoint, caches it until shortly before it expires, then refreshes.
// Build it with NewClientCredentialsTokenSource; it is safe for concurrent use.
type ClientCredentialsTokenSource struct {
	TokenURL     string
	ClientID     string
	ClientSecret string
	Scopes       []string

	HTTPClient *http.Client     // optional; default with a 30s timeout
	Now        func() time.Time // clock for tests; default time.Now

	// The cache lives behind a pointer so the constructor can take the
	// config by value without copying a lock.
	cache *tokenCache
}

// NewClientCredentialsTokenSource validates src and returns it ready for use.
// TokenURL must be https (http only on a loopback host, for local issuers).
func NewClientCredentialsTokenSource(src ClientCredentialsTokenSource) (*ClientCredentialsTokenSource, error) {
	switch {
	case src.TokenURL == "":
		return nil, errors.New("traust: TokenURL is required")
	case src.ClientID == "":
		return nil, errors.New("traust: ClientID is required")
	case src.ClientSecret == "":
		return nil, errors.New("traust: ClientSecret is required")
	}
	endpoint, err := url.Parse(src.TokenURL)
	if err != nil || endpoint.Host == "" {
		return nil, fmt.Errorf("traust: TokenURL %q is not a valid URL", src.TokenURL)
	}
	if endpoint.Scheme != "https" && !(endpoint.Scheme == "http" && isLoopbackHost(endpoint.Hostname())) {
		return nil, fmt.Errorf("traust: TokenURL must be https, or http on a loopback host: %s", src.TokenURL)
	}
	src.cache = &tokenCache{}
	return &src, nil
}

type tokenCache struct {
	mu     sync.Mutex
	token  string
	expiry time.Time
}

// TokenFetchError reports a failed client-credentials grant: the reason it
// failed, the issuer's HTTP status on a refusal, and the underlying cause.
// It never carries the client secret or the response body.
type TokenFetchError struct {
	TokenURL   string
	Reason     string
	StatusCode int
	Err        error
}

func (e *TokenFetchError) Error() string {
	if e.StatusCode != 0 {
		return fmt.Sprintf("traust: client-credentials grant to %s: %s: HTTP %d", e.TokenURL, e.Reason, e.StatusCode)
	}
	if e.Err != nil {
		return fmt.Sprintf("traust: client-credentials grant to %s: %s: %v", e.TokenURL, e.Reason, e.Err)
	}
	return fmt.Sprintf("traust: client-credentials grant to %s: %s", e.TokenURL, e.Reason)
}

func (e *TokenFetchError) Unwrap() error { return e.Err }

// BearerToken returns the cached access token, fetching a fresh one when it is
// missing or past its refresh point. Concurrent callers share one fetch, and a
// failed refresh never returns the stale token.
func (s *ClientCredentialsTokenSource) BearerToken(ctx context.Context) (string, error) {
	if s.cache == nil {
		return "", errors.New("traust: ClientCredentialsTokenSource must be built with NewClientCredentialsTokenSource")
	}
	s.cache.mu.Lock()
	defer s.cache.mu.Unlock()
	now := s.Now
	if now == nil {
		now = time.Now
	}
	if s.cache.token != "" && now().Before(s.cache.expiry) {
		return s.cache.token, nil
	}
	token, expiry, err := s.fetch(ctx, now)
	if err != nil {
		return "", err
	}
	s.cache.token, s.cache.expiry = token, expiry
	return token, nil
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func (s *ClientCredentialsTokenSource) fetch(ctx context.Context, now func() time.Time) (string, time.Time, error) {
	form := url.Values{"grant_type": {"client_credentials"}}
	if len(s.Scopes) > 0 {
		form.Set("scope", strings.Join(s.Scopes, " "))
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "building the token request", Err: err}
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	// RFC 6749 §2.3.1: the credentials are form-urlencoded before Basic auth.
	req.SetBasicAuth(url.QueryEscape(s.ClientID), url.QueryEscape(s.ClientSecret))
	client := s.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: tokenHTTPTimeout}
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "posting the grant", Err: err}
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(io.LimitReader(resp.Body, tokenBodyLimit))
	if err != nil {
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "reading the token response", Err: err}
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "issuer refused the grant", StatusCode: resp.StatusCode}
	}
	var tok tokenResponse
	if err := json.Unmarshal(body, &tok); err != nil {
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "malformed token response", Err: err}
	}
	switch {
	case tok.AccessToken == "":
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "token response without access_token"}
	case !strings.EqualFold(strings.TrimSpace(tok.TokenType), "bearer"):
		// The server's token_type value stays out of the message: a hostile
		// issuer could echo the Basic-auth credentials in it.
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "token response with a non-bearer token_type"}
	case tok.ExpiresIn <= 0:
		// Fail closed: without expires_in the token cannot be refreshed before it expires.
		return "", time.Time{}, &TokenFetchError{TokenURL: s.TokenURL, Reason: "token response without expires_in"}
	}
	lifetime := time.Duration(tok.ExpiresIn) * time.Second
	lead := tokenRefreshLead
	if lifetime <= tokenRefreshLead {
		lead = lifetime / 2
	}
	return tok.AccessToken, now().Add(lifetime - lead), nil
}

func isLoopbackHost(host string) bool {
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}
