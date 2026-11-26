package findings_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/findings"
)

// The assessments read rides its collection (GET /v1/assessments?repository_ids=...) and
// decodes the collection's own page shape (items plus cursor), the repository set riding the
// request as one exploded query value per id — the encoding core's bound reads.
func TestSetScopedReadsDecodeTheirPages(t *testing.T) {
	set := []uuid.UUID{uuid.New(), uuid.New()}
	pages := []struct {
		name string
		body string
		want int
	}{
		{"a non-empty page", `{"items":[{"id":"` + uuid.New().String() + `","occurred_at":"2026-01-01T00:00:00Z","repository_ref_id":"` + uuid.New().String() + `","assessor":"scanner","assessor_version":"1.0","kind":"sbom","status":"recorded"}],"next":null}`, 1},
		{"an empty page", `{"items":[],"next":null}`, 0},
	}
	for _, page := range pages {
		t.Run(page.name, func(t *testing.T) {
			var query url.Values
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				query = r.URL.Query()
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(page.body))
			}))
			defer server.Close()

			client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("t")})
			if err != nil {
				t.Fatal(err)
			}
			found, err := findings.New(client).ListAssessments(context.Background(), &traust.ListAssessmentsParams{RepositoryIds: &set})
			if err != nil {
				t.Fatal(err)
			}
			if len(found.Items) != page.want {
				t.Fatalf("decoded %d rows from %s, want %d", len(found.Items), page.body, page.want)
			}
			if ids := query["repository_ids"]; len(ids) != len(set) || ids[0] != set[0].String() || ids[1] != set[1].String() {
				t.Fatalf("the request carried repository_ids=%v, want one value per id of %v", ids, set)
			}
		})
	}
}

// The assessments read fails closed on its documented edges: a 5xx is retried and then passes
// through as a typed problem, a malformed 2xx body is an error and never a zero-value page,
// and a problem code the generated set does not name still decodes its document.
func TestSetScopedReadsFailClosed(t *testing.T) {
	set := []uuid.UUID{uuid.New(), uuid.New()}
	read := func(c *findings.Client) (int, error) {
		found, err := c.ListAssessments(context.Background(), &traust.ListAssessmentsParams{RepositoryIds: &set})
		if found == nil {
			return 0, err
		}
		return len(found.Items), err
	}
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
			contentType: "application/json", body: `{"items":[{"id":`, attempts: 1,
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

			client, err := traust.NewClient(traust.Config{
				BaseURL: server.URL, Tokens: traust.StaticToken("t"),
				Retries: 1, Backoff: time.Millisecond,
			})
			if err != nil {
				t.Fatal(err)
			}
			count, err := read(findings.New(client))
			if err == nil || count != 0 {
				t.Fatalf("the read returned %d items and %v; every refusal must fail closed", count, err)
			}
			if ids := query["repository_ids"]; len(ids) != len(set) || ids[0] != set[0].String() || ids[1] != set[1].String() {
				t.Fatalf("the request carried repository_ids=%v, want one value per id of %v", ids, set)
			}
			if atomic.LoadInt32(&hits) != edge.attempts {
				t.Fatalf("the edge took %d round trips, want %d", hits, edge.attempts)
			}
			var problem *traust.ProblemError
			if errors.As(err, &problem) != edge.problem {
				t.Fatalf("got %v (%T), want a typed problem: %v", err, err, edge.problem)
			}
			if !edge.problem {
				return
			}
			if problem.StatusCode != edge.status || (problem.Problem != nil) != edge.document {
				t.Fatalf("status=%d document=%+v, want %d documented=%v", problem.StatusCode, problem.Problem, edge.status, edge.document)
			}
			if edge.code != "" && (problem.Problem.Code == nil || *problem.Problem.Code != traust.DenyCode(edge.code) || problem.Problem.RequestId != "req-set") {
				t.Fatalf("problem=%+v, want code %q and its request id", problem.Problem, edge.code)
			}
		})
	}
}
