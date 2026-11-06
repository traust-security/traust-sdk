// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
