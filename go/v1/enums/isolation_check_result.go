// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// IsolationCheckResult — Outcome of one isolation-review dimension or PQC readiness check. `na` means not applicable.
type IsolationCheckResult string

const (
	IsolationCheckResultYes     IsolationCheckResult = "yes"
	IsolationCheckResultPartial IsolationCheckResult = "partial"
	IsolationCheckResultNo      IsolationCheckResult = "no"
	IsolationCheckResultNa      IsolationCheckResult = "na"
)

// IsolationCheckResultValues returns all valid IsolationCheckResult values.
func IsolationCheckResultValues() []IsolationCheckResult {
	return []IsolationCheckResult{
		IsolationCheckResultYes,
		IsolationCheckResultPartial,
		IsolationCheckResultNo,
		IsolationCheckResultNa,
	}
}
