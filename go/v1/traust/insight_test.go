package traust_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

func rollup() map[string]any {
	zero := map[string]int{"critical": 0, "high": 0, "medium": 0, "low": 0, "informational": 0, "unrated": 0}
	open := map[string]int{"critical": 1, "high": 0, "medium": 0, "low": 0, "informational": 0, "unrated": 1}
	return map[string]any{
		"as_of": nil,
		"total": map[string]any{
			"open": open, "hardening": 0, "resolved": 0, "risk_accepted": 0, "false_positive": 0,
			"execution_proven": 0, "unverified": 2, "decided": 0, "remediated": zero, "tracked": 2,
		},
		"products":     []any{},
		"repositories": []any{},
		"top": []any{map[string]any{
			"finding_id":       "00000000-0000-4000-8000-000000000001",
			"repository_url":   "https://forge.example/org/repo",
			"title":            "Unbounded input allocation",
			"severity":         "critical",
			"cvss_score":       "9.8",
			"execution_proven": false,
			"sla_due":          nil,
		}},
		"categories":       []any{},
		"attack_scenarios": []any{},
	}
}

func TestGetInsightAsAsksForTheViewItsTypeDecodes(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/insights/findings-rollup", traustest.JSON(http.StatusOK, rollup()))
	view, err := traust.GetInsightAs[traust.FindingsRollupInsight](context.Background(), core.Client(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if view.Total.Open.Critical != 1 || view.Total.Tracked != 2 || view.AsOf != nil {
		t.Fatalf("total = %+v, as_of = %v", view.Total, view.AsOf)
	}
	top := view.Top[0]
	if top.Severity != traust.FindingsRollupInsightTopFindingSeverityCritical || top.CvssScore == nil || *top.CvssScore != "9.8" || top.SlaDue != nil {
		t.Fatalf("top = %+v", top)
	}
}

func TestEveryContractedViewDecodesFromItsOwnRoute(t *testing.T) {
	tests := []struct {
		route string
		get   func(context.Context, *traust.Client) error
	}{
		{"findings-rollup", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.FindingsRollupInsight](ctx, c, nil)
			return err
		}},
		{"findings-trend", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.FindingsTrendInsight](ctx, c, nil)
			return err
		}},
		{"coverage", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.CoverageInsight](ctx, c, nil)
			return err
		}},
		{"remediation-index", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.RemediationIndexInsight](ctx, c, nil)
			return err
		}},
		{"critical-misses", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.CriticalMissesInsight](ctx, c, nil)
			return err
		}},
		{"pqc-worklist", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.PqcWorklistInsight](ctx, c, nil)
			return err
		}},
		{"compliance-precision", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.CompliancePrecisionInsight](ctx, c, nil)
			return err
		}},
		{"benchmark-scorecards", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.BenchmarkScorecardsInsight](ctx, c, nil)
			return err
		}},
		{"inventory-export", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.InventoryExportInsight](ctx, c, nil)
			return err
		}},
	}
	for _, tt := range tests {
		t.Run(tt.route, func(t *testing.T) {
			core := traustest.New(t)
			core.Handle("GET /v1/insights/"+tt.route, traustest.JSON(http.StatusOK, map[string]any{}))
			if err := tt.get(context.Background(), core.Client()); err != nil {
				t.Fatalf("GET /v1/insights/%s: %v", tt.route, err)
			}
			requests := core.Requests()
			if len(requests) != 1 || requests[0].Path != "/v1/insights/"+tt.route {
				t.Fatalf("requests = %+v", requests)
			}
		})
	}
}

func TestGetInsightAsRefusesAFieldOutsideTheContract(t *testing.T) {
	body := rollup()
	body["unexpected"] = 1
	core := traustest.New(t)
	core.Handle("GET /v1/insights/findings-rollup", traustest.JSON(http.StatusOK, body))
	view, err := traust.GetInsightAs[traust.FindingsRollupInsight](context.Background(), core.Client(), nil)
	if err == nil || view != nil {
		t.Fatalf("view = %+v, err = %v; want an error and no view", view, err)
	}
}

func TestGetInsightAsRefusesAMalformedBody(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/insights/findings-rollup", traustest.JSON(http.StatusOK, map[string]any{"total": "many"}))
	if _, err := traust.GetInsightAs[traust.FindingsRollupInsight](context.Background(), core.Client(), nil); err == nil {
		t.Fatal("want a decode error")
	}
}

func TestGetInsightAsReturnsCoresRefusalAsAProblem(t *testing.T) {
	tests := []struct {
		name   string
		status int
	}{
		{"unknown view", http.StatusNotFound},
		{"unsupported filter", http.StatusUnprocessableEntity},
		{"unauthenticated", http.StatusUnauthorized},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := traustest.New(t)
			core.Handle("GET /v1/insights/coverage", traustest.Problem(tt.status, "refused"))
			_, err := traust.GetInsightAs[traust.CoverageInsight](context.Background(), core.Client(), nil)
			var problem *traust.ProblemError
			if !errors.As(err, &problem) || problem.StatusCode != tt.status {
				t.Fatalf("err = %v; want a ProblemError with status %d", err, tt.status)
			}
		})
	}
}

func TestGetInsightAsHonoursCancellation(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/insights/coverage", traustest.JSON(http.StatusOK, map[string]any{}))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := traust.GetInsightAs[traust.CoverageInsight](ctx, core.Client(), nil); !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v; want context.Canceled", err)
	}
}
