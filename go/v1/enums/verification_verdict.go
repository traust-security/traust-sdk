// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// VerificationVerdict — Values of `verdict` in verification.schema.json.
type VerificationVerdict string

const (
	VerificationVerdictResolved          VerificationVerdict = "resolved"
	VerificationVerdictPartiallyResolved VerificationVerdict = "partially_resolved"
	VerificationVerdictUnresolved        VerificationVerdict = "unresolved"
	VerificationVerdictNewApproach       VerificationVerdict = "new_approach"
	VerificationVerdictRegression        VerificationVerdict = "regression"
	VerificationVerdictFalsePositive     VerificationVerdict = "false_positive"
	VerificationVerdictRiskAccepted      VerificationVerdict = "risk_accepted"
)

// VerificationVerdictValues returns all valid VerificationVerdict values.
func VerificationVerdictValues() []VerificationVerdict {
	return []VerificationVerdict{
		VerificationVerdictResolved,
		VerificationVerdictPartiallyResolved,
		VerificationVerdictUnresolved,
		VerificationVerdictNewApproach,
		VerificationVerdictRegression,
		VerificationVerdictFalsePositive,
		VerificationVerdictRiskAccepted,
	}
}
