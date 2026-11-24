// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
