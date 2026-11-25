// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
