// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
