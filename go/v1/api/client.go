package api

// Bindings over the generated client: bearer-token auth from a rotating
// file, bounded retries on transient failures, typed problem errors, and
// cursor pagination. The generated client.gen.go stays untouched; this file
// is the hand-written surface consumers build against.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

// ClientConfig configures the bound client. BaseURL and Tokens are
// required; the rest have defaults.
type ClientConfig struct {
	BaseURL string
	Tokens  TokenSource

	HTTPClient *http.Client     // optional; default http.DefaultClient wrapped with retries
	Retries    int              // attempts beyond the first; default 2
	Backoff    time.Duration    // between retries; default 250ms, doubling
	Now        func() time.Time // clock for tests; default time.Now
}

// BoundClient is the generated client with auth, retries and typed errors bound.
type BoundClient struct {
	*ClientWithResponses
}

// NewBoundClient validates the configuration and wires the generated client.
func NewBoundClient(cfg ClientConfig) (*BoundClient, error) {
	if cfg.BaseURL == "" {
		return nil, errors.New("api: BaseURL is required")
	}
	if cfg.Tokens == nil {
		return nil, errors.New("api: Tokens is required")
	}
	if cfg.Retries < 0 {
		return nil, fmt.Errorf("api: Retries must not be negative, got %d", cfg.Retries)
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
			return fmt.Errorf("api: bearer token: %w", err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	}
	generated, err := NewClientWithResponses(cfg.BaseURL, WithRequestEditorFn(editor), WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("api: building client: %w", err)
	}
	return &BoundClient{ClientWithResponses: generated}, nil
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

// problemFromResponse builds the error for a non-2xx generated response.
func problemFromResponse(statusCode int, body []byte, contentType string) *ProblemError {
	out := &ProblemError{StatusCode: statusCode, Body: string(body)}
	if strings.Contains(contentType, "application/problem+json") || bytes.HasPrefix(bytes.TrimSpace(body), []byte("{")) {
		var problem Problem
		if err := json.Unmarshal(body, &problem); err == nil && problem.Title != "" {
			out.Problem = &problem
		}
	}
	return out
}

// checkResponse turns a failed generated response into a *ProblemError.
// The generated types expose StatusCode() as a method and Body as a field.
func checkResponse(statusCode int, body []byte, contentType string) error {
	if statusCode >= 200 && statusCode < 300 {
		return nil
	}
	return problemFromResponse(statusCode, body, contentType)
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
		case <-time.After(t.backoff << uint(attempt)):
		}
	}
	return resp, err
}

// EachProduct walks the product listing page by page, following the cursor,
// calling fn for every product; a non-nil error stops the walk. One request
// page size is limit (1..1000; core clamps).
func (c *BoundClient) EachProduct(ctx context.Context, limit int, fn func(Product) error) error {
	var cursor *string
	for {
		resp, err := c.ListProductsV1ProductsGetWithResponse(ctx, &ListProductsV1ProductsGetParams{
			Limit:  &limit,
			Cursor: cursor,
		})
		if err != nil {
			return err
		}
		if err := checkResponse(resp.StatusCode(), resp.Body, resp.HTTPResponse.Header.Get("Content-Type")); err != nil {
			return err
		}
		if resp.JSON200 == nil {
			return problemFromResponse(resp.StatusCode(), resp.Body, resp.HTTPResponse.Header.Get("Content-Type"))
		}
		for _, item := range resp.JSON200.Items {
			if err := fn(item); err != nil {
				return err
			}
		}
		if resp.JSON200.Next == nil || *resp.JSON200.Next == "" {
			return nil
		}
		cursor = resp.JSON200.Next
	}
}

// EachRepository walks the repository listing the same way.
func (c *BoundClient) EachRepository(ctx context.Context, limit int, fn func(Repository) error) error {
	var cursor *string
	for {
		resp, err := c.ListRepositoriesV1RepositoriesGetWithResponse(ctx, &ListRepositoriesV1RepositoriesGetParams{
			Limit:  &limit,
			Cursor: cursor,
		})
		if err != nil {
			return err
		}
		if err := checkResponse(resp.StatusCode(), resp.Body, resp.HTTPResponse.Header.Get("Content-Type")); err != nil {
			return err
		}
		if resp.JSON200 == nil {
			return problemFromResponse(resp.StatusCode(), resp.Body, resp.HTTPResponse.Header.Get("Content-Type"))
		}
		for _, item := range resp.JSON200.Items {
			if err := fn(item); err != nil {
				return err
			}
		}
		if resp.JSON200.Next == nil || *resp.JSON200.Next == "" {
			return nil
		}
		cursor = resp.JSON200.Next
	}
}
