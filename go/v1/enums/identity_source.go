// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
