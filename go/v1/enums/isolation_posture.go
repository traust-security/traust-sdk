// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
