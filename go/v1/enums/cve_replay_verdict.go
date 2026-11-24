// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
