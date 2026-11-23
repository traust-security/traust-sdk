// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

package types

type ProductDefinitionsMap struct {
	Mappings ProductDefinitionsMapMappings `json:"mappings"`
}

type ProductDefinitionsMapMappings struct {
	Packages map[string]string `json:"packages,omitempty"`
	Repos    map[string]string `json:"repos,omitempty"`
}
