// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package types

type SessionResult struct {
	Error   interface{} `json:"error"`
	Session string      `json:"session"`
	Usage   []Usage     `json:"usage"`
}

type Usage struct {
	CacheCreation *int        `json:"cache_creation,omitempty"`
	CacheRead     *int        `json:"cache_read,omitempty"`
	Model         string      `json:"model"`
	OccurredAt    string      `json:"occurred_at"`
	Provider      string      `json:"provider"`
	Session       interface{} `json:"session,omitempty"`
	Skill         interface{} `json:"skill,omitempty"`
	TokensIn      int         `json:"tokens_in"`
	TokensOut     int         `json:"tokens_out"`
}
