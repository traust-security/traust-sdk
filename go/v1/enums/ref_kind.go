// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
