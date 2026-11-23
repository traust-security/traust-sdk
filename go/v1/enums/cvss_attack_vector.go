// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
