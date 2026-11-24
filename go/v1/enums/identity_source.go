// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
