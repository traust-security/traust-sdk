// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
