// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package types

type OrgParameters struct {
	DeclaredBy string                                  `json:"declared_by"`
	DeclaredOn *string                                 `json:"declared_on,omitempty"`
	Note       *string                                 `json:"note,omitempty"`
	Parameters map[string]OrgParametersParametersEntry `json:"parameters"`
	Version    int                                     `json:"version"`
}

type OrgParametersParametersEntry struct {
	Controls []string    `json:"controls"`
	Note     *string     `json:"note,omitempty"`
	Value    interface{} `json:"value"`
}
