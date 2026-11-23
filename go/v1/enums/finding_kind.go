// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
