// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
