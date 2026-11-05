// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
