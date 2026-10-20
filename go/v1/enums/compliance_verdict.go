// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
