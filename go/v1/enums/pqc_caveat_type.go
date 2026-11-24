// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
