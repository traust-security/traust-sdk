// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
