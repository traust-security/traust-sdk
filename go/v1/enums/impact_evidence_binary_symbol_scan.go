// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
