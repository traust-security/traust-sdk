// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
