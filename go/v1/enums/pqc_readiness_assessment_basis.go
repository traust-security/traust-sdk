// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
