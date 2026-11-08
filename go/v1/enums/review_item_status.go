// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
