// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
