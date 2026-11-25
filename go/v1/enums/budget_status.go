// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
