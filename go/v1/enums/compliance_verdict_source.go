// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// ComplianceVerdictSource — deterministic controls: 'check' (code-computed) or 'human_override' — never 'agent' (amendment 1)
type ComplianceVerdictSource string

const (
	ComplianceVerdictSourceCheck         ComplianceVerdictSource = "check"
	ComplianceVerdictSourceAgent         ComplianceVerdictSource = "agent"
	ComplianceVerdictSourceHumanOverride ComplianceVerdictSource = "human_override"
)

// ComplianceVerdictSourceValues returns all valid ComplianceVerdictSource values.
func ComplianceVerdictSourceValues() []ComplianceVerdictSource {
	return []ComplianceVerdictSource{
		ComplianceVerdictSourceCheck,
		ComplianceVerdictSourceAgent,
		ComplianceVerdictSourceHumanOverride,
	}
}
