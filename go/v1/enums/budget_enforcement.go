// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
