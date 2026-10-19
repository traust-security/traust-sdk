// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
