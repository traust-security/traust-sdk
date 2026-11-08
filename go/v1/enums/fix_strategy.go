// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// FixStrategy — Values of `fix_strategy` in remediation.schema.json.
type FixStrategy string

const (
	FixStrategyInputValidation     FixStrategy = "input-validation"
	FixStrategyBoundResource       FixStrategy = "bound-resource"
	FixStrategyRbacScopeDown       FixStrategy = "rbac-scope-down"
	FixStrategyAuthAdd             FixStrategy = "auth-add"
	FixStrategyTlsEnforce          FixStrategy = "tls-enforce"
	FixStrategyUrlAllowlist        FixStrategy = "url-allowlist"
	FixStrategyDigestPin           FixStrategy = "digest-pin"
	FixStrategyConfigDefaultHarden FixStrategy = "config-default-harden"
	FixStrategyApiRestrict         FixStrategy = "api-restrict"
	FixStrategyRemoveFeature       FixStrategy = "remove-feature"
	FixStrategyDependencyBump      FixStrategy = "dependency-bump"
	FixStrategyOther               FixStrategy = "other"
)

// FixStrategyValues returns all valid FixStrategy values.
func FixStrategyValues() []FixStrategy {
	return []FixStrategy{
		FixStrategyInputValidation,
		FixStrategyBoundResource,
		FixStrategyRbacScopeDown,
		FixStrategyAuthAdd,
		FixStrategyTlsEnforce,
		FixStrategyUrlAllowlist,
		FixStrategyDigestPin,
		FixStrategyConfigDefaultHarden,
		FixStrategyApiRestrict,
		FixStrategyRemoveFeature,
		FixStrategyDependencyBump,
		FixStrategyOther,
	}
}
