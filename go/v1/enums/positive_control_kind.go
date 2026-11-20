// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// PositiveControlKind — Expected behavior of the control action, independent of the finding under test.
type PositiveControlKind string

const (
	PositiveControlKindMustSucceed PositiveControlKind = "must_succeed"
	PositiveControlKindMustDeny    PositiveControlKind = "must_deny"
)

// PositiveControlKindValues returns all valid PositiveControlKind values.
func PositiveControlKindValues() []PositiveControlKind {
	return []PositiveControlKind{
		PositiveControlKindMustSucceed,
		PositiveControlKindMustDeny,
	}
}
