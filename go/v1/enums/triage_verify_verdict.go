// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
