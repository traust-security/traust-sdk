// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// CiaImpact — Impact axes actually evidenced (what was read/written/disrupted).
type CiaImpact string

const (
	CiaImpactConfidentiality CiaImpact = "confidentiality"
	CiaImpactIntegrity       CiaImpact = "integrity"
	CiaImpactAvailability    CiaImpact = "availability"
)

// CiaImpactValues returns all valid CiaImpact values.
func CiaImpactValues() []CiaImpact {
	return []CiaImpact{
		CiaImpactConfidentiality,
		CiaImpactIntegrity,
		CiaImpactAvailability,
	}
}
