// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// Effort — Canonical effort-size labels for explicitly selected future-write contracts. Null means unknown size; blocked_external is an independent boolean and never implies a size.
type Effort string

const (
	EffortXs Effort = "xs"
	EffortS  Effort = "s"
	EffortM  Effort = "m"
	EffortL  Effort = "l"
	EffortXl Effort = "xl"
)

// EffortValues returns all valid Effort values.
func EffortValues() []Effort {
	return []Effort{
		EffortXs,
		EffortS,
		EffortM,
		EffortL,
		EffortXl,
	}
}
