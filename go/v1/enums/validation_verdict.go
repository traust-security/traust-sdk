// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
