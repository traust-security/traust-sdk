// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// FuzzRunOutcome — How one harness's fuzz run ended.
type FuzzRunOutcome string

const (
	FuzzRunOutcomeClean   FuzzRunOutcome = "clean"
	FuzzRunOutcomeCrashed FuzzRunOutcome = "crashed"
	FuzzRunOutcomeFailed  FuzzRunOutcome = "failed"
	FuzzRunOutcomeNotRun  FuzzRunOutcome = "not_run"
)

// FuzzRunOutcomeValues returns all valid FuzzRunOutcome values.
func FuzzRunOutcomeValues() []FuzzRunOutcome {
	return []FuzzRunOutcome{
		FuzzRunOutcomeClean,
		FuzzRunOutcomeCrashed,
		FuzzRunOutcomeFailed,
		FuzzRunOutcomeNotRun,
	}
}
