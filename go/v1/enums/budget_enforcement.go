// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// BudgetEnforcement — Values of `enforcement` on a lane's budget in traust.schema.json.
type BudgetEnforcement string

const (
	BudgetEnforcementNone     BudgetEnforcement = "none"
	BudgetEnforcementObserve  BudgetEnforcement = "observe"
	BudgetEnforcementAdvisory BudgetEnforcement = "advisory"
	BudgetEnforcementEnforced BudgetEnforcement = "enforced"
)

// BudgetEnforcementValues returns all valid BudgetEnforcement values.
func BudgetEnforcementValues() []BudgetEnforcement {
	return []BudgetEnforcement{
		BudgetEnforcementNone,
		BudgetEnforcementObserve,
		BudgetEnforcementAdvisory,
		BudgetEnforcementEnforced,
	}
}
