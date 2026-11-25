// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
