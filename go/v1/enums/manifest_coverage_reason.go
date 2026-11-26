// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// ManifestCoverageReason — Why the SBOM tool could not resolve a dependency manifest. An unresolved manifest makes a package missing from the SBOM unknown, never absent.
type ManifestCoverageReason string

const (
	ManifestCoverageReasonNoLockfile          ManifestCoverageReason = "no-lockfile"
	ManifestCoverageReasonUnpinned            ManifestCoverageReason = "unpinned"
	ManifestCoverageReasonUnsupportedManifest ManifestCoverageReason = "unsupported-manifest"
)

// ManifestCoverageReasonValues returns all valid ManifestCoverageReason values.
func ManifestCoverageReasonValues() []ManifestCoverageReason {
	return []ManifestCoverageReason{
		ManifestCoverageReasonNoLockfile,
		ManifestCoverageReasonUnpinned,
		ManifestCoverageReasonUnsupportedManifest,
	}
}
