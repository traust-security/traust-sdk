// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
