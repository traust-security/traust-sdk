// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
