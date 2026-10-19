// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// EvidenceGrade — P4 evidence grade carried from the validation report; E2/E3 demote the event to machine-static class regardless of source type.
type EvidenceGrade string

const (
	EvidenceGradeE0 EvidenceGrade = "E0"
	EvidenceGradeE1 EvidenceGrade = "E1"
	EvidenceGradeE2 EvidenceGrade = "E2"
	EvidenceGradeE3 EvidenceGrade = "E3"
)

// EvidenceGradeValues returns all valid EvidenceGrade values.
func EvidenceGradeValues() []EvidenceGrade {
	return []EvidenceGrade{
		EvidenceGradeE0,
		EvidenceGradeE1,
		EvidenceGradeE2,
		EvidenceGradeE3,
	}
}
