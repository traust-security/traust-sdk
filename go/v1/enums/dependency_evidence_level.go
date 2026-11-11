// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// DependencyEvidenceLevel — Strongest evidence tier that produced this classification: symbol (govulncheck reachability) > binary (ELF) > manifest (lockfile/source grep).
type DependencyEvidenceLevel string

const (
	DependencyEvidenceLevelSymbol      DependencyEvidenceLevel = "symbol"
	DependencyEvidenceLevelSymbolUsage DependencyEvidenceLevel = "symbol-usage"
	DependencyEvidenceLevelBinary      DependencyEvidenceLevel = "binary"
	DependencyEvidenceLevelManifest    DependencyEvidenceLevel = "manifest"
	DependencyEvidenceLevelNone        DependencyEvidenceLevel = "none"
)

// DependencyEvidenceLevelValues returns all valid DependencyEvidenceLevel values.
func DependencyEvidenceLevelValues() []DependencyEvidenceLevel {
	return []DependencyEvidenceLevel{
		DependencyEvidenceLevelSymbol,
		DependencyEvidenceLevelSymbolUsage,
		DependencyEvidenceLevelBinary,
		DependencyEvidenceLevelManifest,
		DependencyEvidenceLevelNone,
	}
}
