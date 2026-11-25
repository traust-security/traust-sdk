// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
