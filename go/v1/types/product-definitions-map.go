// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

package types

type ProductDefinitionsMap struct {
	Mappings ProductDefinitionsMapMappings `json:"mappings"`
}

type ProductDefinitionsMapMappings struct {
	Packages map[string]string `json:"packages,omitempty"`
	Repos    map[string]string `json:"repos,omitempty"`
}
