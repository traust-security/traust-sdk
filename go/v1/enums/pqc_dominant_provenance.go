// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
