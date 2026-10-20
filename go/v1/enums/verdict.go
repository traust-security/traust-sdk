// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
