// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
