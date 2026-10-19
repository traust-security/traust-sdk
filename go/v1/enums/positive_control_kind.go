// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
