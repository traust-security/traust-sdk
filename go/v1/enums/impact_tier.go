// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// ImpactTier — Which analysis tiers were executed.
type ImpactTier string

const (
	ImpactTierL1             ImpactTier = "L1"
	ImpactTierL4             ImpactTier = "L4"
	ImpactTierGovulncheck    ImpactTier = "govulncheck"
	ImpactTierFeaturePattern ImpactTier = "feature_pattern"
	ImpactTierBinaryScan     ImpactTier = "binary_scan"
	ImpactTierManifestScan   ImpactTier = "manifest_scan"
	ImpactTierSourceScan     ImpactTier = "source_scan"
	ImpactTierSbomCrossCheck ImpactTier = "sbom_cross_check"
)

// ImpactTierValues returns all valid ImpactTier values.
func ImpactTierValues() []ImpactTier {
	return []ImpactTier{
		ImpactTierL1,
		ImpactTierL4,
		ImpactTierGovulncheck,
		ImpactTierFeaturePattern,
		ImpactTierBinaryScan,
		ImpactTierManifestScan,
		ImpactTierSourceScan,
		ImpactTierSbomCrossCheck,
	}
}
