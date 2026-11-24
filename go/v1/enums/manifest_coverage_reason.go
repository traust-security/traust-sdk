// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
