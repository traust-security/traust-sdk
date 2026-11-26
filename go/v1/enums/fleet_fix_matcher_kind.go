// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
