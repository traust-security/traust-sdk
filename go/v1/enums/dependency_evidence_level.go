// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
