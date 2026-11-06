// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
