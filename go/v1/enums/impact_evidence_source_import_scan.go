// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package enums

// ImpactEvidenceSourceImportScan — Source-level import/include grep for the module.
type ImpactEvidenceSourceImportScan string

const (
	ImpactEvidenceSourceImportScanImportsFound    ImpactEvidenceSourceImportScan = "imports_found"
	ImpactEvidenceSourceImportScanImportsNotFound ImpactEvidenceSourceImportScan = "imports_not_found"
)

// ImpactEvidenceSourceImportScanValues returns all valid ImpactEvidenceSourceImportScan values.
func ImpactEvidenceSourceImportScanValues() []ImpactEvidenceSourceImportScan {
	return []ImpactEvidenceSourceImportScan{
		ImpactEvidenceSourceImportScanImportsFound,
		ImpactEvidenceSourceImportScanImportsNotFound,
	}
}
