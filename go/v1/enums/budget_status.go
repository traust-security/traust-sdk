// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// BudgetStatus — A lane's month-to-date spend plus an estimate against its budget: within, over, or no budget configured for the lane.
type BudgetStatus string

const (
	BudgetStatusWithin   BudgetStatus = "within"
	BudgetStatusOver     BudgetStatus = "over"
	BudgetStatusNoBudget BudgetStatus = "no_budget"
)

// BudgetStatusValues returns all valid BudgetStatus values.
func BudgetStatusValues() []BudgetStatus {
	return []BudgetStatus{
		BudgetStatusWithin,
		BudgetStatusOver,
		BudgetStatusNoBudget,
	}
}
