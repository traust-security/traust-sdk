// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// ComponentTargetStrategy — Deterministic resolution strategy that verified a component PURL against a pinned repository tree.
type ComponentTargetStrategy string

const (
	ComponentTargetStrategySharedRelease    ComponentTargetStrategy = "shared_release"
	ComponentTargetStrategyPackageRelease   ComponentTargetStrategy = "package_release"
	ComponentTargetStrategyWorkspacePackage ComponentTargetStrategy = "workspace_package"
)

// ComponentTargetStrategyValues returns all valid ComponentTargetStrategy values.
func ComponentTargetStrategyValues() []ComponentTargetStrategy {
	return []ComponentTargetStrategy{
		ComponentTargetStrategySharedRelease,
		ComponentTargetStrategyPackageRelease,
		ComponentTargetStrategyWorkspacePackage,
	}
}
