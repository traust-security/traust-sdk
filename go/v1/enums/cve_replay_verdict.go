// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
