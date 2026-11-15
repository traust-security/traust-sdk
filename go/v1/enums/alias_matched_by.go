// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
