// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
