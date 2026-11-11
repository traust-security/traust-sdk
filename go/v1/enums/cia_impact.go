// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
