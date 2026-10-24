// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
