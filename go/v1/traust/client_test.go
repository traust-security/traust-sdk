package traust

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync/atomic"
	"testing"
	"time"
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
		_ = json.NewEncoder(w).Encode(Principal{Identity: "machine"})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"first-token", "second-token"} {
		writeFile(t, tokenPath, want)
		if _, err := client.Whoami(context.Background()); err != nil {
			t.Fatal(err)
		}
	}
	for i, want := range []string{"Bearer first-token", "Bearer second-token"} {
		if seen[i] != want {
			t.Fatalf("request %d sent %q, want %q", i, seen[i], want)
		}
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
	_, err = client.Whoami(context.Background())
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

// A transient 503 on a GET is retried; the read succeeds without the
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
		_ = json.NewEncoder(w).Encode(Principal{Identity: "machine"})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Tokens: FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.Whoami(context.Background()); err != nil {
		t.Fatal(err)
	}
	if atomic.LoadInt32(&calls) != 2 {
		t.Fatalf("calls=%d, want 2", calls)
	}
}

func TestNewClientValidatesConfig(t *testing.T) {
	if _, err := NewClient(Config{Tokens: FileTokenSource{Path: "x"}}); err == nil {
		t.Fatal("empty BaseURL accepted")
	}
}

// A client with only a URL serves people (OnBehalfOf) and the liveness check; a call that names
// no caller is refused before it is sent, never made anonymously.
func TestAClientWithOnlyAURLNeedsACallerPerCall(t *testing.T) {
	var auth []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = append(auth, r.URL.Path+" "+r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/healthz":
			_, _ = w.Write([]byte(`{"status":"ok"}`))
		default:
			_, _ = w.Write([]byte(`{"identity":"ada@example.com","kind":"human","display":"Ada"}`))
		}
	}))
	defer srv.Close()
	client, err := NewClient(Config{BaseURL: srv.URL})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.Healthz(context.Background()); err != nil {
		t.Fatalf("healthz without a caller: %v", err)
	}
	if _, err := client.Whoami(OnBehalfOf(context.Background(), "ada-token")); err != nil {
		t.Fatalf("a person's call: %v", err)
	}
	if _, err := client.Whoami(context.Background()); !errors.Is(err, ErrNoIdentity) {
		t.Fatalf("a call naming no caller: %v, want ErrNoIdentity", err)
	}
	want := []string{"/v1/healthz ", "/v1/whoami Bearer ada-token"}
	if !slices.Equal(auth, want) {
		t.Fatalf("core saw %q, want %q (the refused call never sent)", auth, want)
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
	mux.HandleFunc("/v1/whoami", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/elsewhere", http.StatusFound)
	})
	mux.HandleFunc("/elsewhere", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&targetHits, 1)
		_ = json.NewEncoder(w).Encode(Principal{Identity: "machine"})
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
	_, err = client.Whoami(context.Background())
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

// closeTracker records the one close a refusal owes the connection.
type closeTracker struct {
	io.Reader
	closed bool
}

func (c *closeTracker) Close() error { c.closed = true; return nil }

// Refusal hands back the typed problem and closes the body it read; a 2xx passes untouched.
func TestRefusal(t *testing.T) {
	body := &closeTracker{Reader: strings.NewReader(`{"title":"Forbidden","status":403}`)}
	refusal := &http.Response{
		StatusCode: http.StatusForbidden,
		Header:     http.Header{"Content-Type": {"application/problem+json"}},
		Body:       body,
	}
	err := Refusal(refusal)
	var problem *ProblemError
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusForbidden || problem.Problem == nil || problem.Problem.Title != "Forbidden" {
		t.Fatalf("refusal = %v, want the typed problem document", err)
	}
	if !body.closed {
		t.Fatal("the refusal's body was left open")
	}
	if err := Refusal(&http.Response{StatusCode: http.StatusOK, Body: http.NoBody}); err != nil {
		t.Fatalf("a 2xx refused: %v", err)
	}
}

// Pages walks one fetch per page until no next cursor, stops at the first error, and stops
// fetching when the consumer stops.
func TestPages(t *testing.T) {
	var fetches int
	var found []string
	for item, err := range Pages(func(cursor *string) ([]string, *string, error) {
		fetches++
		if cursor == nil {
			return []string{"a", "b"}, strp("two"), nil
		}
		return []string{"c"}, nil, nil
	}, nil) {
		if err != nil {
			t.Fatal(err)
		}
		found = append(found, item)
	}
	if fetches != 2 || strings.Join(found, "") != "abc" {
		t.Fatalf("%d fetches walked %v, want 2 over [a b c]", fetches, found)
	}
	for _, err := range Pages(func(*string) ([]string, *string, error) {
		return nil, nil, ErrForbidden
	}, nil) {
		if !errors.Is(err, ErrForbidden) {
			t.Fatalf("the fetch's error was lost: %v", err)
		}
	}
	fetches = 0
	for range Pages(func(cursor *string) ([]string, *string, error) {
		fetches++
		return []string{"a"}, strp("two"), nil
	}, nil) {
		break
	}
	if fetches != 1 {
		t.Fatalf("%d fetches after the consumer stopped, want 1", fetches)
	}
}

// NewPoller names the work its token does and polls until the work finishes.
func TestNewPoller(t *testing.T) {
	var polls int
	poller := NewPoller("an-order", func(context.Context) (string, bool, error) {
		polls++
		if polls < 2 {
			return "building", false, nil
		}
		return "done", true, nil
	})
	if poller.ResumeToken() != "an-order" {
		t.Fatalf("token %q, want an-order", poller.ResumeToken())
	}
	latest, done, err := poller.Poll(context.Background())
	if err != nil || done || latest != "building" {
		t.Fatalf("poll %q done %v err %v, want the work unfinished", latest, done, err)
	}
	if _, err := poller.Wait(context.Background(), 0); err == nil {
		t.Fatal("Wait accepted a zero interval")
	}
	final, err := poller.Wait(context.Background(), time.Millisecond)
	if err != nil || final != "done" {
		t.Fatalf("wait %q err %v, want the finished work", final, err)
	}
}

// The generated client's response types decode core's problem media, whatever its parameters.
func TestWithResponseDecodesCoreProblemMediaAndValidationErrors(t *testing.T) {
	for _, contentType := range []string{"application/problem+json", "application/problem+json; charset=utf-8", "application/json"} {
		t.Run(contentType, func(t *testing.T) {
			code := RoleDenied
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", contentType)
				w.WriteHeader(http.StatusUnauthorized)
				_ = json.NewEncoder(w).Encode(Problem{
					Type: "about:blank", Title: http.StatusText(http.StatusUnauthorized), Status: http.StatusUnauthorized,
					Detail: "typed refusal", RequestId: "req-with-response", Code: &code,
					Issues: &[]DenyIssue{{Code: code, Detail: "typed issue"}},
				})
			}))
			defer server.Close()
			client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("test-token")})
			if err != nil {
				t.Fatal(err)
			}
			caller, err := client.Raw().WhoamiWithResponse(context.Background())
			if err != nil || caller == nil || caller.StatusCode() != http.StatusUnauthorized || caller.JSON401 == nil {
				t.Fatalf("problem media did not populate typed 401: response=%+v err=%v", caller, err)
			}
			problem := caller.JSON401
			if problem.Detail != "typed refusal" || problem.RequestId != "req-with-response" ||
				problem.Code == nil || *problem.Code != code || problem.Issues == nil || len(*problem.Issues) != 1 {
				t.Fatalf("typed problem was decoded as an empty or validation-error body: %+v", problem)
			}
		})
	}
}
