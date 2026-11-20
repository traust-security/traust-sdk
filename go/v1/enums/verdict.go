// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// Verdict — Triage finding verdict taxonomy.
type Verdict string

const (
	VerdictTruePositive  Verdict = "true_positive"
	VerdictHardening     Verdict = "hardening"
	VerdictUndetermined  Verdict = "undetermined"
	VerdictFalsePositive Verdict = "false_positive"
	VerdictDuplicate     Verdict = "duplicate"
)

// VerdictValues returns all valid Verdict values.
func VerdictValues() []Verdict {
	return []Verdict{
		VerdictTruePositive,
		VerdictHardening,
		VerdictUndetermined,
		VerdictFalsePositive,
		VerdictDuplicate,
	}
}
