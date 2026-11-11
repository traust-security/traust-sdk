// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
