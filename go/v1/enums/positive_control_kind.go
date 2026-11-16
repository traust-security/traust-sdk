// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
