// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
