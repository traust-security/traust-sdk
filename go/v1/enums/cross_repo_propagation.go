// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
