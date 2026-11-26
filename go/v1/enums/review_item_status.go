// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
