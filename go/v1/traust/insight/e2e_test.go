//go:build e2e

package insight_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/insight"
)

// e2eClient runs against traust-server built from core at CORE_REF; `make e2e` starts it.
func e2eClient(t *testing.T) *insight.Client {
	t.Helper()
	url, tokens := os.Getenv("TRAUST_E2E_URL"), os.Getenv("TRAUST_E2E_TOKEN_FILE")
	if url == "" || tokens == "" {
		t.Fatal("TRAUST_E2E_URL and TRAUST_E2E_TOKEN_FILE are required; run make e2e")
	}
	c, err := traust.NewClient(traust.Config{BaseURL: url, Tokens: traust.FileTokenSource{Path: tokens}})
	if err != nil {
		t.Fatal(err)
	}
	return insight.New(c)
}

func TestE2EQueryDashboardCountsARepositorySet(t *testing.T) {
	set := []uuid.UUID{uuid.New()}
	result, err := e2eClient(t).QueryDashboard(context.Background(), traust.DashboardQuery{
		Subject: traust.SubjectFindings, Measures: []traust.Measure{traust.MeasureFindings}, RepositoryIds: &set,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range result.Rows {
		if n := insight.Measure(row, traust.MeasureFindings); n != nil && *n != 0 {
			t.Fatalf("a repository the fresh server does not know counted %v findings", *n)
		}
	}
}
