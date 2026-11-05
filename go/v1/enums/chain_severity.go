// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
