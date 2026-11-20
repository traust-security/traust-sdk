// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
