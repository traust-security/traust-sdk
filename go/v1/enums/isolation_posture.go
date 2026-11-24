// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
