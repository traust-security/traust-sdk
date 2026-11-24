// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

package enums

// PqcDominantProvenance — Values of `dominant` in pqc-readiness.schema.json.
type PqcDominantProvenance string

const (
	PqcDominantProvenanceInheritedPlatform       PqcDominantProvenance = "inherited-platform"
	PqcDominantProvenanceInheritedConstrained    PqcDominantProvenance = "inherited-constrained"
	PqcDominantProvenanceDelegatedDependency     PqcDominantProvenance = "delegated-dependency"
	PqcDominantProvenanceNativeFirstParty        PqcDominantProvenance = "native-first-party"
	PqcDominantProvenanceVendored                PqcDominantProvenance = "vendored"
	PqcDominantProvenanceExternalized            PqcDominantProvenance = "externalized"
	PqcDominantProvenanceNotAssessableFromSource PqcDominantProvenance = "not-assessable-from-source"
	PqcDominantProvenanceNone                    PqcDominantProvenance = "none"
)

// PqcDominantProvenanceValues returns all valid PqcDominantProvenance values.
func PqcDominantProvenanceValues() []PqcDominantProvenance {
	return []PqcDominantProvenance{
		PqcDominantProvenanceInheritedPlatform,
		PqcDominantProvenanceInheritedConstrained,
		PqcDominantProvenanceDelegatedDependency,
		PqcDominantProvenanceNativeFirstParty,
		PqcDominantProvenanceVendored,
		PqcDominantProvenanceExternalized,
		PqcDominantProvenanceNotAssessableFromSource,
		PqcDominantProvenanceNone,
	}
}
