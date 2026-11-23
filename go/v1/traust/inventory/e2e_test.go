//go:build e2e

package inventory_test

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/inventory"
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

// e2eClient is the inventory calls over e2eCore.
func e2eClient(t *testing.T) *inventory.Client {
	t.Helper()
	return inventory.New(e2eCore(t))
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
	resp, err := e2eCore(t).GetRepositoryWithResponse(context.Background(), uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	var problem *traust.ProblemError
	_, err = traust.Decode(resp.HTTPResponse, resp.Body, resp.JSON200)
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusNotFound || problem.Problem == nil || problem.Problem.RequestId == "" {
		t.Fatalf("want a 404 problem with a request id, got %v", err)
	}
}

// The wave-f reads refuse without a named set exactly as the batch findings listing does:
// a 422 problem with a request id, never a firehose of everything the server holds.
func TestE2ESetScopedReadsRefuseWithoutRepositories(t *testing.T) {
	client := e2eClient(t)
	_, err := client.ListRefsForRepositories(context.Background(), &traust.ListRefsForRepositoriesParams{})
	var problem *traust.ProblemError
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity || problem.Problem == nil || problem.Problem.RequestId == "" {
		t.Fatalf("want a 422 problem with a request id, got %v", err)
	}
}

// The wave-f reads over a repository set the fresh server does not know decode as an empty
// listing: the refusal, parse and decode branches all take their success paths.
func TestE2ESetScopedReadsListARepositorySet(t *testing.T) {
	client := e2eClient(t)
	set := []uuid.UUID{uuid.New()}
	found, err := client.ListRefsForRepositories(context.Background(), &traust.ListRefsForRepositoriesParams{RepositoryIds: &set})
	if err != nil {
		t.Fatal(err)
	}
	if len(found) != 0 {
		t.Fatalf("a repository the fresh server does not know listed %d rows", len(found))
	}
}
