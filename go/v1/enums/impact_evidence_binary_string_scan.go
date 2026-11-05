// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
