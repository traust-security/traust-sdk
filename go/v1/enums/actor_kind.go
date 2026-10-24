// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
