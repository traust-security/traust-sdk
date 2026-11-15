// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// PqcCaveatType — Values of `caveat_type` in pqc-readiness.schema.json.
type PqcCaveatType string

const (
	PqcCaveatTypeInfraMlkemUnverified PqcCaveatType = "infra_mlkem_unverified"
	PqcCaveatTypeInfraMlkemBlocked    PqcCaveatType = "infra_mlkem_blocked"
)

// PqcCaveatTypeValues returns all valid PqcCaveatType values.
func PqcCaveatTypeValues() []PqcCaveatType {
	return []PqcCaveatType{
		PqcCaveatTypeInfraMlkemUnverified,
		PqcCaveatTypeInfraMlkemBlocked,
	}
}
