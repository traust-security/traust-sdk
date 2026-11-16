// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// PqcFactsAssessmentBasis — Values of `assessment_basis` in pqc-facts.schema.json.
type PqcFactsAssessmentBasis string

const (
	PqcFactsAssessmentBasisSource   PqcFactsAssessmentBasis = "source"
	PqcFactsAssessmentBasisSbomOnly PqcFactsAssessmentBasis = "sbom-only"
)

// PqcFactsAssessmentBasisValues returns all valid PqcFactsAssessmentBasis values.
func PqcFactsAssessmentBasisValues() []PqcFactsAssessmentBasis {
	return []PqcFactsAssessmentBasis{
		PqcFactsAssessmentBasisSource,
		PqcFactsAssessmentBasisSbomOnly,
	}
}
