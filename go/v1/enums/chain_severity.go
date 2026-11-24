// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
