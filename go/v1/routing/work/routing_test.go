package work_test

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/routing/work"
)

// These outputs were produced by the pinned Python rules, not by the Go port.
// Normal Go tests read the fixture offline and do not require Python or Traust.
func TestPythonParity(t *testing.T) {
	var fixture struct {
		Table []struct {
			Name  string
			Input struct {
				Tier         work.RiskTier `json:"tier"`
				Exposure     work.Exposure `json:"exposure"`
				C            int64
				R            *float64
				S            bool
				Sensitive    int64  `json:"S_lines"`
				DepsOnly     bool   `json:"deps_only"`
				AuditAgeDays *int64 `json:"audit_age_days"`
				Changed      bool   `json:"changed"`
				AheadBy      *int64 `json:"ahead_by"`
			}
			Want work.Decision
		}
		Risk []struct {
			Name  string
			Input work.RiskInput
			Want  work.RiskTier
		}
	}
	raw, err := os.ReadFile("testdata/python_decisions.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, &fixture); err != nil {
		t.Fatal(err)
	}
	if len(fixture.Table) == 0 || len(fixture.Risk) == 0 {
		t.Fatal("parity fixture must cover table and risk rules")
	}
	for _, tc := range fixture.Table {
		t.Run(tc.Name, func(t *testing.T) {
			got, err := work.EvaluateTable(work.TableInput{
				RiskTier:              tc.Input.Tier,
				Exposure:              tc.Input.Exposure,
				ChangedLines:          tc.Input.C,
				CoverageChangeRatio:   tc.Input.R,
				SensitiveChange:       tc.Input.S,
				SensitiveChangedLines: tc.Input.Sensitive,
				DependenciesOnly:      tc.Input.DepsOnly,
				AuditAgeDays:          tc.Input.AuditAgeDays,
				PushChanged:           tc.Input.Changed,
				CommitsAhead:          tc.Input.AheadBy,
			})
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.Want {
				t.Fatalf("Go decision = %+v; Python decision = %+v", got, tc.Want)
			}
		})
	}
	for _, tc := range fixture.Risk {
		t.Run(tc.Name, func(t *testing.T) {
			got, err := work.ClassifyRisk(tc.Input)
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.Want {
				t.Fatalf("Go risk tier = %s; Python risk tier = %s", got, tc.Want)
			}
		})
	}
}

func TestInvalidTableInput(t *testing.T) {
	negative := int64(-1)
	for _, tc := range []struct {
		name  string
		field string
		input work.TableInput
	}{
		{"missing risk", "RiskTier", work.TableInput{}},
		{"unknown risk", "RiskTier", work.TableInput{RiskTier: "P4"}},
		{"unknown exposure", "Exposure", work.TableInput{RiskTier: work.RiskP1, Exposure: "external"}},
		{"negative churn", "ChangedLines", work.TableInput{RiskTier: work.RiskP1, ChangedLines: -1}},
		{"negative sensitive lines", "SensitiveChangedLines", work.TableInput{RiskTier: work.RiskP1, SensitiveChangedLines: -1}},
		{"negative age", "AuditAgeDays", work.TableInput{RiskTier: work.RiskP1, AuditAgeDays: &negative}},
		{"negative ahead", "CommitsAhead", work.TableInput{RiskTier: work.RiskP1, CommitsAhead: &negative}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := work.EvaluateTable(tc.input)
			assertInputError(t, err, tc.field)
			if got != (work.Decision{}) {
				t.Fatalf("invalid input returned a usable decision: %+v", got)
			}
		})
	}
	for _, ratio := range []float64{-1, math.NaN(), math.Inf(1), math.Inf(-1)} {
		// Even an otherwise matching rule must not hide invalid input.
		got, err := work.EvaluateTable(work.TableInput{
			RiskTier: work.RiskP0, ChangedLines: 9000, CoverageChangeRatio: &ratio,
		})
		assertInputError(t, err, "CoverageChangeRatio")
		if got != (work.Decision{}) {
			t.Fatalf("invalid ratio returned a usable decision: %+v", got)
		}
	}
}

func TestNegativeLiveCount(t *testing.T) {
	got, err := work.ClassifyRisk(work.RiskInput{LiveCriticalHigh: -1})
	assertInputError(t, err, "LiveCriticalHigh")
	if got != "" {
		t.Fatalf("invalid count returned a risk tier: %s", got)
	}
}

func assertInputError(t *testing.T, err error, field string) {
	t.Helper()
	var inputErr *work.InputError
	if !errors.As(err, &inputErr) {
		t.Fatalf("wanted InputError for %s, got %v", field, err)
	}
	if inputErr.Field != field || inputErr.Detail == "" || inputErr.Error() == "" {
		t.Fatalf("unexpected input error: %+v", inputErr)
	}
}
