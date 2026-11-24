// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
