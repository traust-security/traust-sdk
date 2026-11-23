// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
