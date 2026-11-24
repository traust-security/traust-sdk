// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
