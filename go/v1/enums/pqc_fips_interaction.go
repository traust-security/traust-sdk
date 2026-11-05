// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
