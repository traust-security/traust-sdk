// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
