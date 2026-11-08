// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type ImpactAnalysis struct {
	Metadata ImpactAnalysisMetadata `json:"metadata"`
	Repos    []RepoEntry            `json:"repos"`
	Summary  ImpactAnalysisSummary  `json:"summary"`
}

type Evidence struct {
	BinaryLinkedLibrary   *enums.ImpactEvidenceBinaryLinkedLibrary `json:"binary_linked_library,omitempty"`
	BinaryStringScan      *enums.ImpactEvidenceBinaryStringScan    `json:"binary_string_scan,omitempty"`
	BinarySymbolScan      *enums.ImpactEvidenceBinarySymbolScan    `json:"binary_symbol_scan,omitempty"`
	EvidenceLevel         *enums.DependencyEvidenceLevel           `json:"evidence_level,omitempty"`
	FeaturePatternMatches interface{}                              `json:"feature_pattern_matches,omitempty"`
	Govulncheck           *enums.ImpactEvidenceGovulncheck         `json:"govulncheck,omitempty"`
	GovulncheckTrace      interface{}                              `json:"govulncheck_trace,omitempty"`
	L1DependsOn           interface{}                              `json:"l1_depends_on,omitempty"`
	L1VersionInRange      interface{}                              `json:"l1_version_in_range,omitempty"`
	L4PackageImported     interface{}                              `json:"l4_package_imported,omitempty"`
	L4PackagesFound       interface{}                              `json:"l4_packages_found,omitempty"`
	ManifestScan          *enums.ImpactEvidenceManifestScan        `json:"manifest_scan,omitempty"`
	ManifestVersion       interface{}                              `json:"manifest_version,omitempty"`
	NeedsManualTrace      interface{}                              `json:"needs_manual_trace,omitempty"`
	Notes                 interface{}                              `json:"notes,omitempty"`
	SbomScan              *enums.ImpactEvidenceSbomScan            `json:"sbom_scan,omitempty"`
	SbomShippedVersion    interface{}                              `json:"sbom_shipped_version,omitempty"`
	SourceImportScan      *enums.ImpactEvidenceSourceImportScan    `json:"source_import_scan,omitempty"`
	SymbolUsageScan       *enums.ImpactEvidenceSymbolUsageScan     `json:"symbol_usage_scan,omitempty"`
}

type ImpactAnalysisMetadata struct {
	AdvisorySources       []string                                `json:"advisory_sources,omitempty"`
	Cve                   string                                  `json:"cve"`
	Ecosystem             *enums.Ecosystem                        `json:"ecosystem,omitempty"`
	FeatureDescription    interface{}                             `json:"feature_description,omitempty"`
	FixedVersion          interface{}                             `json:"fixed_version,omitempty"`
	GeneratedAt           string                                  `json:"generated_at"`
	HarnessVersion        string                                  `json:"harness_version"`
	Module                string                                  `json:"module"`
	Options               ImpactAnalysisMetadataOptions           `json:"options"`
	PortfolioGraphDb      interface{}                             `json:"portfolio_graph_db,omitempty"`
	PortfolioGraphVersion interface{}                             `json:"portfolio_graph_version,omitempty"`
	PreparationScope      *ImpactAnalysisMetadataPreparationScope `json:"preparation_scope,omitempty"`
	TiersExecuted         []enums.ImpactTier                      `json:"tiers_executed"`
	VulnerablePackages    []string                                `json:"vulnerable_packages,omitempty"`
	VulnerableRange       string                                  `json:"vulnerable_range"`
	VulnerableSymbols     []string                                `json:"vulnerable_symbols,omitempty"`
}

type ImpactAnalysisMetadataOptions struct {
	BinaryScan  *bool `json:"binary_scan,omitempty"`
	Govulncheck *bool `json:"govulncheck,omitempty"`
	Sweep       *bool `json:"sweep,omitempty"`
}

type ImpactAnalysisMetadataPreparationScope struct {
	Digest  string                                              `json:"digest"`
	Members []ImpactAnalysisMetadataPreparationScopeMembersItem `json:"members"`
}

type ImpactAnalysisMetadataPreparationScopeMembersItem struct {
	Ecosystem enums.Ecosystem `json:"ecosystem"`
	Module    string          `json:"module"`
	Version   interface{}     `json:"version"`
}

type ImpactAnalysisSummary struct {
	Affected           int      `json:"affected"`
	Inconclusive       int      `json:"inconclusive"`
	LikelyAffected     int      `json:"likely_affected"`
	NotImported        int      `json:"not_imported"`
	NotObserved        int      `json:"not_observed"`
	ProductSurfaces    []string `json:"product_surfaces,omitempty"`
	ReposInBlastRadius int      `json:"repos_in_blast_radius"`
	VersionInRange     int      `json:"version_in_range"`
	VersionNotInRange  int      `json:"version_not_in_range"`
}

type RepoEntry struct {
	Classification enums.ImpactClassification `json:"classification"`
	Direct         interface{}                `json:"direct,omitempty"`
	Evidence       Evidence                   `json:"evidence"`
	Products       []string                   `json:"products,omitempty"`
	Repo           string                     `json:"repo"`
	Version        interface{}                `json:"version,omitempty"`
}
