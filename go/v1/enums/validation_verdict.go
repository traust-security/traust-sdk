// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
