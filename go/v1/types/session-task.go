// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
