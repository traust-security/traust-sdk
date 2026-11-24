// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
