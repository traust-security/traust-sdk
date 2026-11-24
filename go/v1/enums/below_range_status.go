// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
