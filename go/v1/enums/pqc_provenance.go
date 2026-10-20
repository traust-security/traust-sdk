// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
