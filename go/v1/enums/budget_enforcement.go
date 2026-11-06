// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
