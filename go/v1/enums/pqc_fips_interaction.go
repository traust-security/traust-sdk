// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
