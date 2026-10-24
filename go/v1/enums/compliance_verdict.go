// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package enums

// ComplianceVerdict — Compliance control assessment verdict.
type ComplianceVerdict string

const (
	ComplianceVerdictSatisfied     ComplianceVerdict = "satisfied"
	ComplianceVerdictNotSatisfied  ComplianceVerdict = "not_satisfied"
	ComplianceVerdictNotApplicable ComplianceVerdict = "not_applicable"
	ComplianceVerdictNotAssessed   ComplianceVerdict = "not_assessed"
)

// ComplianceVerdictValues returns all valid ComplianceVerdict values.
func ComplianceVerdictValues() []ComplianceVerdict {
	return []ComplianceVerdict{
		ComplianceVerdictSatisfied,
		ComplianceVerdictNotSatisfied,
		ComplianceVerdictNotApplicable,
		ComplianceVerdictNotAssessed,
	}
}
