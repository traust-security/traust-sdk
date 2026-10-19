// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
