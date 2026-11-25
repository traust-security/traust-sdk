// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
