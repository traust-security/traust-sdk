// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
