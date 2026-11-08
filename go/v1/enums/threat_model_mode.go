// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
