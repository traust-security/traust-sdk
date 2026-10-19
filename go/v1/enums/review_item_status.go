// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
