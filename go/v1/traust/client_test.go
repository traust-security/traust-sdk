package traust

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

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

func TestPrepareVerificationPreservesTypedRefusals(t *testing.T) {
	for _, tc := range []struct {
		name        string
		contentType string
		body        string
		problem     bool
	}{
		{
			name:        "core problem instead of validation errors",
			contentType: "application/problem+json",
			body:        `{"type":"about:blank","title":"Unprocessable Entity","status":422,"detail":"verification requires a closed work item","request_id":"req-verification"}`,
			problem:     true,
		},
		{
			name:        "malformed problem document",
			contentType: "application/problem+json",
			body:        `{"title":"Unprocessable Entity","detail":`,
		},
		{
			name:        "non-JSON refusal",
			contentType: "text/plain",
			body:        "verification refused\n",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer server.Close()

			client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("test-token")})
			if err != nil {
				t.Fatal(err)
			}
			got, err := client.PrepareVerification(context.Background(), uuid.New())
			var refusal *ProblemError
			if !errors.As(err, &refusal) || got != nil {
				t.Fatalf("got=%v err=%v (%T), want nil and *ProblemError", got, err, err)
			}
			if refusal.StatusCode != http.StatusUnprocessableEntity || refusal.Body != tc.body {
				t.Fatalf("status=%d body=%q, want 422 and %q", refusal.StatusCode, refusal.Body, tc.body)
			}
			if !tc.problem {
				if refusal.Problem != nil {
					t.Fatalf("problem=%+v, want no decoded problem", refusal.Problem)
				}
				return
			}
			document := refusal.Problem
			if document == nil || document.Type != "about:blank" || document.Title != "Unprocessable Entity" ||
				document.Status != http.StatusUnprocessableEntity || document.Detail != "verification requires a closed work item" ||
				document.RequestId != "req-verification" {
				t.Fatalf("problem=%+v, want the original core refusal document", document)
			}
		})
	}
}

func TestMalformedRefusalCannotLeakBodyOrForgeLogLines(t *testing.T) {
	const body = "{invalid\nsensitive-refusal-marker\r\nFORGED EVENT\x1b[31m"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()
	client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("test-token")})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.PrepareVerification(context.Background(), uuid.New())
	var refusal *ProblemError
	if !errors.As(err, &refusal) || refusal.StatusCode != http.StatusUnprocessableEntity || refusal.Body != body {
		t.Fatalf("original refusal is unavailable for explicit inspection: %v", err)
	}
	formatted := err.Error()
	if strings.Contains(formatted, "sensitive-refusal-marker") || strings.Contains(formatted, "FORGED EVENT") || strings.ContainsAny(formatted, "\r\n\x1b") {
		t.Fatalf("formatted refusal leaked untrusted response text: %q", formatted)
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

func TestWithResponseDecodesCoreProblemMediaAndValidationErrors(t *testing.T) {
	for _, contentType := range []string{"application/problem+json", "application/problem+json; charset=utf-8", "application/json"} {
		t.Run(contentType, func(t *testing.T) {
			code := RoleDenied
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				status := http.StatusUnauthorized
				if strings.HasSuffix(r.URL.Path, "/verification-input") {
					status = http.StatusUnprocessableEntity
				}
				w.Header().Set("Content-Type", contentType)
				w.WriteHeader(status)
				_ = json.NewEncoder(w).Encode(Problem{
					Type: "about:blank", Title: http.StatusText(status), Status: status,
					Detail: "typed refusal", RequestId: "req-with-response", Code: &code,
					Issues: &[]DenyIssue{{Code: code, Detail: "typed issue"}},
				})
			}))
			defer server.Close()
			client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("test-token")})
			if err != nil {
				t.Fatal(err)
			}
			repository, err := client.GetRepositoryWithResponse(context.Background(), uuid.New())
			if err != nil || repository == nil || repository.StatusCode() != http.StatusUnauthorized || repository.JSON401 == nil {
				t.Fatalf("problem media did not populate typed 401: response=%+v err=%v", repository, err)
			}
			verification, err := client.PrepareVerificationWithResponse(context.Background(), uuid.New())
			if err != nil || verification == nil || verification.StatusCode() != http.StatusUnprocessableEntity || verification.JSON422 == nil {
				t.Fatalf("problem media did not populate typed 422: response=%+v err=%v", verification, err)
			}
			for _, problem := range []*Problem{repository.JSON401, verification.JSON422} {
				if problem.Detail != "typed refusal" || problem.RequestId != "req-with-response" ||
					problem.Code == nil || *problem.Code != code || problem.Issues == nil || len(*problem.Issues) != 1 {
					t.Fatalf("typed problem was decoded as an empty or validation-error body: %+v", problem)
				}
			}
		})
	}
}

// The set-scoped reads decode a non-empty page and an empty one, and the repository set
// rides the request as one exploded query value per id — the encoding core's bound reads.
func TestSetScopedReadsDecodeTheirPages(t *testing.T) {
	set := []uuid.UUID{uuid.New(), uuid.New()}
	reads := []struct {
		name string
		page string
		read func(*Client) (int, error)
	}{
		{"ListAssessmentsIn", `[{"id":"` + uuid.New().String() + `","occurred_at":"2026-01-01T00:00:00Z","repository_ref_id":"` + uuid.New().String() + `","assessor":"scanner","assessor_version":"1.0","kind":"sbom","status":"recorded"}]`, func(c *Client) (int, error) {
			found, err := c.ListAssessmentsIn(context.Background(), &ListAssessmentsInParams{RepositoryIds: &set})
			return len(found), err
		}},
		{"ListRefsIn", `[{"id":"` + uuid.New().String() + `","repository_id":"` + uuid.New().String() + `","sha":"4c7f0b9a","name":"main"}]`, func(c *Client) (int, error) {
			found, err := c.ListRefsIn(context.Background(), &ListRefsInParams{RepositoryIds: &set})
			return len(found), err
		}},
		{"SummarizeFindingsIn", `[{"key":null,"findings":3,"open":1,"accepted":1,"false_positives":0,"hardening":0,"resolved":1,"open_by_severity":{"high":1}}]`, func(c *Client) (int, error) {
			found, err := c.SummarizeFindingsIn(context.Background(), &SummarizeFindingsInParams{RepositoryIds: &set})
			return len(found), err
		}},
	}
	for _, read := range reads {
		t.Run(read.name, func(t *testing.T) {
			for _, page := range []struct {
				name string
				body string
				want int
			}{
				{"a non-empty page", read.page, 1},
				{"an empty page", "[]", 0},
			} {
				t.Run(page.name, func(t *testing.T) {
					var query url.Values
					server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						query = r.URL.Query()
						w.Header().Set("Content-Type", "application/json")
						_, _ = w.Write([]byte(page.body))
					}))
					defer server.Close()

					client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("t")})
					if err != nil {
						t.Fatal(err)
					}
					count, err := read.read(client)
					if err != nil {
						t.Fatal(err)
					}
					if count != page.want {
						t.Fatalf("decoded %d rows from %s, want %d", count, page.body, page.want)
					}
					if ids := query["repository_ids"]; len(ids) != len(set) || ids[0] != set[0].String() || ids[1] != set[1].String() {
						t.Fatalf("the request carried repository_ids=%v, want one value per id of %v", ids, set)
					}
				})
			}
		})
	}
}

// The set-scoped reads fail closed on their documented edges: a 5xx is retried and then
// passes through as a typed problem, a malformed 2xx body is an error and never a
// zero-value page, and a problem code the generated set does not name still decodes its
// document.
func TestSetScopedReadsFailClosed(t *testing.T) {
	set := []uuid.UUID{uuid.New(), uuid.New()}
	reads := []struct {
		name string
		read func(*Client) (int, error)
	}{
		{"ListAssessmentsIn", func(c *Client) (int, error) {
			found, err := c.ListAssessmentsIn(context.Background(), &ListAssessmentsInParams{RepositoryIds: &set})
			return len(found), err
		}},
		{"ListRefsIn", func(c *Client) (int, error) {
			found, err := c.ListRefsIn(context.Background(), &ListRefsInParams{RepositoryIds: &set})
			return len(found), err
		}},
		{"SummarizeFindingsIn", func(c *Client) (int, error) {
			found, err := c.SummarizeFindingsIn(context.Background(), &SummarizeFindingsInParams{RepositoryIds: &set})
			return len(found), err
		}},
	}
	for _, read := range reads {
		t.Run(read.name, func(t *testing.T) {
			for _, edge := range []struct {
				name        string
				status      int
				contentType string
				body        string
				attempts    int32 // round trips the edge takes, retries included
				problem     bool  // a *ProblemError, else a decode error
				document    bool  // the problem document decoded
				code        string
			}{
				{
					name: "a 503 passes through once the retries exhaust", status: http.StatusServiceUnavailable,
					contentType: "text/plain", body: "unavailable\n", attempts: 2, problem: true,
				},
				{
					name: "a malformed body on 200", status: http.StatusOK,
					contentType: "application/json", body: `[{"id":`, attempts: 1,
				},
				{
					name: "an unrecognized problem code", status: http.StatusConflict,
					contentType: "application/problem+json",
					body:        `{"type":"about:blank","title":"Conflict","status":409,"code":"unrecognized","request_id":"req-set"}`,
					attempts:    1, problem: true, document: true, code: "unrecognized",
				},
			} {
				t.Run(edge.name, func(t *testing.T) {
					var hits int32
					var query url.Values
					server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						atomic.AddInt32(&hits, 1)
						query = r.URL.Query()
						w.Header().Set("Content-Type", edge.contentType)
						w.WriteHeader(edge.status)
						_, _ = w.Write([]byte(edge.body))
					}))
					defer server.Close()

					client, err := NewClient(Config{
						BaseURL: server.URL, Tokens: StaticToken("t"),
						Retries: 1, Backoff: time.Millisecond,
					})
					if err != nil {
						t.Fatal(err)
					}
					count, err := read.read(client)
					if err == nil || count != 0 {
						t.Fatalf("the read returned %d items and %v; every refusal must fail closed", count, err)
					}
					if ids := query["repository_ids"]; len(ids) != len(set) || ids[0] != set[0].String() || ids[1] != set[1].String() {
						t.Fatalf("the request carried repository_ids=%v, want one value per id of %v", ids, set)
					}
					if atomic.LoadInt32(&hits) != edge.attempts {
						t.Fatalf("the edge took %d round trips, want %d", hits, edge.attempts)
					}
					var problem *ProblemError
					if errors.As(err, &problem) != edge.problem {
						t.Fatalf("got %v (%T), want a typed problem: %v", err, err, edge.problem)
					}
					if !edge.problem {
						return
					}
					if problem.StatusCode != edge.status || (problem.Problem != nil) != edge.document {
						t.Fatalf("status=%d document=%+v, want %d documented=%v", problem.StatusCode, problem.Problem, edge.status, edge.document)
					}
					if edge.code != "" && (problem.Problem.Code == nil || *problem.Problem.Code != DenyCode(edge.code) || problem.Problem.RequestId != "req-set") {
						t.Fatalf("problem=%+v, want code %q and its request id", problem.Problem, edge.code)
					}
				})
			}
		})
	}
}
