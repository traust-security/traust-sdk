// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
