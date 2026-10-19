// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// InterfaceComplexity — Values of `complexity` in isolation-review.schema.json.
type InterfaceComplexity string

const (
	InterfaceComplexityLow    InterfaceComplexity = "low"
	InterfaceComplexityMedium InterfaceComplexity = "medium"
	InterfaceComplexityHigh   InterfaceComplexity = "high"
)

// InterfaceComplexityValues returns all valid InterfaceComplexity values.
func InterfaceComplexityValues() []InterfaceComplexity {
	return []InterfaceComplexity{
		InterfaceComplexityLow,
		InterfaceComplexityMedium,
		InterfaceComplexityHigh,
	}
}
