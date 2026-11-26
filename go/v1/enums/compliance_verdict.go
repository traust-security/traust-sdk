// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
