// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
