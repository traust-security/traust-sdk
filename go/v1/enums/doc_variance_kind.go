// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
