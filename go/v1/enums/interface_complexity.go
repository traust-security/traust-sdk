// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
