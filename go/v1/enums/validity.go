// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// Validity — Finding validity: confirmed means the claim is real; false_positive means the claim is wrong; not_affected means a real issue does not apply to this repository, version or deployment. Historical values remain readable.
type Validity string

const (
	ValidityConfirmed     Validity = "confirmed"
	ValidityCorrected     Validity = "corrected"
	ValidityFalsePositive Validity = "false_positive"
	ValidityNotVerified   Validity = "not_verified"
	ValidityHardening     Validity = "hardening"
	ValidityNotAffected   Validity = "not_affected"
	ValidityTruePositive  Validity = "true_positive"
)

// ValidityValues returns all valid Validity values.
func ValidityValues() []Validity {
	return []Validity{
		ValidityConfirmed,
		ValidityCorrected,
		ValidityFalsePositive,
		ValidityNotVerified,
		ValidityHardening,
		ValidityNotAffected,
		ValidityTruePositive,
	}
}
