// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// FleetFixMatcherKind — Values of `kind` in fleet-fix.schema.json.
type FleetFixMatcherKind string

const (
	FleetFixMatcherKindPinnedRefLine FleetFixMatcherKind = "pinned_ref_line"
	FleetFixMatcherKindAstGrep       FleetFixMatcherKind = "ast_grep"
)

// FleetFixMatcherKindValues returns all valid FleetFixMatcherKind values.
func FleetFixMatcherKindValues() []FleetFixMatcherKind {
	return []FleetFixMatcherKind{
		FleetFixMatcherKindPinnedRefLine,
		FleetFixMatcherKindAstGrep,
	}
}
