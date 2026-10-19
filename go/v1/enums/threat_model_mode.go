// Code generated from traust-contracts 3ed68158121322ce06525de170aa383e057b6a7b. DO NOT EDIT.

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
