// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
