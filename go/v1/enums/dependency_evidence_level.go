// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
