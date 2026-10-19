// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
