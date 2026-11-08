// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// ManifestCoverageReason — Why the SBOM tool could not resolve a dependency manifest. An unresolved manifest makes a package missing from the SBOM unknown, never absent.
type ManifestCoverageReason string

const (
	ManifestCoverageReasonNoLockfile ManifestCoverageReason = "no-lockfile"
	ManifestCoverageReasonUnpinned   ManifestCoverageReason = "unpinned"
)

// ManifestCoverageReasonValues returns all valid ManifestCoverageReason values.
func ManifestCoverageReasonValues() []ManifestCoverageReason {
	return []ManifestCoverageReason{
		ManifestCoverageReasonNoLockfile,
		ManifestCoverageReasonUnpinned,
	}
}
