// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
