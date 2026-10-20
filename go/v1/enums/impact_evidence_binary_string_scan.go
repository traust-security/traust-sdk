// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
