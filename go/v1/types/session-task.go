// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package types

type SessionTask struct {
	Args          []string          `json:"args"`
	Input         []string          `json:"input"`
	Mode          string            `json:"mode"`
	Models        map[string]string `json:"models"`
	Output        SessionTaskOutput `json:"output"`
	RunId         interface{}       `json:"run_id"`
	SchemaVersion interface{}       `json:"schema_version"`
	Skill         string            `json:"skill"`
}

type SessionTaskOutput struct {
	Path   interface{} `json:"path"`
	Schema string      `json:"schema"`
}
