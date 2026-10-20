// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// DocVarianceKind — overclaim = docs promise more security than code delivers (the customer-trust class)
type DocVarianceKind string

const (
	DocVarianceKindOverclaim     DocVarianceKind = "overclaim"
	DocVarianceKindUnderclaim    DocVarianceKind = "underclaim"
	DocVarianceKindOmission      DocVarianceKind = "omission"
	DocVarianceKindContradiction DocVarianceKind = "contradiction"
	DocVarianceKindStale         DocVarianceKind = "stale"
)

// DocVarianceKindValues returns all valid DocVarianceKind values.
func DocVarianceKindValues() []DocVarianceKind {
	return []DocVarianceKind{
		DocVarianceKindOverclaim,
		DocVarianceKindUnderclaim,
		DocVarianceKindOmission,
		DocVarianceKindContradiction,
		DocVarianceKindStale,
	}
}
