// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
