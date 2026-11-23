// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
