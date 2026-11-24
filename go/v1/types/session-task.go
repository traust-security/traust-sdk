// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
