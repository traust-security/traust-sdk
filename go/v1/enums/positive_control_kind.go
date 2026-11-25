// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
