// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ImpactEvidenceGovulncheck — Govulncheck reachability classification.
type ImpactEvidenceGovulncheck string

const (
	ImpactEvidenceGovulncheckSymbolReachable            ImpactEvidenceGovulncheck = "symbol_reachable"
	ImpactEvidenceGovulncheckPackageImportedNotObserved ImpactEvidenceGovulncheck = "package_imported_not_observed"
	ImpactEvidenceGovulncheckModuleRequiredNotObserved  ImpactEvidenceGovulncheck = "module_required_not_observed"
)

// ImpactEvidenceGovulncheckValues returns all valid ImpactEvidenceGovulncheck values.
func ImpactEvidenceGovulncheckValues() []ImpactEvidenceGovulncheck {
	return []ImpactEvidenceGovulncheck{
		ImpactEvidenceGovulncheckSymbolReachable,
		ImpactEvidenceGovulncheckPackageImportedNotObserved,
		ImpactEvidenceGovulncheckModuleRequiredNotObserved,
	}
}
