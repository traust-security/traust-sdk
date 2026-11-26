// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// OwaspImpactBasis — Which factor set determines a threat's impact score under the OWASP Risk Rating Methodology.
type OwaspImpactBasis string

const (
	OwaspImpactBasisTechnical OwaspImpactBasis = "technical"
	OwaspImpactBasisBusiness  OwaspImpactBasis = "business"
)

// OwaspImpactBasisValues returns all valid OwaspImpactBasis values.
func OwaspImpactBasisValues() []OwaspImpactBasis {
	return []OwaspImpactBasis{
		OwaspImpactBasisTechnical,
		OwaspImpactBasisBusiness,
	}
}
