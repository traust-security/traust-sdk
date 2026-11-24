// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
