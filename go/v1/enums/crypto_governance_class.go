// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// CryptoGovernanceClass — Values of `class` in pqc-readiness.schema.json.
type CryptoGovernanceClass string

const (
	CryptoGovernanceClassSelf            CryptoGovernanceClass = "self"
	CryptoGovernanceClassLanguageRuntime CryptoGovernanceClass = "language-runtime"
	CryptoGovernanceClassPlatform        CryptoGovernanceClass = "platform"
	CryptoGovernanceClassInfrastructure  CryptoGovernanceClass = "infrastructure"
)

// CryptoGovernanceClassValues returns all valid CryptoGovernanceClass values.
func CryptoGovernanceClassValues() []CryptoGovernanceClass {
	return []CryptoGovernanceClass{
		CryptoGovernanceClassSelf,
		CryptoGovernanceClassLanguageRuntime,
		CryptoGovernanceClassPlatform,
		CryptoGovernanceClassInfrastructure,
	}
}
