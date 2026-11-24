// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
