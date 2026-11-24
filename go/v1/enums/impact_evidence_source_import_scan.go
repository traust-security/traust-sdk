// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
