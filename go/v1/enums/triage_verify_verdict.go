// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
