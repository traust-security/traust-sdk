// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// DispositionAssurance — Historical disposition assurance labels, retained for compatibility. Evidence strength is now evidence_grade E0-E3; human review is actor provenance, not an automatic evidence upgrade.
type DispositionAssurance string

const (
	DispositionAssuranceExecutionProven DispositionAssurance = "execution_proven"
	DispositionAssuranceHumanReviewed   DispositionAssurance = "human_reviewed"
	DispositionAssuranceMachineVerified DispositionAssurance = "machine_verified"
	DispositionAssuranceClaimed         DispositionAssurance = "claimed"
)

// DispositionAssuranceValues returns all valid DispositionAssurance values.
func DispositionAssuranceValues() []DispositionAssurance {
	return []DispositionAssurance{
		DispositionAssuranceExecutionProven,
		DispositionAssuranceHumanReviewed,
		DispositionAssuranceMachineVerified,
		DispositionAssuranceClaimed,
	}
}
