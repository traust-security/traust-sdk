// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
