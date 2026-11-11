// Package traustest is a fake traust-core for testing code that uses the SDK: register the
// responses your code should see, take a client bound to the fake, and check what it sent.
package traustest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// Request is one call the fake core received.
type Request struct {
	Method string
	Path   string
	Query  url.Values
	Body   []byte
}

// Server is a fake traust-core on an httptest server.
type Server struct {
	mux      *http.ServeMux
	server   *httptest.Server
	mu       sync.Mutex
	requests []Request
}

// New starts a fake core closed when the test ends. A request no handler matches is a 404
// problem naming the route.
func New(t testing.TB) *Server {
	t.Helper()
	s := Start()
	t.Cleanup(s.Close)
	return s
}

// Start starts a fake core outside a test (examples, TestMain); Close stops it.
func Start() *Server {
	s := &Server{mux: http.NewServeMux()}
	s.mux.Handle("/", Problem(http.StatusNotFound, "no traustest handler for this route"))
	s.server = httptest.NewServer(http.HandlerFunc(s.serve))
	return s
}

// Close stops the fake core.
func (s *Server) Close() { s.server.Close() }

// Handle answers requests matching pattern (net/http patterns: "GET /v1/repositories/{id}").
func (s *Server) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

// Client is a client bound to the fake core with a static token.
func (s *Server) Client() *traust.Client {
	client, err := traust.NewClient(traust.Config{BaseURL: s.server.URL, Tokens: traust.StaticToken("traustest")})
	if err != nil {
		panic("traustest: " + err.Error()) // the URL and token are always set
	}
	return client
}

// URL is the fake core's base URL.
func (s *Server) URL() string { return s.server.URL }

// Requests are the calls received so far, in order.
func (s *Server) Requests() []Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Request(nil), s.requests...)
}

func (s *Server) serve(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	s.mu.Lock()
	s.requests = append(s.requests, Request{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Body: body})
	s.mu.Unlock()
	r.Body = io.NopCloser(bytes.NewReader(body))
	s.mux.ServeHTTP(w, r)
}

// JSON answers with status and v encoded as JSON.
func JSON(status int, v any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(v)
	})
}

// Problem answers with core's application/problem+json error.
func Problem(status int, detail string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(traust.Problem{
			Type: "about:blank", Title: http.StatusText(status), Status: status,
			Detail: detail, Instance: r.URL.Path, RequestId: "traustest",
		})
	})
}

// Status answers with status and no body (204 for an empty claim).
func Status(status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(status) })
}
