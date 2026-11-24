// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
