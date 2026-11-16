package traust

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestQueryOrdersAllFollowsPageCursorsAndStops(t *testing.T) {
	first := uuid.MustParse("00000000-0000-4000-8000-000000000001")
	second := uuid.MustParse("00000000-0000-4000-8000-000000000002")
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
				var query OrderQueryRequest
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
					_ = json.NewEncoder(w).Encode(OrderQueryPage{Items: []OrderQueryItem{{Order: RunOrder{Id: &first}, Cursor: "item-one"}}, Next: new("batch-two")})
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
				_ = json.NewEncoder(w).Encode(OrderQueryPage{Items: []OrderQueryItem{{Order: RunOrder{Id: &second}, Cursor: "item-two"}}})
			}))
			defer server.Close()
			client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("fixture-token")})
			if err != nil {
				t.Fatal(err)
			}
			var found []uuid.UUID
			var lastError error
			for item, err := range client.QueryOrdersAll(context.Background(), OrderQueryRequest{Q: new("matching-scope"), Cursor: tt.start}) {
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
				var problem *ProblemError
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
