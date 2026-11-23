// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
