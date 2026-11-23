// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// TriageVerifyVerdict — exploitable|mitigated|needs_manual_test are the skill's canonical values; confirmed|refuted|hardening are accepted presentation aliases.
type TriageVerifyVerdict string

const (
	TriageVerifyVerdictExploitable     TriageVerifyVerdict = "exploitable"
	TriageVerifyVerdictMitigated       TriageVerifyVerdict = "mitigated"
	TriageVerifyVerdictNeedsManualTest TriageVerifyVerdict = "needs_manual_test"
	TriageVerifyVerdictConfirmed       TriageVerifyVerdict = "confirmed"
	TriageVerifyVerdictRefuted         TriageVerifyVerdict = "refuted"
	TriageVerifyVerdictHardening       TriageVerifyVerdict = "hardening"
)

// TriageVerifyVerdictValues returns all valid TriageVerifyVerdict values.
func TriageVerifyVerdictValues() []TriageVerifyVerdict {
	return []TriageVerifyVerdict{
		TriageVerifyVerdictExploitable,
		TriageVerifyVerdictMitigated,
		TriageVerifyVerdictNeedsManualTest,
		TriageVerifyVerdictConfirmed,
		TriageVerifyVerdictRefuted,
		TriageVerifyVerdictHardening,
	}
}
