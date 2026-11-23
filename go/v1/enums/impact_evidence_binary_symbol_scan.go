// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package enums

// ImpactEvidenceBinarySymbolScan — Vulnerable symbol names present in the shipped binary's strings (.dynstr) — the binary references the vulnerable function, not just the library.
type ImpactEvidenceBinarySymbolScan string

const (
	ImpactEvidenceBinarySymbolScanSymbolsPresent ImpactEvidenceBinarySymbolScan = "symbols_present"
	ImpactEvidenceBinarySymbolScanSymbolsAbsent  ImpactEvidenceBinarySymbolScan = "symbols_absent"
)

// ImpactEvidenceBinarySymbolScanValues returns all valid ImpactEvidenceBinarySymbolScan values.
func ImpactEvidenceBinarySymbolScanValues() []ImpactEvidenceBinarySymbolScan {
	return []ImpactEvidenceBinarySymbolScan{
		ImpactEvidenceBinarySymbolScanSymbolsPresent,
		ImpactEvidenceBinarySymbolScanSymbolsAbsent,
	}
}
