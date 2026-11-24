// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
