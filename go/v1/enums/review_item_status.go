// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
