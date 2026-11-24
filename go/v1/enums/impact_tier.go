// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
