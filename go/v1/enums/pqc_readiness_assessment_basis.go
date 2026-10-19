// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
