// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
