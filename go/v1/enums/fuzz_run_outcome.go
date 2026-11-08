// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
