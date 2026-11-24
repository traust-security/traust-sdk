// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// CvssPrivilegesRequired — Privileges the probe identity ACTUALLY held at success time.
type CvssPrivilegesRequired string

const (
	CvssPrivilegesRequiredNone CvssPrivilegesRequired = "none"
	CvssPrivilegesRequiredLow  CvssPrivilegesRequired = "low"
	CvssPrivilegesRequiredHigh CvssPrivilegesRequired = "high"
)

// CvssPrivilegesRequiredValues returns all valid CvssPrivilegesRequired values.
func CvssPrivilegesRequiredValues() []CvssPrivilegesRequired {
	return []CvssPrivilegesRequired{
		CvssPrivilegesRequiredNone,
		CvssPrivilegesRequiredLow,
		CvssPrivilegesRequiredHigh,
	}
}
