// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// EventValidity — Finding validity for disposition events. 4-value set (not_verified is the audit-stage default and is never set by an event). 'hardening' = defense-in-depth gap with no concrete exploit path.
type EventValidity string

const (
	EventValidityConfirmed     EventValidity = "confirmed"
	EventValidityFalsePositive EventValidity = "false_positive"
	EventValidityCorrected     EventValidity = "corrected"
	EventValidityHardening     EventValidity = "hardening"
)

// EventValidityValues returns all valid EventValidity values.
func EventValidityValues() []EventValidity {
	return []EventValidity{
		EventValidityConfirmed,
		EventValidityFalsePositive,
		EventValidityCorrected,
		EventValidityHardening,
	}
}
