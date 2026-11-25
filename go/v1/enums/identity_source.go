// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
