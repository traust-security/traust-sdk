package runs_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/runs"
)

// fakeSkill serves POST /v1/orders, recording the decoded submit body, and
// GET /v1/orders/{id} for the order it queued; the order has already finished.
func fakeSkill(t *testing.T) (*runs.Client, func() map[string]any) {
	t.Helper()
	id := uuid.New()
	var mu sync.Mutex
	var submitted map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v1/orders":
			mu.Lock()
			submitted = map[string]any{}
			_ = json.NewDecoder(r.Body).Decode(&submitted)
			mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(traust.RunOrder{Id: &id, Asset: "scan", State: traust.RunOrderStateQueued})
		case r.Method == http.MethodGet && r.URL.Path == "/v1/orders/"+id.String():
			finished := time.Now()
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(traust.RunOrder{Id: &id, Asset: "scan", State: traust.RunOrderStateSucceeded, FinishedAt: &finished})
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}))
	t.Cleanup(server.Close)
	client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("machine")})
	if err != nil {
		t.Fatal(err)
	}
	return runs.New(client), func() map[string]any {
		mu.Lock()
		defer mu.Unlock()
		return submitted
	}
}

func TestBeginSkillSubmitsTheAssetByName(t *testing.T) {
	repo := uuid.New()
	tests := []struct {
		name     string
		opts     []runs.RunOption
		wantRepo any
	}{
		{name: "no options", opts: nil, wantRepo: nil},
		{name: "for one repository", opts: []runs.RunOption{runs.ForRepository(repo)}, wantRepo: repo.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, submitted := fakeSkill(t)
			poller, err := client.BeginSkill(t.Context(), "scan", tt.opts...)
			if err != nil {
				t.Fatal(err)
			}
			body := submitted()
			order, ok := body["order"].(map[string]any)
			if !ok {
				t.Fatalf("submit body %+v, want it wrapped in an order", body)
			}
			if order["asset"] != "scan" {
				t.Fatalf("submit body %+v, want the asset by name", order)
			}
			if order["priority"] != float64(1) {
				t.Fatalf("submit body priority %v, want 1", order["priority"])
			}
			if repo := order["repository_id"]; tt.wantRepo == nil {
				if repo != nil {
					t.Fatalf("submit body sent repository_id %v, want none", repo)
				}
			} else if repo != tt.wantRepo {
				t.Fatalf("submit body repository_id %v, want %v", repo, tt.wantRepo)
			}
			for _, absent := range []string{"config", "lane", "rule", "event_source", "run_batch_id", "work_item_id"} {
				if v, has := body[absent]; has && v != nil {
					t.Fatalf("submit body sent %s: %+v", absent, v)
				}
				if v, has := order[absent]; has && v != nil {
					t.Fatalf("submit body sent %s: %+v", absent, v)
				}
			}
			finished, done, err := poller.Poll(t.Context())
			if err != nil || !done || finished.State != traust.RunOrderStateSucceeded || poller.ResumeToken() != finished.Id.String() {
				t.Fatalf("poll order %+v done %v err %v, want the submitted order finished by its id", finished, done, err)
			}
		})
	}
}

func TestBeginSkillRequiresAnAsset(t *testing.T) {
	client, _ := fakeSkill(t)
	if _, err := client.BeginSkill(t.Context(), ""); err == nil {
		t.Fatal("BeginSkill accepted an empty asset")
	}
}

func TestBeginSkillPassesCoreRefusalsThrough(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(traust.Problem{Title: "Forbidden"})
	}))
	t.Cleanup(server.Close)
	client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("machine")})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := runs.New(client).BeginSkill(t.Context(), "scan"); !errors.Is(err, traust.ErrForbidden) {
		t.Fatalf("a refused submit: %v, want ErrForbidden", err)
	}
}
