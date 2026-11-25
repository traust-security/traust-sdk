// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
