// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
