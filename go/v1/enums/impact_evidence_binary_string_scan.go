// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ImpactEvidenceBinaryStringScan — ELF binary string scan result.
type ImpactEvidenceBinaryStringScan string

const (
	ImpactEvidenceBinaryStringScanPackagePathPresent ImpactEvidenceBinaryStringScan = "package_path_present"
	ImpactEvidenceBinaryStringScanPackagePathAbsent  ImpactEvidenceBinaryStringScan = "package_path_absent"
	ImpactEvidenceBinaryStringScanNotScanned         ImpactEvidenceBinaryStringScan = "not_scanned"
)

// ImpactEvidenceBinaryStringScanValues returns all valid ImpactEvidenceBinaryStringScan values.
func ImpactEvidenceBinaryStringScanValues() []ImpactEvidenceBinaryStringScan {
	return []ImpactEvidenceBinaryStringScan{
		ImpactEvidenceBinaryStringScanPackagePathPresent,
		ImpactEvidenceBinaryStringScanPackagePathAbsent,
		ImpactEvidenceBinaryStringScanNotScanned,
	}
}
