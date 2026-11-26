// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
