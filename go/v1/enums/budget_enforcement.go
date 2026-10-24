// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// BudgetEnforcement — Values of `enforcement` in budget-policy.schema.json.
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
