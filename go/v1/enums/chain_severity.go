// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// ChainSeverity — Values of `chain_severity` in validation.schema.json.
type ChainSeverity string

const (
	ChainSeverityCritical ChainSeverity = "critical"
	ChainSeverityHigh     ChainSeverity = "high"
	ChainSeverityMedium   ChainSeverity = "medium"
	ChainSeverityLow      ChainSeverity = "low"
)

// ChainSeverityValues returns all valid ChainSeverity values.
func ChainSeverityValues() []ChainSeverity {
	return []ChainSeverity{
		ChainSeverityCritical,
		ChainSeverityHigh,
		ChainSeverityMedium,
		ChainSeverityLow,
	}
}
