// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
