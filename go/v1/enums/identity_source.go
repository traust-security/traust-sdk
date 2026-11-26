// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// IdentitySource — The trusted mechanism that verified a principal, never an actor string supplied by an agent.
type IdentitySource string

const (
	IdentitySourceLocalUser IdentitySource = "local-user"
	IdentitySourceDatabase  IdentitySource = "database"
	IdentitySourceOidc      IdentitySource = "oidc"
)

// IdentitySourceValues returns all valid IdentitySource values.
func IdentitySourceValues() []IdentitySource {
	return []IdentitySource{
		IdentitySourceLocalUser,
		IdentitySourceDatabase,
		IdentitySourceOidc,
	}
}
