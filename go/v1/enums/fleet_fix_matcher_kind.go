// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
