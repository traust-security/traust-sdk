// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
