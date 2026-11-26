// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
