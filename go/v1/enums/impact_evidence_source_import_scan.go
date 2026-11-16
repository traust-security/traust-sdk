// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
