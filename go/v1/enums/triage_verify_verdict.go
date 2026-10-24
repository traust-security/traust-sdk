// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
