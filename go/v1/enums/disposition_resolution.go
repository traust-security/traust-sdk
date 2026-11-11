// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// DispositionResolution — Canonical decision resolution is open, resolved or risk_accepted. Historical progress, partial-resolution and regression labels remain readable; the owner applies line-specific or operational semantics.
type DispositionResolution string

const (
	DispositionResolutionOpen                 DispositionResolution = "open"
	DispositionResolutionFixInProgress        DispositionResolution = "fix_in_progress"
	DispositionResolutionResolved             DispositionResolution = "resolved"
	DispositionResolutionPartiallyResolved    DispositionResolution = "partially_resolved"
	DispositionResolutionRiskAccepted         DispositionResolution = "risk_accepted"
	DispositionResolutionRegressionIntroduced DispositionResolution = "regression_introduced"
)

// DispositionResolutionValues returns all valid DispositionResolution values.
func DispositionResolutionValues() []DispositionResolution {
	return []DispositionResolution{
		DispositionResolutionOpen,
		DispositionResolutionFixInProgress,
		DispositionResolutionResolved,
		DispositionResolutionPartiallyResolved,
		DispositionResolutionRiskAccepted,
		DispositionResolutionRegressionIntroduced,
	}
}
