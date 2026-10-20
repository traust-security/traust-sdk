// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
