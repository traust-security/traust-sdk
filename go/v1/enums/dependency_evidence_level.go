// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
