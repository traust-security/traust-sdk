// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
