// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type VulnFindings struct {
	Findings              []VulnFindingsFinding  `json:"findings"`
	FocusAreas            []string               `json:"focus_areas"`
	KnownFindings         []KnownFinding         `json:"known_findings,omitempty"`
	Metadata              VulnFindingsMetadata   `json:"metadata"`
	ReviewSurfaceCoverage map[string]interface{} `json:"review_surface_coverage,omitempty"`
	ScannedAt             string                 `json:"scanned_at"`
	Summary               VulnFindingsSummary    `json:"summary"`
	Target                string                 `json:"target"`
}

type KnownFinding struct {
	MatchesBaselineId string `json:"matches_baseline_id"`
	Title             string `json:"title"`
}

type VulnFindingsFinding struct {
	Category         string         `json:"category"`
	Confidence       float64        `json:"confidence"`
	ConfidenceReason *string        `json:"confidence_reason,omitempty"`
	Cwe              *string        `json:"cwe,omitempty"`
	Description      string         `json:"description"`
	ExploitScenario  *string        `json:"exploit_scenario,omitempty"`
	File             string         `json:"file"`
	Fingerprint      *string        `json:"fingerprint,omitempty"`
	FingerprintAlgo  *string        `json:"fingerprint_algo,omitempty"`
	Id               string         `json:"id"`
	Line             interface{}    `json:"line"`
	Recommendation   string         `json:"recommendation"`
	ScannerRef       *string        `json:"scanner_ref,omitempty"`
	Severity         enums.Severity `json:"severity"`
	Title            string         `json:"title"`
}

type VulnFindingsMetadata struct {
	Additional       *VulnFindingsMetadataAdditional `json:"additional,omitempty"`
	Baseline         interface{}                     `json:"baseline"`
	BaselineFindings int                             `json:"baseline_findings"`
	HarnessVersion   string                          `json:"harness_version"`
	Repo             string                          `json:"repo"`
	RepoSlug         string                          `json:"repo_slug"`
	ScannedRef       string                          `json:"scanned_ref"`
}

type VulnFindingsMetadataAdditional struct {
	CoverageDiff *VulnFindingsMetadataAdditionalCoverageDiff `json:"coverage_diff,omitempty"`
}

type VulnFindingsMetadataAdditionalCoverageDiff struct {
	Anchor              string   `json:"anchor"`
	BaselineReport      string   `json:"baseline_report"`
	CallersCovered      []string `json:"callers_covered"`
	ChangedFilesCovered []string `json:"changed_files_covered"`
	RefusedOrSkipped    []string `json:"refused_or_skipped"`
	ResolutionSource    string   `json:"resolution_source"`
}

type VulnFindingsSummary struct {
	Critical      int `json:"critical"`
	High          int `json:"high"`
	Informational int `json:"informational"`
	Known         int `json:"known"`
	Low           int `json:"low"`
	LowConfidence int `json:"low_confidence"`
	Medium        int `json:"medium"`
	Total         int `json:"total"`
}
