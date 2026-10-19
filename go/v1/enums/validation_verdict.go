// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// ValidationVerdict — Validation step/finding verdict.
type ValidationVerdict string

const (
	ValidationVerdictConfirmed      ValidationVerdict = "confirmed"
	ValidationVerdictRefuted        ValidationVerdict = "refuted"
	ValidationVerdictInconclusive   ValidationVerdict = "inconclusive"
	ValidationVerdictBlockedByScope ValidationVerdict = "blocked_by_scope"
	ValidationVerdictNotAttempted   ValidationVerdict = "not_attempted"
)

// ValidationVerdictValues returns all valid ValidationVerdict values.
func ValidationVerdictValues() []ValidationVerdict {
	return []ValidationVerdict{
		ValidationVerdictConfirmed,
		ValidationVerdictRefuted,
		ValidationVerdictInconclusive,
		ValidationVerdictBlockedByScope,
		ValidationVerdictNotAttempted,
	}
}
