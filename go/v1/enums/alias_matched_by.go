// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// AliasMatchedBy — 'fingerprint' matches auto-confirm; looser tiers pend in needs_review.
type AliasMatchedBy string

const (
	AliasMatchedByFingerprint AliasMatchedBy = "fingerprint"
	AliasMatchedByPathCwe     AliasMatchedBy = "path_cwe"
	AliasMatchedByPathSet     AliasMatchedBy = "path_set"
	AliasMatchedByTitle       AliasMatchedBy = "title"
	AliasMatchedByManual      AliasMatchedBy = "manual"
)

// AliasMatchedByValues returns all valid AliasMatchedBy values.
func AliasMatchedByValues() []AliasMatchedBy {
	return []AliasMatchedBy{
		AliasMatchedByFingerprint,
		AliasMatchedByPathCwe,
		AliasMatchedByPathSet,
		AliasMatchedByTitle,
		AliasMatchedByManual,
	}
}
