// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
