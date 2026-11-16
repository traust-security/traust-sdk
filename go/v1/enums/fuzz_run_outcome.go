// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
