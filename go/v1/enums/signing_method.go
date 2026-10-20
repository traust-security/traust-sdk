// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
