// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
