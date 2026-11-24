// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
