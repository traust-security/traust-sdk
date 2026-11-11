// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// IsolationPosture — Values of `overall` in isolation-review.schema.json.
type IsolationPosture string

const (
	IsolationPostureStrong      IsolationPosture = "strong"
	IsolationPostureAdequate    IsolationPosture = "adequate"
	IsolationPostureWeak        IsolationPosture = "weak"
	IsolationPostureCriticalGap IsolationPosture = "critical-gap"
)

// IsolationPostureValues returns all valid IsolationPosture values.
func IsolationPostureValues() []IsolationPosture {
	return []IsolationPosture{
		IsolationPostureStrong,
		IsolationPostureAdequate,
		IsolationPostureWeak,
		IsolationPostureCriticalGap,
	}
}
