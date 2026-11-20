// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package types

type ProductDefinitionsMap struct {
	Mappings ProductDefinitionsMapMappings `json:"mappings"`
}

type ProductDefinitionsMapMappings struct {
	Packages map[string]string `json:"packages,omitempty"`
	Repos    map[string]string `json:"repos,omitempty"`
}
