// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
