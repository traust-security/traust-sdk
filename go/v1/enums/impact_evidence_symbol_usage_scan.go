// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
