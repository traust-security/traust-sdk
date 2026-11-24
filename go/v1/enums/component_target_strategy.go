// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ComponentTargetStrategy — Deterministic resolution strategy that verified a component PURL against a pinned repository tree.
type ComponentTargetStrategy string

const (
	ComponentTargetStrategySharedRelease      ComponentTargetStrategy = "shared_release"
	ComponentTargetStrategyPackageRelease     ComponentTargetStrategy = "package_release"
	ComponentTargetStrategyWorkspacePackage   ComponentTargetStrategy = "workspace_package"
	ComponentTargetStrategyExternalProvenance ComponentTargetStrategy = "external_provenance"
)

// ComponentTargetStrategyValues returns all valid ComponentTargetStrategy values.
func ComponentTargetStrategyValues() []ComponentTargetStrategy {
	return []ComponentTargetStrategy{
		ComponentTargetStrategySharedRelease,
		ComponentTargetStrategyPackageRelease,
		ComponentTargetStrategyWorkspacePackage,
		ComponentTargetStrategyExternalProvenance,
	}
}
