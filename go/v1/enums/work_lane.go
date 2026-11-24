// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// WorkLane — Operational vocabulary for traust_ops.run_order.lane; Core enforces new writes, while stored historical text remains readable.
type WorkLane string

const (
	WorkLaneFullAuditValidate    WorkLane = "full-audit+validate"
	WorkLaneFullAudit            WorkLane = "full-audit"
	WorkLaneImpactLane           WorkLane = "impact-lane"
	WorkLaneIacLane              WorkLane = "iac-lane"
	WorkLaneIacBaseline          WorkLane = "iac-baseline"
	WorkLaneReleasePassthrough   WorkLane = "release-passthrough"
	WorkLaneThreatModelReview    WorkLane = "threat-model-review"
	WorkLaneDepsLane             WorkLane = "deps-lane"
	WorkLaneDiffScan             WorkLane = "diff-scan"
	WorkLaneDiffScanQuarterly    WorkLane = "diff-scan-quarterly"
	WorkLaneThreatModelQuarterly WorkLane = "threat-model-quarterly"
	WorkLaneNone                 WorkLane = "none"
)

// WorkLaneValues returns all valid WorkLane values.
func WorkLaneValues() []WorkLane {
	return []WorkLane{
		WorkLaneFullAuditValidate,
		WorkLaneFullAudit,
		WorkLaneImpactLane,
		WorkLaneIacLane,
		WorkLaneIacBaseline,
		WorkLaneReleasePassthrough,
		WorkLaneThreatModelReview,
		WorkLaneDepsLane,
		WorkLaneDiffScan,
		WorkLaneDiffScanQuarterly,
		WorkLaneThreatModelQuarterly,
		WorkLaneNone,
	}
}
