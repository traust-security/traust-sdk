// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
