package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

// The bound client sends the file's token on every request and re-reads it,
// so a rotated projected token is picked up without a restart.
func TestFileTokenSourceRotates(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "first-token\n")

	var seen []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = append(seen, r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(PageProduct{})
	}))
	defer server.Close()

	client, err := NewBoundClient(ClientConfig{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"first-token", "second-token"} {
		writeFile(t, tokenPath, want)
		if err := client.EachProduct(context.Background(), 1, func(Product) error { return nil }); err != nil {
			t.Fatal(err)
		}
	}
	for i, want := range []string{"Bearer first-token", "Bearer second-token"} {
		if seen[i] != want {
			t.Fatalf("request %d sent %q, want %q", i, seen[i], want)
		}
	}
}

func TestEachProductFollowsCursorAndReportsProblems(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "t")

	page := func(items []Product, next *string) PageProduct {
		return PageProduct{Items: items, Next: next}
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer t" {
			problem(w, http.StatusUnauthorized, "a bearer token is required")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("cursor") == "" {
			_ = json.NewEncoder(w).Encode(page([]Product{{Slug: "a"}, {Slug: "b"}}, strp("c2")))
			return
		}
		_ = json.NewEncoder(w).Encode(page([]Product{{Slug: "c"}}, nil))
	}))
	defer server.Close()

	client, err := NewBoundClient(ClientConfig{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	var slugs []string
	if err := client.EachProduct(context.Background(), 2, func(p Product) error {
		slugs = append(slugs, p.Slug)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if len(slugs) != 3 || slugs[0] != "a" || slugs[2] != "c" {
		t.Fatalf("walked %v, want [a b c]", slugs)
	}
}

func TestProblemErrorCarriesTheDocument(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "t")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		problem(w, http.StatusUnauthorized, "a bearer token is required")
	}))
	defer server.Close()

	client, err := NewBoundClient(ClientConfig{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	err = client.EachProduct(context.Background(), 10, func(Product) error { return nil })
	var problemErr *ProblemError
	if !errors.As(err, &problemErr) {
		t.Fatalf("got %v (%T), want *ProblemError", err, err)
	}
	if problemErr.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status %d, want 401", problemErr.StatusCode)
	}
	if problemErr.Problem == nil || problemErr.Problem.Title != "Unauthorized" {
		t.Fatalf("problem %+v, want Unauthorized", problemErr.Problem)
	}
	if problemErr.Problem.RequestId == "" {
		t.Fatal("problem carries no request id")
	}
}

// A transient 503 on a GET is retried; the walk succeeds without the
// consumer knowing.
func TestRetryAbsorbsTransientFailures(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "t")

	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(PageProduct{Items: []Product{{Slug: "a"}}})
	}))
	defer server.Close()

	client, err := NewBoundClient(ClientConfig{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	var count int
	if err := client.EachProduct(context.Background(), 10, func(Product) error { count++; return nil }); err != nil {
		t.Fatal(err)
	}
	if count != 1 || atomic.LoadInt32(&calls) != 2 {
		t.Fatalf("items=%d calls=%d, want 1 item over 2 calls", count, calls)
	}
}

func TestNewBoundClientValidatesConfig(t *testing.T) {
	if _, err := NewBoundClient(ClientConfig{Tokens: FileTokenSource{Path: "x"}}); err == nil {
		t.Fatal("empty BaseURL accepted")
	}
	if _, err := NewBoundClient(ClientConfig{BaseURL: "http://x"}); err == nil {
		t.Fatal("missing Tokens accepted")
	}
}

// A caller-supplied http.Client keeps its CheckRedirect policy: the bound
// client replaces only the transport. Regression test for a review finding
// (girish1430): reconstructing the client from Transport+Timeout alone
// silently dropped CheckRedirect and Jar, following redirects a caller
// configured away.
func TestCallerHTTPClientPolicySurvives(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "t")

	var targetHits int32
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/products", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/elsewhere", http.StatusFound)
	})
	mux.HandleFunc("/elsewhere", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&targetHits, 1)
		_ = json.NewEncoder(w).Encode(PageProduct{})
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	caller := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	client, err := NewBoundClient(ClientConfig{
		BaseURL:    server.URL,
		Tokens:     FileTokenSource{Path: tokenPath},
		HTTPClient: caller,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = client.EachProduct(context.Background(), 10, func(Product) error { return nil })
	if err == nil {
		t.Fatal("302 surfaced as success; CheckRedirect was dropped")
	}
	var problemErr *ProblemError
	if !errors.As(err, &problemErr) || problemErr.StatusCode != http.StatusFound {
		t.Fatalf("got %v, want a problem error carrying 302", err)
	}
	if atomic.LoadInt32(&targetHits) != 0 {
		t.Fatal("redirect target was followed; CheckRedirect was dropped")
	}
	// The caller's client itself stays unmodified.
	if caller.Transport != nil {
		t.Fatal("caller's transport was mutated")
	}
}

func problem(w http.ResponseWriter, status int, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Problem{
		Type:      "about:blank",
		Title:     http.StatusText(status),
		Status:    status,
		Detail:    detail,
		RequestId: "req-test",
	})
}

func strp(s string) *string { return &s }
