// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
