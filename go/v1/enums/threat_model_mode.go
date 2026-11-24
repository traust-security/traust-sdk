// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ThreatModelMode — Values of `mode` in threat-model.schema.json.
type ThreatModelMode string

const (
	ThreatModelModeInterview              ThreatModelMode = "interview"
	ThreatModelModeBootstrap              ThreatModelMode = "bootstrap"
	ThreatModelModeBootstrapThenInterview ThreatModelMode = "bootstrap-then-interview"
)

// ThreatModelModeValues returns all valid ThreatModelMode values.
func ThreatModelModeValues() []ThreatModelMode {
	return []ThreatModelMode{
		ThreatModelModeInterview,
		ThreatModelModeBootstrap,
		ThreatModelModeBootstrapThenInterview,
	}
}
