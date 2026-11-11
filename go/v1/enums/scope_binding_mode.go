// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
