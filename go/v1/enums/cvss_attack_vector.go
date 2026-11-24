// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
