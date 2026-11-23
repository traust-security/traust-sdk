// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
