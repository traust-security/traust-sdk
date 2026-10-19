// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
