// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
