// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
