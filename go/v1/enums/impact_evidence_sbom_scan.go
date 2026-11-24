// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package enums

// ImpactEvidenceSbomScan — Cross-check against syft SBOMs of shipped images: what the delivered artifact actually contains.
type ImpactEvidenceSbomScan string

const (
	ImpactEvidenceSbomScanShippedInRange    ImpactEvidenceSbomScan = "shipped_in_range"
	ImpactEvidenceSbomScanShippedOutOfRange ImpactEvidenceSbomScan = "shipped_out_of_range"
	ImpactEvidenceSbomScanModuleNotInSbom   ImpactEvidenceSbomScan = "module_not_in_sbom"
)

// ImpactEvidenceSbomScanValues returns all valid ImpactEvidenceSbomScan values.
func ImpactEvidenceSbomScanValues() []ImpactEvidenceSbomScan {
	return []ImpactEvidenceSbomScan{
		ImpactEvidenceSbomScanShippedInRange,
		ImpactEvidenceSbomScanShippedOutOfRange,
		ImpactEvidenceSbomScanModuleNotInSbom,
	}
}
