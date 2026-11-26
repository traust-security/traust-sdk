// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
