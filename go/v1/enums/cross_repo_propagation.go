// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// CrossRepoPropagation — Whether the ORIGINAL repo has consumed the fix: consumed = lockfile/vendor/SBOM at or past the fixed version; pending = still on a vulnerable version; module_absent = dependency no longer present (removal is also a fix — verifier judges); not_applicable = Shape 2/3 (fix legitimately lives only in fi
type CrossRepoPropagation string

const (
	CrossRepoPropagationConsumed      CrossRepoPropagation = "consumed"
	CrossRepoPropagationPending       CrossRepoPropagation = "pending"
	CrossRepoPropagationModuleAbsent  CrossRepoPropagation = "module_absent"
	CrossRepoPropagationNotApplicable CrossRepoPropagation = "not_applicable"
)

// CrossRepoPropagationValues returns all valid CrossRepoPropagation values.
func CrossRepoPropagationValues() []CrossRepoPropagation {
	return []CrossRepoPropagation{
		CrossRepoPropagationConsumed,
		CrossRepoPropagationPending,
		CrossRepoPropagationModuleAbsent,
		CrossRepoPropagationNotApplicable,
	}
}
