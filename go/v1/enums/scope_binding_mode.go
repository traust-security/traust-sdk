// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
