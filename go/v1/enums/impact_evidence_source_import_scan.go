// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
