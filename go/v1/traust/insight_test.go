package traust_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"

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
	score := "9.8"
	counts := traust.FindingsRollupInsightSeverityCounts{Critical: 1, Unrated: 1}
	fixture := traust.FindingsRollupInsight{
		Total:           traust.FindingsRollupInsightTally{Open: counts, Unverified: 2, Tracked: 2},
		Products:        []traust.FindingsRollupInsightProduct{},
		Repositories:    []traust.FindingsRollupInsightRepository{},
		Categories:      []traust.FindingsRollupInsightCategory{},
		AttackScenarios: []traust.FindingsRollupInsightAttackScenario{},
		Top: []traust.FindingsRollupInsightTopFinding{{
			FindingId:     uuid.MustParse("00000000-0000-4000-8000-000000000001"),
			RepositoryUrl: "https://forge.example/org/repo",
			Title:         "Unbounded input allocation",
			Severity:      traust.FindingsRollupInsightTopFindingSeverityCritical,
			CvssScore:     &score,
		}},
	}
	core := traustest.New(t)
	core.Handle("GET /v1/insights/findings-rollup", traustest.JSON(http.StatusOK, fixture))
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
		{"sla", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.SlaInsight](ctx, c, nil)
			return err
		}},
		{"census", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.CensusInsight](ctx, c, nil)
			return err
		}},
		{"attack-coverage", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.AttackCoverageInsight](ctx, c, nil)
			return err
		}},
		{"pqc-portfolio", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.PqcPortfolioInsight](ctx, c, nil)
			return err
		}},
		{"threat-register", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.ThreatRegisterInsight](ctx, c, nil)
			return err
		}},
		{"exposure", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.ExposureInsight](ctx, c, nil)
			return err
		}},
		{"executive-summary", func(ctx context.Context, c *traust.Client) error {
			_, err := traust.GetInsightAs[traust.ExecutiveSummaryInsight](ctx, c, nil)
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
		name    string
		status  int
		handler http.Handler
	}{
		{"unknown view", http.StatusNotFound, traustest.Problem(http.StatusNotFound, "refused")},
		{"unsupported filter", http.StatusUnprocessableEntity, traustest.Problem(http.StatusUnprocessableEntity, "refused")},
		{"unauthenticated", http.StatusUnauthorized, traustest.Problem(http.StatusUnauthorized, "refused")},
		{"server failure", http.StatusServiceUnavailable, traustest.Problem(http.StatusServiceUnavailable, "refused")},
		{"unrecognised problem code", http.StatusConflict, problemWithCode(http.StatusConflict, "a-code-this-sdk-predates")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := traustest.New(t)
			core.Handle("GET /v1/insights/coverage", tt.handler)
			view, err := traust.GetInsightAs[traust.CoverageInsight](context.Background(), core.Client(), nil)
			var problem *traust.ProblemError
			if view != nil || !errors.As(err, &problem) || problem.StatusCode != tt.status {
				t.Fatalf("view = %v, err = %v; want no view and a ProblemError with status %d", view, err, tt.status)
			}
			if !strings.Contains(err.Error(), "coverage insight") {
				t.Fatalf("err = %v; want it to name the view", err)
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

func problemWithCode(status int, code traust.DenyCode) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(traust.Problem{
			Type: "about:blank", Title: http.StatusText(status), Status: status, Code: &code,
			Detail: "refused", Instance: r.URL.Path, RequestId: "traustest",
		})
	})
}

func is[T any](v T) T { return v }

// The executive summary composes the component views: its panels decode, strictly, into
// the components' own types (testdata is a document valid against the contracts).
func TestTheExecutiveSummaryDecodesIntoItsComponentViews(t *testing.T) {
	body, err := os.ReadFile("testdata/executive-summary.json")
	if err != nil {
		t.Fatal(err)
	}
	core := traustest.New(t)
	core.Handle("GET /v1/insights/executive-summary", traustest.JSON(http.StatusOK, json.RawMessage(body)))
	summary, err := traust.GetInsightAs[traust.ExecutiveSummaryInsight](context.Background(), core.Client(), nil)
	if err != nil {
		t.Fatalf("GetInsightAs: %v", err)
	}
	// each panel is the component view's own type (this does not compile on a copy)
	rollup := is[traust.FindingsRollupInsight](summary.Rollup)
	attack := is[*traust.AttackCoverageInsight](summary.Attack)
	pqc := is[*traust.PqcPortfolioInsight](summary.Pqc)
	if rollup.Total.Open.High != 1 || attack == nil || attack.AttackVersion != "19.1" {
		t.Fatalf("rollup/attack = %+v / %+v", rollup.Total, attack)
	}
	if pqc == nil || len(pqc.Scopes) != 1 || summary.PqcUnavailable != nil {
		t.Fatalf("pqc = %+v, unavailable = %v", pqc, summary.PqcUnavailable)
	}
	register := is[traust.ThreatRegisterInsight](summary.Isolation.ThreatRegister)
	layer := is[traust.CensusInsightLayer](summary.Iac.Findings)
	if register.Models != 0 || layer.Findings.Low != 1 {
		t.Fatalf("register/layer = %+v / %+v", register, layer)
	}
	rows := is[*[]traust.ExposureInsightRollup](summary.Exposure)
	if rows == nil || len(*rows) != 1 || *(*rows)[0].Key != "GHSA-xxxx-xxxx-xxxx" {
		t.Fatalf("exposure = %+v", rows)
	}
}
