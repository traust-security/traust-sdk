// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
