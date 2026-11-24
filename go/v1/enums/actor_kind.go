// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
