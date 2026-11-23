package inventory_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/inventory"
)

func readAll(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

func of[T any](value T) *T { return &value }

// The repository disposition doors (foundry plan repo-unreachable): the URL correction
// and the verify reach core's routes with the typed bodies, the listing carries the
// dispositioned filter, and the answer decodes the disposition fields core serves.
func TestRepositoryDispositionDoorsRoundTrip(t *testing.T) {
	var seen []struct {
		method string
		path   string
		query  url.Values
		body   []byte
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := readAll(r)
		seen = append(seen, struct {
			method string
			path   string
			query  url.Values
			body   []byte
		}{r.Method, r.URL.Path, r.URL.Query(), body})
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/v1/repositories" && r.Method == http.MethodGet:
			_ = json.NewEncoder(w).Encode(traust.PageRepository{Items: []traust.Repository{{
				Id:         of(uuid.New()),
				Url:        "https://github.com/org/repo",
				Archived:   of(false),
				UpstreamId: of(uuid.New()),
				Public:     true,
				CreatedAt:  of(time.Date(2026, 11, 15, 12, 0, 0, 0, time.UTC)),
				GoneStreak: of(2),
			}}})
		case strings.HasSuffix(r.URL.Path, "/url") && r.Method == http.MethodPut:
			disposition := traust.RepoUnreachable
			reason := traust.GoneReasonGone
			evidence := "the forge reports no such repository"
			_ = json.NewEncoder(w).Encode(traust.Repository{
				Id: of(uuid.New()), Url: "https://github.com/org/relocated",
				Disposition: &disposition, DispositionReason: &reason, DispositionEvidence: &evidence,
			})
		case strings.HasSuffix(r.URL.Path, "/verify") && r.Method == http.MethodPost:
			status := traust.ComparisonStatusOk
			_ = json.NewEncoder(w).Encode(traust.RepositoryVerification{
				Facts:      traust.ForgeFacts{Status: status},
				Repository: traust.Repository{Id: of(uuid.New()), Url: "https://github.com/org/repo"},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	core, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("t")})
	if err != nil {
		t.Fatal(err)
	}
	client := inventory.New(core)
	ctx := context.Background()

	repositoryID := uuid.New()
	corrected, err := client.CorrectRepositoryUrl(ctx, repositoryID, traust.UrlCorrection{
		Url: "https://github.com/org/relocated",
	})
	if err != nil {
		t.Fatal(err)
	}
	if corrected.Url != "https://github.com/org/relocated" || corrected.Disposition == nil ||
		*corrected.Disposition != traust.RepoUnreachable || *corrected.DispositionReason != traust.GoneReasonGone {
		t.Fatalf("the correction's answer drifted: %+v", corrected)
	}

	answer, err := client.VerifyRepository(ctx, repositoryID)
	if err != nil {
		t.Fatal(err)
	}
	if answer.Facts.Status != traust.ComparisonStatusOk || answer.Repository.Id == nil {
		t.Fatalf("the verify's answer drifted: %+v", answer)
	}

	active := true
	paged, err := client.ListRepositories(ctx, &traust.ListRepositoriesParams{Dispositioned: &active})
	if err != nil {
		t.Fatal(err)
	}
	if len(paged.Items) != 1 || paged.Items[0].Url != "https://github.com/org/repo" {
		t.Fatalf("the listing's answer drifted: %+v", paged)
	}

	if len(seen) != 3 {
		t.Fatalf("calls %d, want the three doors", len(seen))
	}
	for _, call := range seen {
		switch {
		case call.method == http.MethodPut:
			if call.path != "/v1/repositories/"+repositoryID.String()+"/url" ||
				string(call.body) != `{"url":"https://github.com/org/relocated"}` {
				t.Fatalf("the correction was sent wrong: %s %s %s", call.method, call.path, call.body)
			}
		case call.method == http.MethodPost:
			if call.path != "/v1/repositories/"+repositoryID.String()+"/verify" || len(call.body) != 0 {
				t.Fatalf("the verify was sent wrong: %s %s %s", call.method, call.path, call.body)
			}
		case call.method == http.MethodGet:
			if call.path != "/v1/repositories" || call.query.Get("dispositioned") != "true" {
				t.Fatalf("the listing was sent wrong: %s %s %s", call.method, call.path, call.query)
			}
		default:
			t.Fatalf("an unexpected door was called: %s %s", call.method, call.path)
		}
	}
}

// The doors' failure paths (Qodo #77): a refusal surfaces as the *ProblemError core
// sent, wrapped with the call's name and the repository it concerns — a caller
// triaging several repositories can tell which one failed.
func TestRepositoryDispositionDoorsNameTheFailedRepository(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusConflict)
		_ = json.NewEncoder(w).Encode(traust.Problem{
			Type: "about:blank", Title: "the URL is occupied", Detail: "already another repository",
			Status: http.StatusConflict, RequestId: "r1",
		})
	}))
	defer server.Close()

	client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.StaticToken("t")})
	if err != nil {
		t.Fatal(err)
	}
	repositoryID := uuid.New()
	_, err = inventory.New(client).CorrectRepositoryUrl(context.Background(), repositoryID, traust.UrlCorrection{
		Url: "https://github.com/other/seat",
	})
	if err == nil {
		t.Fatal("the occupied-URL refusal must surface")
	}
	var problem *traust.ProblemError
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusConflict {
		t.Fatalf("the refusal lost its problem shape: %v", err)
	}
	if !strings.Contains(err.Error(), "CorrectRepositoryUrl("+repositoryID.String()) {
		t.Fatalf("the error does not name the repository it concerns: %v", err)
	}
	if !strings.Contains(err.Error(), "already another repository") {
		t.Fatalf("the error lost core's words: %v", err)
	}
}
