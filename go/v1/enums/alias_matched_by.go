// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
