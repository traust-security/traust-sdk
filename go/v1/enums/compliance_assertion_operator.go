// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// ComplianceAssertionOperator — Values of `operator` in compliance-mapping.schema.json.
type ComplianceAssertionOperator string

const (
	ComplianceAssertionOperatorEquals      ComplianceAssertionOperator = "equals"
	ComplianceAssertionOperatorNotEquals   ComplianceAssertionOperator = "not_equals"
	ComplianceAssertionOperatorContains    ComplianceAssertionOperator = "contains"
	ComplianceAssertionOperatorNotContains ComplianceAssertionOperator = "not_contains"
	ComplianceAssertionOperatorRegex       ComplianceAssertionOperator = "regex"
	ComplianceAssertionOperatorNotRegex    ComplianceAssertionOperator = "not_regex"
	ComplianceAssertionOperatorGte         ComplianceAssertionOperator = "gte"
	ComplianceAssertionOperatorLte         ComplianceAssertionOperator = "lte"
	ComplianceAssertionOperatorExists      ComplianceAssertionOperator = "exists"
	ComplianceAssertionOperatorAbsent      ComplianceAssertionOperator = "absent"
	ComplianceAssertionOperatorIn          ComplianceAssertionOperator = "in"
	ComplianceAssertionOperatorNotIn       ComplianceAssertionOperator = "not_in"
)

// ComplianceAssertionOperatorValues returns all valid ComplianceAssertionOperator values.
func ComplianceAssertionOperatorValues() []ComplianceAssertionOperator {
	return []ComplianceAssertionOperator{
		ComplianceAssertionOperatorEquals,
		ComplianceAssertionOperatorNotEquals,
		ComplianceAssertionOperatorContains,
		ComplianceAssertionOperatorNotContains,
		ComplianceAssertionOperatorRegex,
		ComplianceAssertionOperatorNotRegex,
		ComplianceAssertionOperatorGte,
		ComplianceAssertionOperatorLte,
		ComplianceAssertionOperatorExists,
		ComplianceAssertionOperatorAbsent,
		ComplianceAssertionOperatorIn,
		ComplianceAssertionOperatorNotIn,
	}
}
