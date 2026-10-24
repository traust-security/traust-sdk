// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
