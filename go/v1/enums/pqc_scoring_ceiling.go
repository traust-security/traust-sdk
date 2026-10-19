// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
