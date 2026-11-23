package runs_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/runs"
)

func problem(w http.ResponseWriter, status int, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(traust.Problem{
		Type: "about:blank", Title: http.StatusText(status), Status: status,
		Detail: detail, RequestId: "req-test",
	})
}

func TestQueryOrdersAllFollowsPageCursorsAndStops(t *testing.T) {
	first := uuid.MustParse("00000000-0000-4000-8000-000000000001")
	second := uuid.MustParse("00000000-0000-4000-8000-000000000002")
	new := func(s string) *string { return &s }
	for _, tt := range []struct {
		name       string
		start      *string
		refuse     bool
		stop       bool
		want       []uuid.UUID
		wantStatus int
		wantCalls  int
	}{
		{name: "complete filtered history", want: []uuid.UUID{first, second}, wantCalls: 2},
		{name: "resume supplied cursor", start: new("batch-two"), want: []uuid.UUID{second}, wantCalls: 1},
		{name: "surface later refusal", refuse: true, want: []uuid.UUID{first}, wantStatus: 403, wantCalls: 2},
		{name: "consumer stops without fetching next page", stop: true, want: []uuid.UUID{first}, wantCalls: 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				var query traust.OrderQueryRequest
				if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
					t.Error(err)
					w.WriteHeader(400)
					return
				}
				if query.Q == nil || *query.Q != "matching-scope" {
					problem(w, 422, "query scope missing")
					return
				}
				w.Header().Set("Content-Type", "application/json")
				if query.Cursor == nil {
					_ = json.NewEncoder(w).Encode(traust.OrderQueryPage{Items: []traust.OrderQueryItem{{Order: traust.RunOrder{Id: &first}, Cursor: "item-one"}}, Next: new("batch-two")})
					return
				}
				if *query.Cursor != "batch-two" {
					problem(w, 422, "wrong page boundary")
					return
				}
				if tt.refuse {
					problem(w, 403, "access revoked")
					return
				}
				_ = json.NewEncoder(w).Encode(traust.OrderQueryPage{Items: []traust.OrderQueryItem{{Order: traust.RunOrder{Id: &second}, Cursor: "item-two"}}})
			}))
			defer server.Close()
			client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("fixture-token")})
			if err != nil {
				t.Fatal(err)
			}
			var found []uuid.UUID
			var lastError error
			for item, err := range runs.New(client).QueryOrdersAll(context.Background(), traust.OrderQueryRequest{Q: new("matching-scope"), Cursor: tt.start}) {
				if err != nil {
					lastError = err
					break
				}
				found = append(found, *item.Order.Id)
				if tt.stop {
					break
				}
			}
			if !reflect.DeepEqual(found, tt.want) {
				t.Fatalf("history changed across cursor pages: %v, want %v", found, tt.want)
			}
			if tt.wantStatus == 0 {
				if lastError != nil {
					t.Fatal(lastError)
				}
			} else {
				var problem *traust.ProblemError
				if !errors.As(lastError, &problem) || problem.StatusCode != tt.wantStatus {
					t.Fatalf("lost iteration refusal: %v", lastError)
				}
			}
			if calls != tt.wantCalls {
				t.Fatalf("iteration made %d requests, want %d", calls, tt.wantCalls)
			}
		})
	}
}
