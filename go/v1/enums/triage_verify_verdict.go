// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
