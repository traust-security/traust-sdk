// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
