// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
