package traust

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

	"github.com/google/uuid"
)

func products(client *Client, limit int) ([]Product, error) {
	var found []Product
	for p, err := range client.ListProductsAll(context.Background(), &ListProductsParams{Limit: &limit}) {
		if err != nil {
			return found, err
		}
		found = append(found, p)
	}
	return found, nil
}

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

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"first-token", "second-token"} {
		writeFile(t, tokenPath, want)
		if _, err := products(client, 1); err != nil {
			t.Fatal(err)
		}
	}
	for i, want := range []string{"Bearer first-token", "Bearer second-token"} {
		if seen[i] != want {
			t.Fatalf("request %d sent %q, want %q", i, seen[i], want)
		}
	}
}

func TestListProductsAllFollowsCursorAndReportsProblems(t *testing.T) {
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

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	found, err := products(client, 2)
	if err != nil {
		t.Fatal(err)
	}
	var slugs []string
	for _, p := range found {
		slugs = append(slugs, p.Slug)
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

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = products(client, 10)
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

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	found, err := products(client, 10)
	if err != nil {
		t.Fatal(err)
	}
	count := len(found)
	if count != 1 || atomic.LoadInt32(&calls) != 2 {
		t.Fatalf("items=%d calls=%d, want 1 item over 2 calls", count, calls)
	}
}

// The scheduling door's refusal is a problem+json document: the typed deny
// code and its issues ride ProblemError so a caller can branch on them
// without parsing the detail string.
func TestPreflightSurfacesTheDenyProblem(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "t")

	var method, path string
	code := RoleDenied
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, path = r.Method, r.URL.Path
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(Problem{
			Type:      "about:blank",
			Title:     "Forbidden",
			Status:    http.StatusForbidden,
			Detail:    "the stage may not run on this subject now",
			Code:      &code,
			Issues:    &[]DenyIssue{{Code: RoleDenied, Detail: "the lane's role may not run this stage"}},
			RequestId: "req-test",
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Preflight(context.Background(), PreflightRequest{Asset: "repository"})
	var problemErr *ProblemError
	if !errors.As(err, &problemErr) {
		t.Fatalf("got %v (%T), want *ProblemError", err, err)
	}
	if method != http.MethodPost || path != "/v1/preflight" {
		t.Fatalf("request %s %s, want POST /v1/preflight", method, path)
	}
	if problemErr.StatusCode != http.StatusForbidden {
		t.Fatalf("status %d, want 403", problemErr.StatusCode)
	}
	if problemErr.Problem.Code == nil || *problemErr.Problem.Code != RoleDenied {
		t.Fatalf("code %v, want %q", problemErr.Problem.Code, RoleDenied)
	}
	if problemErr.Problem.Issues == nil || len(*problemErr.Problem.Issues) != 1 || (*problemErr.Problem.Issues)[0].Code != RoleDenied {
		t.Fatalf("issues %+v, want one %q issue", problemErr.Problem.Issues, RoleDenied)
	}
}

// The re-keyed read names the work item in the path, and an effort core
// refuses (here: none for that work item) reads as *ProblemError, never a
// zero-value effort.
func TestGetRemediationNamesTheWorkItemAndRefusesAbsence(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "t")

	var path string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		problem(w, http.StatusNotFound, "no remediation effort names this work item")
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	id := uuid.New()
	_, err = client.GetRemediation(context.Background(), id)
	var problemErr *ProblemError
	if !errors.As(err, &problemErr) {
		t.Fatalf("got %v (%T), want *ProblemError", err, err)
	}
	if path != "/v1/remediations/"+id.String() {
		t.Fatalf("requested %s, want /v1/remediations/%s", path, id)
	}
	if problemErr.StatusCode != http.StatusNotFound {
		t.Fatalf("status %d, want 404", problemErr.StatusCode)
	}
}

func TestNewClientValidatesConfig(t *testing.T) {
	if _, err := NewClient(Config{Tokens: FileTokenSource{Path: "x"}}); err == nil {
		t.Fatal("empty BaseURL accepted")
	}
	if _, err := NewClient(Config{BaseURL: "http://x"}); err == nil {
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
	client, err := NewClient(Config{
		BaseURL:    server.URL,
		Tokens:     FileTokenSource{Path: tokenPath},
		HTTPClient: caller,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = products(client, 10)
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

func TestDecode(t *testing.T) {
	order := &RunOrder{Asset: "a"}
	for _, tc := range []struct {
		name        string
		status      int
		contentType string
		body        string
		value       *RunOrder
		problem     bool
	}{
		{"decoded body", 200, "application/json", `{"asset":"a"}`, order, false},
		{"problem document", 409, "application/problem+json", `{"title":"Conflict","status":409}`, nil, true},
		{"refusal without a document", 502, "text/plain", "bad gateway", nil, true},
		{"2xx without a body", 200, "application/json", "", nil, false},
		{"2xx with a null body", 200, "application/json", "null", order, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response := &http.Response{StatusCode: tc.status, Header: http.Header{"Content-Type": {tc.contentType}}}
			got, err := Decode(response, []byte(tc.body), tc.value)
			var problem *ProblemError
			switch {
			case tc.problem:
				if !errors.As(err, &problem) || problem.StatusCode != tc.status || got != nil {
					t.Fatalf("got=%v err=%v", got, err)
				}
			case tc.value == order && tc.body != "null":
				if err != nil || got != order {
					t.Fatalf("got=%v err=%v", got, err)
				}
			default:
				if err == nil || errors.As(err, &problem) || got != nil {
					t.Fatalf("a 2xx without a body must fail as such: got=%v err=%v", got, err)
				}
			}
		})
	}
}
