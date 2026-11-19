// Package transport implements the HTTP transport for the ledger service.
//
// Wire contract:
//
//	POST /v1/ledger/events                        — human lane (countersign, severity)
//	POST /v1/ledger/layer/submit?layer_id=        — machine lane batch submit
//	POST /v1/ledger/layer/resolve?layer_id=       — resolve needs_review item
//	POST /v1/ledger/fingerprint                   — stamp finding fingerprints
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/traust-security/traust-sdk/go/v1/ledger/internal/route"
)

const eventsPath = "/v1/ledger/events"

type eventEnvelope struct {
	Kind  string          `json:"kind"`
	Event json.RawMessage `json:"event"`
}

// StatusError is returned when the ledger service responds with a non-2xx status.
type StatusError struct {
	StatusCode int
	Body       []byte
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("ingest: http status %d: %s", e.StatusCode, string(e.Body))
}

// ErrResponseTooLarge is returned when a response body exceeds the limit
// the provider was configured with (WithResponseLimit): the read stops at
// the limit, the body is never buffered whole.
var ErrResponseTooLarge = errors.New("ledger response exceeds the configured limit")

// HTTPProvider implements the ledger service HTTP transport.
type HTTPProvider struct {
	baseURL string
	client  *http.Client
	auth    func(*http.Request)
	limit   int64
}

// Option configures the HTTP provider.
type Option func(*HTTPProvider)

// WithHTTPClient sets a custom *http.Client.
func WithHTTPClient(c *http.Client) Option {
	return func(p *HTTPProvider) { p.client = c }
}

// WithBearerToken sets a static bearer token for all requests.
func WithBearerToken(token string) Option {
	return func(p *HTTPProvider) {
		p.auth = func(r *http.Request) {
			r.Header.Set("Authorization", "Bearer "+token)
		}
	}
}

// WithAuthFunc sets a custom auth function called before each request.
func WithAuthFunc(fn func(*http.Request)) Option {
	return func(p *HTTPProvider) { p.auth = fn }
}

// WithResponseLimit bounds what one response body may buffer: a body over
// the limit fails with ErrResponseTooLarge instead of being read whole. A
// limit of zero or less keeps reads unbounded.
func WithResponseLimit(limit int64) Option {
	return func(p *HTTPProvider) { p.limit = limit }
}

// New creates an HTTPProvider targeting the given ledger service base URL.
func New(baseURL string, opts ...Option) *HTTPProvider {
	p := &HTTPProvider{
		baseURL: strings.TrimRight(baseURL, "/"),
		client:  http.DefaultClient,
	}
	for _, o := range opts {
		o(p)
	}
	return p
}

// Ingest wraps the payload in the human-lane event envelope and POSTs to
// POST /v1/ledger/events.
func (p *HTTPProvider) Ingest(ctx context.Context, kind string, payload []byte) ([]byte, error) {
	body, err := json.Marshal(eventEnvelope{Kind: kind, Event: payload})
	if err != nil {
		return nil, fmt.Errorf("ingest: marshal envelope: %w", err)
	}
	return p.post(ctx, eventsPath, body)
}

// Submit POSTs a batch-submit body to POST /v1/ledger/layer/submit?layer_id=.
func (p *HTTPProvider) Submit(ctx context.Context, layerID string, payload []byte) ([]byte, error) {
	path := route.Layer(route.Submit, layerID, nil)
	return p.post(ctx, path, payload)
}

// Post sends a raw JSON body to an absolute API path (e.g. /v1/ledger/fingerprint).
func (p *HTTPProvider) Post(ctx context.Context, path string, payload []byte) ([]byte, error) {
	return p.post(ctx, path, payload)
}

// Query issues an HTTP request to the ledger service read API.
func (p *HTTPProvider) Query(ctx context.Context, method, path string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, p.baseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("query: build request: %w", err)
	}
	if p.auth != nil {
		p.auth(req)
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	// The status is checked before the body is buffered: an error
	// response is a StatusError whatever its body weighs, so a misbehaving
	// service cannot hide a refusal behind an oversized error body.
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, &StatusError{StatusCode: resp.StatusCode, Body: p.readErrorBody(resp.Body)}
	}
	body, err := p.readBody(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("query: read response: %w", err)
	}
	return body, nil
}

// readBody buffers a response body under the configured limit: a body over
// the limit is refused before it is whole in memory.
func (p *HTTPProvider) readBody(body io.Reader) ([]byte, error) {
	if p.limit <= 0 {
		return io.ReadAll(body)
	}
	// One byte past the limit detects the exceed; at the maximum int64
	// there is no one past a response can reach, so the limit itself is
	// the read and nothing overflows.
	probe := p.limit + 1
	if probe <= 0 {
		probe = p.limit
	}
	buffered, err := io.ReadAll(io.LimitReader(body, probe))
	if err != nil {
		return nil, err
	}
	if int64(len(buffered)) > p.limit {
		return nil, ErrResponseTooLarge
	}
	return buffered, nil
}

// readErrorBody reads an error response's diagnostic body, best effort:
// the status is the failure, so a body that cannot be read whole (or at
// all) never masks it.
func (p *HTTPProvider) readErrorBody(body io.Reader) []byte {
	if p.limit <= 0 {
		buffered, _ := io.ReadAll(body)
		return buffered
	}
	buffered, _ := io.ReadAll(io.LimitReader(body, p.limit))
	return buffered
}

func (p *HTTPProvider) post(ctx context.Context, path string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("ingest: build request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if p.auth != nil {
		p.auth(req)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ingest: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// The status wins over the body's weight here too: a refusal is a
	// StatusError whatever its body weighs.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &StatusError{
			StatusCode: resp.StatusCode,
			Body:       p.readErrorBody(resp.Body),
		}
	}

	respBody, err := p.readBody(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ingest: read response: %w", err)
	}

	return respBody, nil
}
