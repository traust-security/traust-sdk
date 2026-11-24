// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

package types

type ProductDefinitionsMap struct {
	Mappings ProductDefinitionsMapMappings `json:"mappings"`
}

type ProductDefinitionsMapMappings struct {
	Packages map[string]string `json:"packages,omitempty"`
	Repos    map[string]string `json:"repos,omitempty"`
}
