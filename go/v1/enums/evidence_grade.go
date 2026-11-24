// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

package enums

// EvidenceGrade — The single evidence-strength scale, E0 strongest through E3 weakest. An absent or null grade is ungraded; a bare claim or human approval does not supply a grade.
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
