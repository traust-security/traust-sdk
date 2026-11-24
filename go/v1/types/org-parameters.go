// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
