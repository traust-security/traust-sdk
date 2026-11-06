// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
