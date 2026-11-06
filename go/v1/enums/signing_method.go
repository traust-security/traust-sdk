// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
