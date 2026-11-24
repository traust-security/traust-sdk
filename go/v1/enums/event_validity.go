// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package enums

// EventValidity — Validity on historical layer events, including canonical confirmed, false_positive and not_affected. Finding kind and restatement operations are separate concepts.
type EventValidity string

const (
	EventValidityConfirmed     EventValidity = "confirmed"
	EventValidityFalsePositive EventValidity = "false_positive"
	EventValidityCorrected     EventValidity = "corrected"
	EventValidityHardening     EventValidity = "hardening"
	EventValidityNotAffected   EventValidity = "not_affected"
)

// EventValidityValues returns all valid EventValidity values.
func EventValidityValues() []EventValidity {
	return []EventValidity{
		EventValidityConfirmed,
		EventValidityFalsePositive,
		EventValidityCorrected,
		EventValidityHardening,
		EventValidityNotAffected,
	}
}
