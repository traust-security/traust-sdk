// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
