// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// ReviewItemStatus — Queue state.
type ReviewItemStatus string

const (
	ReviewItemStatusPending   ReviewItemStatus = "pending"
	ReviewItemStatusConfirmed ReviewItemStatus = "confirmed"
	ReviewItemStatusRejected  ReviewItemStatus = "rejected"
)

// ReviewItemStatusValues returns all valid ReviewItemStatus values.
func ReviewItemStatusValues() []ReviewItemStatus {
	return []ReviewItemStatus{
		ReviewItemStatusPending,
		ReviewItemStatusConfirmed,
		ReviewItemStatusRejected,
	}
}
