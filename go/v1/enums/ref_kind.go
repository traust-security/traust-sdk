// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

package enums

// RefKind — Kind of git reference being audited (branch, tag, default branch, or dist-git stream).
type RefKind string

const (
	RefKindBranch  RefKind = "branch"
	RefKindTag     RefKind = "tag"
	RefKindDefault RefKind = "default"
	RefKindStream  RefKind = "stream"
)

// RefKindValues returns all valid RefKind values.
func RefKindValues() []RefKind {
	return []RefKind{
		RefKindBranch,
		RefKindTag,
		RefKindDefault,
		RefKindStream,
	}
}
