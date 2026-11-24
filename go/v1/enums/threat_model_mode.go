// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
