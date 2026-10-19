// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
