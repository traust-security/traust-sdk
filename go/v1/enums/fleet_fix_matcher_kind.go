// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
