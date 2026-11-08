package traust

// Bindings over the generated client: bearer-token auth, bounded retries on transient
// failures, typed problem errors and cursor paging. calls.gen.go builds the `(T, error)` calls
// and `…All` iterators on top of these.

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

// TokenSource supplies the bearer token for each request.
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

// Config configures a Client. BaseURL and Tokens are required; the rest have defaults.
type Config struct {
	BaseURL string
	Tokens  TokenSource

	HTTPClient *http.Client     // optional; default http.DefaultClient wrapped with retries
	Retries    int              // attempts beyond the first; default 2
	Backoff    time.Duration    // between retries; default 250ms, doubling
	Now        func() time.Time // clock for tests; default time.Now
}

// Client is traust-core's client: every operation as a `(T, error)` call and every paged listing
// as an `…All` iterator, with auth, retries and typed errors.
type Client struct {
	*ClientWithResponses
}

// NewClient validates cfg and connects to traust-core.
func NewClient(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, errors.New("traust: BaseURL is required")
	}
	if cfg.Tokens == nil {
		return nil, errors.New("traust: Tokens is required")
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
	now := cfg.Now
	if now == nil {
		now = time.Now
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
	wrapped.Transport = &retryTransport{base: inner, retries: cfg.Retries, backoff: cfg.Backoff, now: now}
	httpClient = &wrapped
	editor := func(ctx context.Context, req *http.Request) error {
		token, err := cfg.Tokens.BearerToken(ctx)
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
	return &Client{ClientWithResponses: generated}, nil
}

// ProblemError reports a non-retryable failed response. Problem carries the
// RFC 7807 document when the server sent one; otherwise Body holds the raw
// response.
type ProblemError struct {
	StatusCode int
	Problem    *Problem
	Body       string
}

func (e *ProblemError) Error() string {
	if e.Problem != nil {
		detail := e.Problem.Detail
		if detail == "" {
			detail = e.Problem.Title
		}
		request := ""
		if e.Problem.RequestId != "" {
			request = " (request " + e.Problem.RequestId + ")"
		}
		return fmt.Sprintf("traust-core: %d %s: %s%s", e.StatusCode, e.Problem.Title, detail, request)
	}
	body := e.Body
	if len(body) > 200 {
		body = body[:200]
	}
	return fmt.Sprintf("traust-core: %d: %s", e.StatusCode, body)
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

func checkRefusal(response *http.Response) error {
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
	now     func() time.Time
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

// pages iterates a cursor-paged listing, one fetch per page, until core returns no next cursor.
func pages[T any](fetch func(cursor *string) ([]T, *string, error)) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var cursor *string
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
