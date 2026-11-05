// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
