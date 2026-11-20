// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// PqcFipsInteraction — Values of `verdict` in pqc-decision-tree.schema.json.
type PqcFipsInteraction string

const (
	PqcFipsInteractionBlockedByProviderVersion PqcFipsInteraction = "blocked-by-provider-version"
	PqcFipsInteractionNoPenalty                PqcFipsInteraction = "no-penalty"
	PqcFipsInteractionFipsValidationGap        PqcFipsInteraction = "fips-validation-gap"
	PqcFipsInteractionPqcBlockedByFipsMode     PqcFipsInteraction = "pqc-blocked-by-fips-mode"
)

// PqcFipsInteractionValues returns all valid PqcFipsInteraction values.
func PqcFipsInteractionValues() []PqcFipsInteraction {
	return []PqcFipsInteraction{
		PqcFipsInteractionBlockedByProviderVersion,
		PqcFipsInteractionNoPenalty,
		PqcFipsInteractionFipsValidationGap,
		PqcFipsInteractionPqcBlockedByFipsMode,
	}
}
