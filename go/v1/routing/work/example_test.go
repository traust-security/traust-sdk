package work_test

import (
	"fmt"

	"github.com/traust-security/traust-sdk/go/v1/routing/work"
)

func ExampleEvaluateTable() {
	// The controller has already collected a usable baseline comparison and
	// handled bootstrap/events. It supplies normalized measurements here.
	age := int64(30)
	decision, err := work.EvaluateTable(work.TableInput{
		RiskTier:              work.RiskP1,
		Exposure:              work.ExposurePrivateInternal,
		ChangedLines:          500,
		SensitiveChange:       true,
		SensitiveChangedLines: 200,
		AuditAgeDays:          &age,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decision.Rule, decision.Lane)
	// The caller still assembles the worklist and applies execution gates.
	// Output: rule-3 full-audit
}

func ExampleClassifyRisk() {
	tier, err := work.ClassifyRisk(work.RiskInput{
		LiveCriticalHigh: 5,
		Archived:         true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tier)
	// Live risk takes precedence over archival state.
	// Output: P0
}

func ExampleDecidePrimary() {
	result, err := work.DecidePrimary(work.PrimaryInput{
		// Known inventory membership, with no HEAD code-audit baseline.
		Inventory: &work.InventoryInput{Designation: work.DesignationExternal},
		Events:    []work.RescanEvent{{Source: work.EventCVE}},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// The event has no audited target. Bootstrap still selects the first audit.
	fmt.Println(result.EventsHonored[0].Queued)
	fmt.Println(result.Decisions[0].Rule, result.Decisions[0].Lane)
	fmt.Println(result.Decisions[0].RiskTier, result.Decisions[0].Exposure)
	// Output:
	// false
	// rule-1-bootstrap full-audit
	// P2 private-external
}
