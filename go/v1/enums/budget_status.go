// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
