package traust

// Bindings over the generated client: bearer-token auth, bounded retries on transient
// failures, typed problem errors and cursor paging. The generated calls — root's and every
// domain package's — build the `(T, error)` calls and `…All` iterators on top of these.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"net/http"
	"os"
	"strings"
	"time"
)

const healthzPath = "/v1/healthz"

// TokenSource supplies the app's own bearer token for each request: its machine identity, for
// the calls it makes as itself (background jobs), never for a person's.
type TokenSource interface {
	BearerToken(ctx context.Context) (string, error)
}

// FileTokenSource reads the token file on every request, so projected
// service-account tokens rotate (hourly, by default) without a restart.
type FileTokenSource struct {
	Path string
}

func (s FileTokenSource) BearerToken(_ context.Context) (string, error) {
	raw, err := os.ReadFile(s.Path)
	if err != nil {
		return "", fmt.Errorf("reading token file %s: %w", s.Path, err)
	}
	token := strings.TrimSpace(string(raw))
	if token == "" {
		return "", fmt.Errorf("token file %s is empty", s.Path)
	}
	return token, nil
}

// StaticToken is a fixed bearer token, for scripts, local development and tests; services
// use FileTokenSource so rotated tokens are picked up.
type StaticToken string

func (t StaticToken) BearerToken(_ context.Context) (string, error) {
	if strings.TrimSpace(string(t)) == "" {
		return "", errors.New("traust: static token is empty")
	}
	return string(t), nil
}

type onBehalfOfKey struct{}

// OnBehalfOf sends the calls made with ctx under a person's bearer token instead of the client's,
// so core authorizes and attributes them to that person: a BFF acting for its signed-in user.
func OnBehalfOf(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, onBehalfOfKey{}, token)
}

func bearer(ctx context.Context, tokens TokenSource) (string, error) {
	token, person := ctx.Value(onBehalfOfKey{}).(string)
	if !person {
		if tokens == nil {
			return "", ErrNoIdentity
		}
		return tokens.BearerToken(ctx)
	}
	if strings.TrimSpace(token) == "" {
		return "", errors.New("OnBehalfOf token is empty")
	}
	return token, nil
}

// Config configures a Client. BaseURL is required. Tokens is the app's own identity, set only
// by an app that makes calls as itself; every other call names its person with OnBehalfOf. The
// rest have defaults.
type Config struct {
	BaseURL string
	Tokens  TokenSource // optional

	HTTPClient *http.Client  // optional; default http.DefaultClient wrapped with retries
	Retries    int           // attempts beyond the first; default 2
	Backoff    time.Duration // between retries; default 250ms, doubling
}

// Client is traust-core's client: the transport — auth, retries, typed errors — the domain
// clients call through, with the health and identity reads.
type Client struct {
	raw *ClientWithResponses
}

// Raw is the generated client under the typed calls: raw HTTP responses, for the domain
// packages and the rare caller that needs a response's headers.
func (c *Client) Raw() *ClientWithResponses { return c.raw }

// NewClient validates cfg and connects to traust-core.
func NewClient(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, errors.New("traust: BaseURL is required")
	}
	if cfg.Retries < 0 {
		return nil, fmt.Errorf("traust: Retries must not be negative, got %d", cfg.Retries)
	}
	if cfg.Retries == 0 {
		cfg.Retries = 2
	}
	if cfg.Backoff == 0 {
		cfg.Backoff = 250 * time.Millisecond
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	inner := httpClient.Transport
	if inner == nil {
		inner = http.DefaultTransport
	}
	// Copy the caller's client and replace only the transport, so
	// CheckRedirect, Jar and the rest survive; building a fresh client here
	// used to drop them (a caller's ErrUseLastResponse policy was silently
	// ignored and redirects followed).
	wrapped := *httpClient
	wrapped.Transport = &retryTransport{base: inner, retries: cfg.Retries, backoff: cfg.Backoff}
	httpClient = &wrapped
	editor := func(ctx context.Context, req *http.Request) error {
		token, err := bearer(ctx, cfg.Tokens)
		if errors.Is(err, ErrNoIdentity) && strings.HasSuffix(req.URL.Path, healthzPath) {
			return nil // core's liveness check takes no caller
		}
		if err != nil {
			return fmt.Errorf("traust: bearer token: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
	generated, err := NewClientWithResponses(cfg.BaseURL, WithRequestEditorFn(editor), WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("traust: building client: %w", err)
	}
	return &Client{raw: generated}, nil
}

// CheckResponse returns nil for a 2xx response, otherwise its *ProblemError: the RFC 7807
// document when the server sent one, else the raw body.
func CheckResponse(statusCode int, body []byte, contentType string) error {
	if statusCode >= 200 && statusCode < 300 {
		return nil
	}
	out := &ProblemError{StatusCode: statusCode, Body: string(body)}
	if strings.Contains(contentType, "application/problem+json") || bytes.HasPrefix(bytes.TrimSpace(body), []byte("{")) {
		var problem Problem
		if err := json.Unmarshal(body, &problem); err == nil && problem.Title != "" {
			out.Problem = &problem
		}
	}
	return out
}

// Refusal returns nil for a 2xx response, otherwise its *ProblemError with the body read and
// closed. The generated calls check every raw response with it before decoding.
func Refusal(response *http.Response) error {
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	}
	defer func() { _ = response.Body.Close() }()
	body, err := io.ReadAll(response.Body)
	refusal := CheckResponse(response.StatusCode, body, response.Header.Get("Content-Type"))
	if err != nil {
		return errors.Join(refusal, fmt.Errorf("traust-core: reading refusal body: %w", err))
	}
	return refusal
}

// Decode returns a generated response's decoded body, or its *ProblemError when core refused.
// The generated calls use it.
func Decode[T any](response *http.Response, body []byte, value *T) (*T, error) {
	if err := CheckResponse(response.StatusCode, body, response.Header.Get("Content-Type")); err != nil {
		return nil, err
	}
	if value == nil || bytes.Equal(bytes.TrimSpace(body), []byte("null")) {
		return nil, fmt.Errorf("traust-core: HTTP %d without a JSON body", response.StatusCode)
	}
	return value, nil
}

// retryTransport retries idempotent requests (GET/HEAD) on 429 and
// 502/503/504 and on transport errors, at most retries extra attempts with
// doubling backoff. Other methods and statuses pass through untouched.
type retryTransport struct {
	base    http.RoundTripper
	retries int
	backoff time.Duration
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	retryable := req.Method == http.MethodGet || req.Method == http.MethodHead
	var resp *http.Response
	var err error
	for attempt := 0; ; attempt++ {
		// The body of a GET/HEAD is nil or already consumed safely; replay
		// by reference is fine for idempotent requests.
		resp, err = t.base.RoundTrip(req)
		if !retryable || attempt >= t.retries {
			break
		}
		if err == nil && resp.StatusCode != http.StatusTooManyRequests &&
			resp.StatusCode != http.StatusBadGateway &&
			resp.StatusCode != http.StatusServiceUnavailable &&
			resp.StatusCode != http.StatusGatewayTimeout {
			break
		}
		if resp != nil {
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()
			resp = nil
		}
		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(t.backoff << attempt):
		}
	}
	return resp, err
}

// Pages iterates a cursor-paged listing, one fetch per page, until core returns no next
// cursor. The generated `…All` iterators are built on it.
func Pages[T any](fetch func(cursor *string) ([]T, *string, error), start *string) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		cursor := start
		for {
			items, next, err := fetch(cursor)
			if err != nil {
				var zero T
				yield(zero, err)
				return
			}
			for _, item := range items {
				if !yield(item, nil) {
					return
				}
			}
			if next == nil || *next == "" {
				return
			}
			cursor = next
		}
	}
}
