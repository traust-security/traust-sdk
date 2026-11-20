// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// CveReplayVerdict — Whether the original recorded audit detected a replayed advisory; a missed verdict requires human-confirmed historical presence.
type CveReplayVerdict string

const (
	CveReplayVerdictDetected      CveReplayVerdict = "detected"
	CveReplayVerdictUnclear       CveReplayVerdict = "unclear"
	CveReplayVerdictMissed        CveReplayVerdict = "missed"
	CveReplayVerdictNotApplicable CveReplayVerdict = "not_applicable"
)

// CveReplayVerdictValues returns all valid CveReplayVerdict values.
func CveReplayVerdictValues() []CveReplayVerdict {
	return []CveReplayVerdict{
		CveReplayVerdictDetected,
		CveReplayVerdictUnclear,
		CveReplayVerdictMissed,
		CveReplayVerdictNotApplicable,
	}
}
