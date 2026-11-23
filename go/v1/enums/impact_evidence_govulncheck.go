// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
