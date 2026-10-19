// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// PatchEvidenceKind — Kind of base-versus-patch evidence carried for a fix. Defined once in remediation.schema.json and $ref'd by verification.schema.json, so a `proves` claim means the same thing on both sides of the fix. What each kind may legitimately conclude is bounded by the per-path evidence ceilings in the harness's docs/disposition-ledger.md §8a; a kind that claims more than its ceiling is a defect in the kind.
type PatchEvidenceKind string

const (
	PatchEvidenceKindRegression          PatchEvidenceKind = "regression"
	PatchEvidenceKindMutation            PatchEvidenceKind = "mutation"
	PatchEvidenceKindProperty            PatchEvidenceKind = "property"
	PatchEvidenceKindScannerDifferential PatchEvidenceKind = "scanner_differential"
	PatchEvidenceKindExploit             PatchEvidenceKind = "exploit"
)

// PatchEvidenceKindValues returns all valid PatchEvidenceKind values.
func PatchEvidenceKindValues() []PatchEvidenceKind {
	return []PatchEvidenceKind{
		PatchEvidenceKindRegression,
		PatchEvidenceKindMutation,
		PatchEvidenceKindProperty,
		PatchEvidenceKindScannerDifferential,
		PatchEvidenceKindExploit,
	}
}
