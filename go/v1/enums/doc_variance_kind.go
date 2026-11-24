// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
