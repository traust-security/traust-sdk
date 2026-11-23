package runs_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/runs"
)

// The scheduling door's refusal is a problem+json document: the typed deny
// code and its issues ride ProblemError so a caller can branch on them
// without parsing the detail string.
func TestPreflightSurfacesTheDenyProblem(t *testing.T) {
	var method, path string
	code := traust.RoleDenied
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, path = r.Method, r.URL.Path
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(traust.Problem{
			Type:      "about:blank",
			Title:     "Forbidden",
			Status:    http.StatusForbidden,
			Detail:    "the stage may not run on this subject now",
			Code:      &code,
			Issues:    &[]traust.DenyIssue{{Code: traust.RoleDenied, Detail: "the lane's role may not run this stage"}},
			RequestId: "req-test",
		})
	}))
	defer server.Close()

	client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("t")})
	if err != nil {
		t.Fatal(err)
	}
	_, err = runs.New(client).Preflight(context.Background(), traust.PreflightRequest{Asset: "repository"})
	var problemErr *traust.ProblemError
	if !errors.As(err, &problemErr) {
		t.Fatalf("got %v (%T), want *ProblemError", err, err)
	}
	if method != http.MethodPost || path != "/v1/preflight" {
		t.Fatalf("request %s %s, want POST /v1/preflight", method, path)
	}
	if problemErr.StatusCode != http.StatusForbidden {
		t.Fatalf("status %d, want 403", problemErr.StatusCode)
	}
	if problemErr.Problem.Code == nil || *problemErr.Problem.Code != traust.RoleDenied {
		t.Fatalf("code %v, want %q", problemErr.Problem.Code, traust.RoleDenied)
	}
	if problemErr.Problem.Issues == nil || len(*problemErr.Problem.Issues) != 1 || (*problemErr.Problem.Issues)[0].Code != traust.RoleDenied {
		t.Fatalf("issues %+v, want one %q issue", problemErr.Problem.Issues, traust.RoleDenied)
	}
}

func TestQueryOrdersPreservesNullAndEmptyIDs(t *testing.T) {
	empty := []uuid.UUID{}
	for _, tt := range []struct {
		name string
		ids  *[]uuid.UUID
		want string
	}{
		{name: "unrestricted", ids: nil, want: "null"},
		{name: "empty", ids: &empty, want: "[]"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost || r.URL.Path != "/v1/orders/query" || r.URL.RawQuery != "" {
					t.Errorf("unexpected query transport: %s %s", r.Method, r.URL)
				}
				var body map[string]json.RawMessage
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Error(err)
				}
				if string(body["ids"]) != tt.want {
					t.Errorf("ids = %s, want %s", body["ids"], tt.want)
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(traust.OrderQueryPage{Items: []traust.OrderQueryItem{}})
			}))
			defer server.Close()
			client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("test-token")})
			if err != nil {
				t.Fatal(err)
			}
			page, err := runs.New(client).QueryOrders(context.Background(), traust.OrderQueryRequest{Ids: tt.ids})
			if err != nil {
				t.Fatal(err)
			}
			if len(page.Items) != 0 || page.Next != nil {
				t.Fatalf("unexpected empty page: %+v", page)
			}
		})
	}
}

func TestQueryOrdersKeepsPerItemCursorsAndTypedProblems(t *testing.T) {
	strp := func(s string) *string { return &s }
	tests := []struct {
		name        string
		status      int
		contentType string
		body        string
		request     traust.OrderQueryRequest
		success     bool
		problem     bool
		code        traust.DenyCode
		cancelled   bool
	}{
		{
			name: "per-item cursors and unknown repository", status: http.StatusOK,
			contentType: "application/json",
			body: `{"items":[
				{"order":{"id":"00000000-0000-4000-8000-000000000001"},"repository_url":"https://forge.example/a","cursor":"first"},
				{"order":{"id":"00000000-0000-4000-8000-000000000002"},"repository_url":null,"cursor":"second"}
			],"next":"second"}`,
			success: true,
		},
		{
			name: "invalid cursor", status: http.StatusUnprocessableEntity,
			contentType: "application/problem+json",
			body:        `{"type":"about:blank","title":"Unprocessable Entity","status":422,"detail":"invalid or mismatched order query cursor","request_id":"req-query"}`,
			request:     traust.OrderQueryRequest{Cursor: strp("invalid")}, problem: true,
		},
		{
			name: "malformed successful body", status: http.StatusOK,
			contentType: "application/json", body: `{"items":[{"order":`,
		},
		{
			name: "non-JSON successful body", status: http.StatusOK,
			contentType: "text/plain", body: "not a page\n",
		},
		{
			name: "null page", status: http.StatusOK,
			contentType: "application/json", body: "null",
		},
		{
			name: "missing page", status: http.StatusOK,
			contentType: "application/json",
		},
		{
			name: "raw service unavailable", status: http.StatusServiceUnavailable,
			contentType: "text/plain", body: "unavailable\n", problem: true,
		},
		{
			name: "typed unauthorized", status: http.StatusUnauthorized,
			contentType: "application/problem+json",
			body:        `{"type":"about:blank","title":"Unauthorized","status":401,"code":"role_denied","request_id":"req-query"}`,
			problem:     true, code: traust.RoleDenied,
		},
		{
			name: "unknown problem code", status: http.StatusConflict,
			contentType: "application/problem+json",
			body:        `{"type":"about:blank","title":"Conflict","status":409,"code":"unrecognized","request_id":"req-query"}`,
			problem:     true, code: traust.DenyCode("unrecognized"),
		},
		{
			name: "already cancelled context", status: http.StatusOK,
			contentType: "application/json", body: `{"items":[],"next":null}`, cancelled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var hits int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				atomic.AddInt32(&hits, 1)
				w.Header().Set("Content-Type", tt.contentType)
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.body))
			}))
			defer server.Close()
			client, err := traust.NewClient(traust.Config{
				BaseURL: server.URL, Tokens: traust.StaticToken("test-token"),
				Retries: 1, Backoff: time.Millisecond,
			})
			if err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if tt.cancelled {
				cancel()
			}
			page, err := runs.New(client).QueryOrders(ctx, tt.request)
			wantHits := int32(1)
			if tt.cancelled {
				wantHits = 0
			}
			if got := atomic.LoadInt32(&hits); got != wantHits {
				t.Fatalf("got %d round trips, want %d", got, wantHits)
			}
			if tt.success {
				if err != nil || page == nil {
					t.Fatalf("page=%+v err=%v, want a successful page", page, err)
				}
				if len(page.Items) != 2 || page.Items[0].Cursor != "first" || page.Items[1].Cursor != "second" {
					t.Fatalf("item positions lost: %+v", page)
				}
				if page.Next == nil || *page.Next != "second" || page.Items[1].RepositoryUrl != nil {
					t.Fatalf("page boundary or unknown repository changed: %+v", page)
				}
				return
			}
			if err == nil || page != nil {
				t.Fatalf("page=%+v err=%v, want no successful page and an error", page, err)
			}
			if tt.cancelled {
				if !errors.Is(err, context.Canceled) {
					t.Fatalf("got %v, want context.Canceled", err)
				}
				return
			}
			var refused *traust.ProblemError
			if errors.As(err, &refused) != tt.problem {
				t.Fatalf("got %v (%T), want typed problem=%v", err, err, tt.problem)
			}
			if !tt.problem {
				return
			}
			if refused.StatusCode != tt.status || refused.Body != tt.body {
				t.Fatalf("status=%d body=%q, want %d and %q", refused.StatusCode, refused.Body, tt.status, tt.body)
			}
			if tt.contentType == "text/plain" {
				if refused.Problem != nil {
					t.Fatalf("problem=%+v, want no decoded problem", refused.Problem)
				}
				return
			}
			document := refused.Problem
			if document == nil || document.Status != tt.status || document.RequestId != "req-query" {
				t.Fatalf("problem=%+v, want status %d and request id req-query", document, tt.status)
			}
			if tt.code != "" && (document.Code == nil || *document.Code != tt.code) {
				t.Fatalf("problem=%+v, want code %q", document, tt.code)
			}
		})
	}
}
