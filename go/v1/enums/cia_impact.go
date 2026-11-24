// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
