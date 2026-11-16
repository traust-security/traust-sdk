// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
