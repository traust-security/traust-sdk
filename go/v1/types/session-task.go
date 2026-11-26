// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

package types

type SessionTask struct {
	Args            []string          `json:"args"`
	Input           []string          `json:"input"`
	LeaseGeneration interface{}       `json:"lease_generation,omitempty"`
	Mode            string            `json:"mode"`
	Models          map[string]string `json:"models"`
	Output          SessionTaskOutput `json:"output"`
	RunId           interface{}       `json:"run_id"`
	SchemaVersion   interface{}       `json:"schema_version"`
	Skill           string            `json:"skill"`
}

type SessionTaskOutput struct {
	Path   interface{} `json:"path"`
	Schema string      `json:"schema"`
}
