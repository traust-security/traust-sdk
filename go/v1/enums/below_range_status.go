// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// BelowRangeStatus — Whether a commit's below-range packets cover every package: degraded when some packages went unanswered by the advisory source or some advisories were unreadable.
type BelowRangeStatus string

const (
	BelowRangeStatusComplete BelowRangeStatus = "complete"
	BelowRangeStatusDegraded BelowRangeStatus = "degraded"
)

// BelowRangeStatusValues returns all valid BelowRangeStatus values.
func BelowRangeStatusValues() []BelowRangeStatus {
	return []BelowRangeStatus{
		BelowRangeStatusComplete,
		BelowRangeStatusDegraded,
	}
}
