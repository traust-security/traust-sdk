// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

package enums

// ActorKind — Actor type for disposition events. Determines identity-provider verification requirements and auto-accept eligibility.
type ActorKind string

const (
	ActorKindHuman   ActorKind = "human"
	ActorKindMachine ActorKind = "machine"
)

// ActorKindValues returns all valid ActorKind values.
func ActorKindValues() []ActorKind {
	return []ActorKind{
		ActorKindHuman,
		ActorKindMachine,
	}
}
