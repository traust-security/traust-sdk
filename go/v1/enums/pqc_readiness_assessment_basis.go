// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// PqcReadinessAssessmentBasis — Values of `assessment_basis` in pqc-readiness.schema.json.
type PqcReadinessAssessmentBasis string

const (
	PqcReadinessAssessmentBasisSource        PqcReadinessAssessmentBasis = "source"
	PqcReadinessAssessmentBasisSbomOnly      PqcReadinessAssessmentBasis = "sbom-only"
	PqcReadinessAssessmentBasisSourceRuntime PqcReadinessAssessmentBasis = "source+runtime"
)

// PqcReadinessAssessmentBasisValues returns all valid PqcReadinessAssessmentBasis values.
func PqcReadinessAssessmentBasisValues() []PqcReadinessAssessmentBasis {
	return []PqcReadinessAssessmentBasis{
		PqcReadinessAssessmentBasisSource,
		PqcReadinessAssessmentBasisSbomOnly,
		PqcReadinessAssessmentBasisSourceRuntime,
	}
}
