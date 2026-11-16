// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
