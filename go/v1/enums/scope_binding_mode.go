// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

package enums

// ScopeBindingMode — Values of `scope_binding_mode` in validation.schema.json.
type ScopeBindingMode string

const (
	ScopeBindingModeExplicit ScopeBindingMode = "explicit"
	ScopeBindingModeInline   ScopeBindingMode = "inline"
	ScopeBindingModeInferred ScopeBindingMode = "inferred"
	ScopeBindingModeMixed    ScopeBindingMode = "mixed"
	ScopeBindingModeNone     ScopeBindingMode = "none"
)

// ScopeBindingModeValues returns all valid ScopeBindingMode values.
func ScopeBindingModeValues() []ScopeBindingMode {
	return []ScopeBindingMode{
		ScopeBindingModeExplicit,
		ScopeBindingModeInline,
		ScopeBindingModeInferred,
		ScopeBindingModeMixed,
		ScopeBindingModeNone,
	}
}
