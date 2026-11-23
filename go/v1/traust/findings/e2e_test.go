//go:build e2e

package findings_test

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/findings"
)

// e2eCore runs against traust-server built from core at CORE_REF; `make e2e` starts it.
func e2eCore(t *testing.T) *traust.Client {
	t.Helper()
	url, tokens := os.Getenv("TRAUST_E2E_URL"), os.Getenv("TRAUST_E2E_TOKEN_FILE")
	if url == "" || tokens == "" {
		t.Fatal("TRAUST_E2E_URL and TRAUST_E2E_TOKEN_FILE are required; run make e2e")
	}
	c, err := traust.NewClient(traust.Config{BaseURL: url, Tokens: traust.FileTokenSource{Path: tokens}})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestE2EBatchFindingsListingRefusesWithoutRepositories(t *testing.T) {
	resp, err := e2eCore(t).ListFindingsInWithResponse(context.Background(), &traust.ListFindingsInParams{})
	if err != nil {
		t.Fatal(err)
	}
	var problem *traust.ProblemError
	_, err = traust.Decode(resp.HTTPResponse, resp.Body, resp.JSON200)
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity || problem.Problem == nil || problem.Problem.RequestId == "" {
		t.Fatalf("want a 422 problem with a request id, got %v", err)
	}
}

func TestE2EBatchFindingsListingPagesARepositorySet(t *testing.T) {
	set := []uuid.UUID{uuid.New()}
	var seen int
	for _, err := range e2eClient(t).ListFindingsInAll(context.Background(), &traust.ListFindingsInParams{
		RepositoryIds: &set,
	}) {
		if err != nil {
			t.Fatal(err)
		}
		seen++
	}
	if seen != 0 {
		t.Fatalf("a repository the fresh server does not know listed %d findings", seen)
	}
}

// The wave-f assessments read refuses without a named set exactly as the batch listing does:
// a 422 problem with a request id, never a firehose of everything the server holds.
func TestE2ESetScopedReadsRefuseWithoutRepositories(t *testing.T) {
	client := e2eClient(t)
	_, err := client.ListAssessments(context.Background(), &traust.ListAssessmentsParams{})
	var problem *traust.ProblemError
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity || problem.Problem == nil || problem.Problem.RequestId == "" {
		t.Fatalf("want a 422 problem with a request id, got %v", err)
	}
}

// The wave-f assessments read over a repository set the fresh server does not know decodes as
// an empty page: the refusal, parse and decode branches all take their success paths.
func TestE2ESetScopedReadsListARepositorySet(t *testing.T) {
	client := e2eClient(t)
	set := []uuid.UUID{uuid.New()}
	found, err := client.ListAssessments(context.Background(), &traust.ListAssessmentsParams{RepositoryIds: &set})
	if err != nil {
		t.Fatal(err)
	}
	if found == nil || len(found.Items) != 0 {
		t.Fatalf("a repository the fresh server does not know listed %d rows", len(found.Items))
	}
}
