// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
