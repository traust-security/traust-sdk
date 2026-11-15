// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// ImpactEvidenceManifestScan — Lockfile/manifest scan result (Python/Rust/JS/Java — manifest-level evidence, never proves reachability).
type ImpactEvidenceManifestScan string

const (
	ImpactEvidenceManifestScanModulePinnedInRange    ImpactEvidenceManifestScan = "module_pinned_in_range"
	ImpactEvidenceManifestScanModulePinnedOutOfRange ImpactEvidenceManifestScan = "module_pinned_out_of_range"
	ImpactEvidenceManifestScanModuleNotInManifests   ImpactEvidenceManifestScan = "module_not_in_manifests"
	ImpactEvidenceManifestScanNoManifestsFound       ImpactEvidenceManifestScan = "no_manifests_found"
)

// ImpactEvidenceManifestScanValues returns all valid ImpactEvidenceManifestScan values.
func ImpactEvidenceManifestScanValues() []ImpactEvidenceManifestScan {
	return []ImpactEvidenceManifestScan{
		ImpactEvidenceManifestScanModulePinnedInRange,
		ImpactEvidenceManifestScanModulePinnedOutOfRange,
		ImpactEvidenceManifestScanModuleNotInManifests,
		ImpactEvidenceManifestScanNoManifestsFound,
	}
}
