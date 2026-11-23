// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
