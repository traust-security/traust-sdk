// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
