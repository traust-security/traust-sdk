// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
