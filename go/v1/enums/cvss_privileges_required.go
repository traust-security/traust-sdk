// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
