// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

type ProductDefinitionsMap struct {
	Mappings ProductDefinitionsMapMappings `json:"mappings"`
}

type ProductDefinitionsMapMappings struct {
	Packages map[string]string `json:"packages,omitempty"`
	Repos    map[string]string `json:"repos,omitempty"`
}
