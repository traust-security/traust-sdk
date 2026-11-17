// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
