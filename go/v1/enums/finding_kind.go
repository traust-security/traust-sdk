// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// FindingKind — Finding.kind in the authoritative findings model; independent of whether the finding is real.
type FindingKind string

const (
	FindingKindVulnerability FindingKind = "vulnerability"
	FindingKindWeakness      FindingKind = "weakness"
)

// FindingKindValues returns all valid FindingKind values.
func FindingKindValues() []FindingKind {
	return []FindingKind{
		FindingKindVulnerability,
		FindingKindWeakness,
	}
}
