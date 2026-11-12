//go:build e2e

package traust

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
)

// Runs against traust-server built from core at CORE_REF; `make e2e` starts it.
func e2eClient(t *testing.T) *Client {
	t.Helper()
	url, tokens := os.Getenv("TRAUST_E2E_URL"), os.Getenv("TRAUST_E2E_TOKEN_FILE")
	if url == "" || tokens == "" {
		t.Fatal("TRAUST_E2E_URL and TRAUST_E2E_TOKEN_FILE are required; run make e2e")
	}
	c, err := NewClient(Config{BaseURL: url, Tokens: FileTokenSource{Path: tokens}})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestE2ECallerIsVerified(t *testing.T) {
	resp, err := e2eClient(t).WhoamiWithResponse(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if resp.JSON200 == nil || resp.JSON200.Identity == "" {
		t.Fatalf("whoami: %d %s", resp.StatusCode(), resp.Body)
	}
}

func TestE2EPagingDecodesAnEmptyInventory(t *testing.T) {
	var seen int
	for _, err := range e2eClient(t).ListRepositoriesAll(context.Background(), nil) {
		if err != nil {
			t.Fatal(err)
		}
		seen++
	}
	if seen != 0 {
		t.Fatalf("fresh server listed %d repositories", seen)
	}
}

func TestE2EUnknownResourceIsATypedProblem(t *testing.T) {
	resp, err := e2eClient(t).GetRepositoryWithResponse(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	var problem *ProblemError
	_, err = Decode(resp.HTTPResponse, resp.Body, resp.JSON200)
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusNotFound || problem.Problem == nil || problem.Problem.RequestId == "" {
		t.Fatalf("want a 404 problem with a request id, got %v", err)
	}
}

func TestE2EUnverifiedTokenIsUnauthorized(t *testing.T) {
	tokens := t.TempDir() + "/token"
	if err := os.WriteFile(tokens, []byte("not-a-token"), 0o600); err != nil {
		t.Fatal(err)
	}
	c, err := NewClient(Config{BaseURL: os.Getenv("TRAUST_E2E_URL"), Tokens: FileTokenSource{Path: tokens}})
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.WhoamiWithResponse(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var problem *ProblemError
	_, err = Decode(resp.HTTPResponse, resp.Body, resp.JSON200)
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnauthorized || problem.Problem == nil {
		t.Fatalf("want a 401 problem, got %v", err)
	}
}

func TestE2EBatchFindingsListingRefusesWithoutRepositories(t *testing.T) {
	resp, err := e2eClient(t).ListFindingsInWithResponse(context.Background(), &ListFindingsInParams{})
	if err != nil {
		t.Fatal(err)
	}
	var problem *ProblemError
	_, err = Decode(resp.HTTPResponse, resp.Body, resp.JSON200)
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity || problem.Problem == nil || problem.Problem.RequestId == "" {
		t.Fatalf("want a 422 problem with a request id, got %v", err)
	}
}

func TestE2ECountFindingsInCountsARepositorySet(t *testing.T) {
	set := []uuid.UUID{uuid.New()}
	counted, err := e2eClient(t).CountFindingsIn(context.Background(), &CountFindingsInParams{
		RepositoryIds: &set,
	})
	if err != nil {
		t.Fatal(err)
	}
	if counted == nil || counted.Total != 0 {
		t.Fatalf("a repository the fresh server does not know counted %v", counted)
	}
}

func TestE2ECountFindingsInRefusesWithoutRepositories(t *testing.T) {
	resp, err := e2eClient(t).CountFindingsInWithResponse(context.Background(), &CountFindingsInParams{})
	if err != nil {
		t.Fatal(err)
	}
	var problem *ProblemError
	_, err = Decode(resp.HTTPResponse, resp.Body, resp.JSON200)
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity || problem.Problem == nil || problem.Problem.RequestId == "" {
		t.Fatalf("want a 422 problem with a request id, got %v", err)
	}
}

func TestE2EBatchFindingsListingPagesARepositorySet(t *testing.T) {
	set := []uuid.UUID{uuid.New()}
	var seen int
	for _, err := range e2eClient(t).ListFindingsInAll(context.Background(), &ListFindingsInParams{
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

// The wave-f reads refuse without a named set exactly as the batch listing and count do:
// a 422 problem with a request id, never a firehose of everything the server holds.
func TestE2ESetScopedReadsRefuseWithoutRepositories(t *testing.T) {
	client := e2eClient(t)
	reads := []struct {
		name string
		read func() error
	}{
		{"assessments", func() error {
			_, err := client.ListAssessmentsIn(context.Background(), &ListAssessmentsInParams{})
			return err
		}},
		{"refs", func() error {
			_, err := client.ListRefsIn(context.Background(), &ListRefsInParams{})
			return err
		}},
		{"summaries", func() error {
			_, err := client.SummarizeFindingsIn(context.Background(), &SummarizeFindingsInParams{})
			return err
		}},
	}
	for _, read := range reads {
		t.Run(read.name, func(t *testing.T) {
			err := read.read()
			var problem *ProblemError
			if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity || problem.Problem == nil || problem.Problem.RequestId == "" {
				t.Fatalf("want a 422 problem with a request id, got %v", err)
			}
		})
	}
}

// The wave-f reads over a repository set the fresh server does not know decode as an empty
// page: the refusal, parse and decode branches all take their success paths.
func TestE2ESetScopedReadsListARepositorySet(t *testing.T) {
	client := e2eClient(t)
	set := []uuid.UUID{uuid.New()}
	reads := []struct {
		name string
		read func() (int, error)
	}{
		{"assessments", func() (int, error) {
			found, err := client.ListAssessmentsIn(context.Background(), &ListAssessmentsInParams{RepositoryIds: &set})
			return len(found), err
		}},
		{"refs", func() (int, error) {
			found, err := client.ListRefsIn(context.Background(), &ListRefsInParams{RepositoryIds: &set})
			return len(found), err
		}},
		{"summaries", func() (int, error) {
			found, err := client.SummarizeFindingsIn(context.Background(), &SummarizeFindingsInParams{RepositoryIds: &set})
			return len(found), err
		}},
	}
	for _, read := range reads {
		t.Run(read.name, func(t *testing.T) {
			seen, err := read.read()
			if err != nil {
				t.Fatal(err)
			}
			if seen != 0 {
				t.Fatalf("a repository the fresh server does not know listed %d rows", seen)
			}
		})
	}
}
