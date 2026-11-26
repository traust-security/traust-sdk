// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// PqcScoringCeiling — Values of `scoring_ceiling` in pqc-decision-tree.schema.json.
type PqcScoringCeiling string

const (
	PqcScoringCeilingPartial PqcScoringCeiling = "partial"
	PqcScoringCeilingNo      PqcScoringCeiling = "no"
)

// PqcScoringCeilingValues returns all valid PqcScoringCeiling values.
func PqcScoringCeilingValues() []PqcScoringCeiling {
	return []PqcScoringCeiling{
		PqcScoringCeilingPartial,
		PqcScoringCeilingNo,
	}
}
