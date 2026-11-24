// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
