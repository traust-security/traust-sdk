// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// ImpactEvidenceSymbolUsageScan — Source grep for the advisory's vulnerable symbol names — textual API usage, stronger than a pin, weaker than reachability.
type ImpactEvidenceSymbolUsageScan string

const (
	ImpactEvidenceSymbolUsageScanSymbolsUsed     ImpactEvidenceSymbolUsageScan = "symbols_used"
	ImpactEvidenceSymbolUsageScanSymbolsNotFound ImpactEvidenceSymbolUsageScan = "symbols_not_found"
)

// ImpactEvidenceSymbolUsageScanValues returns all valid ImpactEvidenceSymbolUsageScan values.
func ImpactEvidenceSymbolUsageScanValues() []ImpactEvidenceSymbolUsageScan {
	return []ImpactEvidenceSymbolUsageScan{
		ImpactEvidenceSymbolUsageScanSymbolsUsed,
		ImpactEvidenceSymbolUsageScanSymbolsNotFound,
	}
}
