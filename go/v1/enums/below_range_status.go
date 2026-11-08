// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
