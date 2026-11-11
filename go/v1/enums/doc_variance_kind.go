// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
