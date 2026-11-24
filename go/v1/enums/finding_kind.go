// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
