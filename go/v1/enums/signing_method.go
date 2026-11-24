// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package enums

// SigningMethod — How the Ledger signs layer Merkle roots: keypair (a cosign private key) or identity (keyless Sigstore/OIDC).
type SigningMethod string

const (
	SigningMethodKeypair  SigningMethod = "keypair"
	SigningMethodIdentity SigningMethod = "identity"
)

// SigningMethodValues returns all valid SigningMethod values.
func SigningMethodValues() []SigningMethod {
	return []SigningMethod{
		SigningMethodKeypair,
		SigningMethodIdentity,
	}
}
