// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
