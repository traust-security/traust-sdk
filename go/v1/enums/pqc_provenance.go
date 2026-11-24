// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// PqcProvenance — Values of `provenance` in pqc-decision-tree.schema.json.
type PqcProvenance string

const (
	PqcProvenanceInheritedPlatform       PqcProvenance = "inherited-platform"
	PqcProvenanceInheritedConstrained    PqcProvenance = "inherited-constrained"
	PqcProvenanceDelegatedDependency     PqcProvenance = "delegated-dependency"
	PqcProvenanceNativeFirstParty        PqcProvenance = "native-first-party"
	PqcProvenanceVendored                PqcProvenance = "vendored"
	PqcProvenanceExternalized            PqcProvenance = "externalized"
	PqcProvenanceNotAssessableFromSource PqcProvenance = "not-assessable-from-source"
)

// PqcProvenanceValues returns all valid PqcProvenance values.
func PqcProvenanceValues() []PqcProvenance {
	return []PqcProvenance{
		PqcProvenanceInheritedPlatform,
		PqcProvenanceInheritedConstrained,
		PqcProvenanceDelegatedDependency,
		PqcProvenanceNativeFirstParty,
		PqcProvenanceVendored,
		PqcProvenanceExternalized,
		PqcProvenanceNotAssessableFromSource,
	}
}
