package work_test

import (
	"encoding/json"
	"math"
	"os"
	"reflect"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/routing/work"
)

func TestPrimaryPythonParity(t *testing.T) {
	var fixture struct {
		EventLanes []struct {
			Source work.EventSource
			Tier   work.RiskTier
			Lane   *work.Lane
		} `json:"event_lanes"`
		Primary []struct {
			Name  string
			Input work.PrimaryInput
			Want  work.PrimaryResult
		}
	}
	raw, err := os.ReadFile("testdata/python_primary.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, &fixture); err != nil {
		t.Fatal(err)
	}
	if len(fixture.EventLanes) == 0 || len(fixture.Primary) == 0 {
		t.Fatal("primary fixture must exercise event lanes and composed selection")
	}
	for _, tc := range fixture.EventLanes {
		t.Run("event-lane/"+string(tc.Source)+"/"+string(tc.Tier), func(t *testing.T) {
			lane, selected := work.EventLane(tc.Source, tc.Tier)
			if selected != (tc.Lane != nil) || (selected && lane != *tc.Lane) || (!selected && lane != "") {
				t.Fatalf("EventLane = %q, %v; Python lane = %v", lane, selected, tc.Lane)
			}
		})
	}
	for _, tc := range fixture.Primary {
		t.Run(tc.Name, func(t *testing.T) {
			before, err := json.Marshal(tc.Input)
			if err != nil {
				t.Fatal(err)
			}
			got, err := work.DecidePrimary(tc.Input)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.Want) {
				t.Fatalf("Go primary selection = %+v; Python selection = %+v", got, tc.Want)
			}
			after, err := json.Marshal(tc.Input)
			if err != nil {
				t.Fatal(err)
			}
			if string(before) != string(after) {
				t.Fatal("selection mutated input facts or consumed an event")
			}
		})
	}
}

func TestInvalidPrimaryInput(t *testing.T) {
	negative := int64(-1)
	for _, tc := range []struct {
		name  string
		field string
		input work.PrimaryInput
	}{
		{"missing status", "Audit.Status", work.PrimaryInput{Audit: &work.AuditInput{}}},
		{"unknown status", "Audit.Status", work.PrimaryInput{Audit: &work.AuditInput{Status: "other"}}},
		{"bootstrap status is output only", "Audit.Status", work.PrimaryInput{Audit: &work.AuditInput{Status: work.StatusNeverAudited}}},
		{"unknown designation", "Inventory.Designation", work.PrimaryInput{Inventory: &work.InventoryInput{Designation: "public"}}},
		{"negative risk", "LiveCriticalHigh", work.PrimaryInput{Audit: &work.AuditInput{
			Status: work.StatusOK, Risk: work.RiskInput{LiveCriticalHigh: -1},
		}}},
		{"unknown exposure", "Exposure", work.PrimaryInput{Audit: &work.AuditInput{
			Status: work.StatusOK, Exposure: "external",
		}}},
		{"negative age", "AuditAgeDays", work.PrimaryInput{Audit: &work.AuditInput{
			Status: work.StatusOK, AuditAgeDays: &negative,
		}}},
		{"negative churn", "ChangedLines", work.PrimaryInput{Audit: &work.AuditInput{
			Status: work.StatusOK, Change: &work.ChangeInput{ChangedLines: -1},
		}}},
		{"negative sensitive lines", "SensitiveChangedLines", work.PrimaryInput{Audit: &work.AuditInput{
			Status: work.StatusOK, Change: &work.ChangeInput{SensitiveChangedLines: -1},
		}}},
		{"negative commits", "CommitsAhead", work.PrimaryInput{Audit: &work.AuditInput{
			Status: work.StatusOK, Change: &work.ChangeInput{CommitsAhead: &negative},
		}}},
		{"missing changed comparison", "Audit.Change", work.PrimaryInput{Audit: &work.AuditInput{
			Status: work.StatusOK, PushChanged: true,
		}}},
		{"unqueued event needs comparison", "Audit.Change", work.PrimaryInput{
			Audit:  &work.AuditInput{Status: work.StatusOK, PushChanged: true},
			Events: []work.RescanEvent{{Source: work.EventMethodology}},
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := work.DecidePrimary(tc.input)
			assertInputError(t, err, tc.field)
			if !reflect.DeepEqual(got, work.PrimaryResult{}) {
				t.Fatalf("invalid input returned a partial result: %+v", got)
			}
		})
	}
	for _, ratio := range []float64{-1, math.NaN(), math.Inf(1), math.Inf(-1)} {
		got, err := work.DecidePrimary(work.PrimaryInput{
			Audit:  &work.AuditInput{Status: work.StatusOK, Change: &work.ChangeInput{CoverageChangeRatio: &ratio}},
			Events: []work.RescanEvent{{Source: work.EventExternalReport}},
		})
		assertInputError(t, err, "CoverageChangeRatio")
		if !reflect.DeepEqual(got, work.PrimaryResult{}) {
			t.Fatalf("invalid input returned an event decision: %+v", got)
		}
	}
}

func TestQueuedEventDoesNotRequireComparison(t *testing.T) {
	got, err := work.DecidePrimary(work.PrimaryInput{
		Audit:  &work.AuditInput{Status: work.StatusOK, PushChanged: true},
		Events: []work.RescanEvent{{Source: work.EventExternalReport}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Decisions) != 1 || got.Decisions[0].Rule != "rule-1" {
		t.Fatalf("event did not supersede comparison: %+v", got)
	}
}
