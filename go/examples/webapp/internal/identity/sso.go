package identity

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// ssoCookie carries one sign-in request's state and PKCE verifier until the issuer returns. Lax,
// because the issuer's redirect back is a top-level navigation from another site.
const ssoCookie = "traust_sso"

// discoveryLimit bounds what the app reads of a provider's documents.
const discoveryLimit = 1 << 20

// httpTimeout bounds the app's asks of the issuer.
const httpTimeout = 10 * time.Second

// SSO is the issuer the app signs people in with: authorization code with PKCE and state. Core
// verifies the token endpoint's access_token (its own login keeps that one, never the id_token),
// so that is the one the app keeps too.
type SSO struct {
	Issuer       string
	ClientID     string
	ClientSecret string
	RedirectURL  string

	client  *http.Client
	mu      sync.Mutex
	found   endpoints
	fetched bool
}

// endpoints is what the app needs of the issuer's discovery document.
type endpoints struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
}

// NewSSO builds the sign-in provider; no issuer configured is no SSO (nil), a half-configured one
// is an error.
func NewSSO(issuer, clientID, clientSecret, redirectURL string) (*SSO, error) {
	if issuer == "" {
		if clientID != "" || clientSecret != "" || redirectURL != "" {
			return nil, errors.New("identity: SSO client settings need an issuer")
		}
		return nil, nil
	}
	if clientID == "" || redirectURL == "" {
		return nil, errors.New("identity: an SSO issuer needs a client id and a redirect URL")
	}
	if to, err := url.Parse(redirectURL); err != nil || !to.IsAbs() || to.Host == "" {
		return nil, fmt.Errorf("identity: the SSO redirect URL %q is not an absolute URL", redirectURL)
	}
	return &SSO{
		Issuer: issuer, ClientID: clientID, ClientSecret: clientSecret, RedirectURL: redirectURL,
		client: &http.Client{Timeout: httpTimeout},
	}, nil
}

// Link is what the sign-in page offers SSO from; nothing when no issuer is set.
func (s *SSO) Link() string {
	if s == nil {
		return ""
	}
	return "/sso"
}

// start redirects the person to the issuer's authorization endpoint, one state and one PKCE
// challenge behind them in a cookie.
func (s *SSO) start(w http.ResponseWriter, r *http.Request) error {
	found, err := s.endpoints(r.Context())
	if err != nil {
		return err
	}
	state, verifier, err := mint()
	if err != nil {
		return err
	}
	sum := sha256.Sum256([]byte(verifier))
	http.SetCookie(w, &http.Cookie{
		Name: ssoCookie, Value: state + "." + verifier, Path: "/sso", MaxAge: 600,
		HttpOnly: true, Secure: r.TLS != nil, SameSite: http.SameSiteLaxMode,
	})
	ask := url.Values{
		"response_type":         {"code"},
		"client_id":             {s.ClientID},
		"redirect_uri":          {s.RedirectURL},
		"scope":                 {"openid email"}, // core's identity claim is email by default
		"state":                 {state},
		"code_challenge":        {base64.RawURLEncoding.EncodeToString(sum[:])},
		"code_challenge_method": {"S256"},
	}
	http.Redirect(w, r, found.AuthorizationEndpoint+"?"+ask.Encode(), http.StatusSeeOther)
	return nil
}

// exchange trades the authorization code for the access_token core verifies, proving the request's
// verifier. The state cookie is only spent once.
func (s *SSO) exchange(ctx context.Context, code, verifier string) (string, error) {
	found, err := s.endpoints(ctx)
	if err != nil {
		return "", err
	}
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {s.RedirectURL},
		"client_id":     {s.ClientID},
		"code_verifier": {verifier},
	}
	if s.ClientSecret != "" {
		form.Set("client_secret", s.ClientSecret)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, found.TokenEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("asking %s for a token: %w", s.Issuer, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("asking %s for a token: %w", s.Issuer, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, discoveryLimit))
	if err != nil {
		return "", fmt.Errorf("reading %s's token answer: %w", s.Issuer, err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s refused the code: %s", s.Issuer, oauthError(body))
	}
	var grant struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &grant); err != nil || grant.AccessToken == "" {
		return "", fmt.Errorf("%s's token answer has no access_token", s.Issuer)
	}
	return grant.AccessToken, nil
}

// answered confirms the callback answers the request its cookie names: same state, and a code.
func answered(r *http.Request) (verifier, code string, ok bool) {
	cookie, err := r.Cookie(ssoCookie)
	if err != nil {
		return "", "", false
	}
	state, verifier, found := strings.Cut(cookie.Value, ".")
	if !found {
		return "", "", false
	}
	asked := r.URL.Query()
	if asked.Get("state") == "" || subtle.ConstantTimeCompare([]byte(state), []byte(asked.Get("state"))) != 1 {
		return "", "", false
	}
	if code = asked.Get("code"); code == "" {
		return "", "", false
	}
	return verifier, code, true
}

// endpoints discovers the issuer's endpoints once, then keeps them.
func (s *SSO) endpoints(ctx context.Context) (endpoints, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.fetched {
		return s.found, nil
	}
	doc := strings.TrimSuffix(s.Issuer, "/") + "/.well-known/openid-configuration"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, doc, nil)
	if err != nil {
		return endpoints{}, fmt.Errorf("reading %s's discovery document: %w", s.Issuer, err)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return endpoints{}, fmt.Errorf("reading %s's discovery document: %w", s.Issuer, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return endpoints{}, fmt.Errorf("%s answered %d for its discovery document", s.Issuer, resp.StatusCode)
	}
	var found endpoints
	if err := json.NewDecoder(io.LimitReader(resp.Body, discoveryLimit)).Decode(&found); err != nil {
		return endpoints{}, fmt.Errorf("%s's discovery document is not JSON: %w", s.Issuer, err)
	}
	if found.AuthorizationEndpoint == "" || found.TokenEndpoint == "" {
		return endpoints{}, fmt.Errorf("%s's discovery document names no endpoints", s.Issuer)
	}
	s.found, s.fetched = found, true
	return found, nil
}

// mint mints one sign-in request's state and PKCE verifier.
func mint() (state, verifier string, err error) {
	raw := make([]byte, 64)
	if _, err := rand.Read(raw); err != nil {
		return "", "", fmt.Errorf("minting a sign-in request: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(raw[:32]), base64.RawURLEncoding.EncodeToString(raw[32:]), nil
}

// oauthError words a token refusal without echoing anything secret-sized.
func oauthError(body []byte) string {
	var refused struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
	}
	if err := json.Unmarshal(body, &refused); err == nil && refused.Error != "" {
		if refused.ErrorDescription != "" {
			return refused.ErrorDescription
		}
		return refused.Error
	}
	return http.StatusText(http.StatusBadRequest)
}

// spent drops the sign-in request's cookie.
func spent(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{Name: ssoCookie, Path: "/sso", MaxAge: -1})
}
