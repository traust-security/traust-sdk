// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
