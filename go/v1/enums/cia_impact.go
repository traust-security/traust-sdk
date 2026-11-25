// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
