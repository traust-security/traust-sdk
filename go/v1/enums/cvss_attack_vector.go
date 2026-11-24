// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package enums

// CvssAttackVector — Vector the exploit ACTUALLY used.
type CvssAttackVector string

const (
	CvssAttackVectorNetwork  CvssAttackVector = "network"
	CvssAttackVectorAdjacent CvssAttackVector = "adjacent"
	CvssAttackVectorLocal    CvssAttackVector = "local"
	CvssAttackVectorPhysical CvssAttackVector = "physical"
)

// CvssAttackVectorValues returns all valid CvssAttackVector values.
func CvssAttackVectorValues() []CvssAttackVector {
	return []CvssAttackVector{
		CvssAttackVectorNetwork,
		CvssAttackVectorAdjacent,
		CvssAttackVectorLocal,
		CvssAttackVectorPhysical,
	}
}
