// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
