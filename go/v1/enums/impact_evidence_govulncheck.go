// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
