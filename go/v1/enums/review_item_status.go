// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// ReviewItemStatus — Review queue workflow status, independent of finding validity.
type ReviewItemStatus string

const (
	ReviewItemStatusPending          ReviewItemStatus = "pending"
	ReviewItemStatusConfirmed        ReviewItemStatus = "confirmed"
	ReviewItemStatusRejected         ReviewItemStatus = "rejected"
	ReviewItemStatusAccepted         ReviewItemStatus = "accepted"
	ReviewItemStatusClosedUnreviewed ReviewItemStatus = "closed_unreviewed"
)

// ReviewItemStatusValues returns all valid ReviewItemStatus values.
func ReviewItemStatusValues() []ReviewItemStatus {
	return []ReviewItemStatus{
		ReviewItemStatusPending,
		ReviewItemStatusConfirmed,
		ReviewItemStatusRejected,
		ReviewItemStatusAccepted,
		ReviewItemStatusClosedUnreviewed,
	}
}
