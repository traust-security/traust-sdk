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
