// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

package types

import (
	"github.com/traust-security/traust-sdk/go/v1/enums"
)

type AdapterResult struct {
	Findings   []AdapterFinding `json:"findings"`
	FocusAreas []string         `json:"focus_areas,omitempty"`
	Metadata   Metadata         `json:"metadata"`
	ScannedAt  string           `json:"scanned_at"`
	Summary    *Summary         `json:"summary,omitempty"`
	Target     string           `json:"target"`
}

type AdapterFinding struct {
	AsvsReferences      []string                   `json:"asvs_references,omitempty"`
	AttackPattern       *string                    `json:"attack_pattern,omitempty"`
	BlockedExternal     *bool                      `json:"blocked_external,omitempty"`
	Capec               []string                   `json:"capec,omitempty"`
	Category            *string                    `json:"category,omitempty"`
	Cvss                *ReportFindingCvss         `json:"cvss,omitempty"`
	Cwes                []string                   `json:"cwes"`
	Dependency          *ReportFindingDependency   `json:"dependency,omitempty"`
	Description         string                     `json:"description"`
	Disposition         *ReportDisposition         `json:"disposition,omitempty"`
	EffectiveSeverity   *enums.Severity            `json:"effective_severity,omitempty"`
	Evidence            []EvidenceBlock            `json:"evidence,omitempty"`
	Fingerprint         *string                    `json:"fingerprint,omitempty"`
	FingerprintAlgo     *string                    `json:"fingerprint_algo,omitempty"`
	Id                  string                     `json:"id"`
	IsolationBoundary   *string                    `json:"isolation_boundary,omitempty"`
	IsolationDimensions []enums.IsolationDimension `json:"isolation_dimensions,omitempty"`
	Locations           []Location                 `json:"locations"`
	Origin              *enums.FindingOrigin       `json:"origin,omitempty"`
	OwaspLlmReferences  []string                   `json:"owasp_llm_references,omitempty"`
	Passes              []int                      `json:"passes,omitempty"`
	PeachReferences     []string                   `json:"peach_references,omitempty"`
	PqcClassification   *enums.PqcClassification   `json:"pqc_classification,omitempty"`
	Remediation         string                     `json:"remediation"`
	RemediationEffort   *enums.RemediationEffort   `json:"remediation_effort,omitempty"`
	Severity            enums.Severity             `json:"severity"`
	SourceFindings      []string                   `json:"source_findings,omitempty"`
	Title               string                     `json:"title"`
	ValidationStatus    *enums.Validity            `json:"validation_status,omitempty"`
}

type Metadata struct {
	ScannerVersion *string `json:"scanner_version,omitempty"`
	Tool           string  `json:"tool"`
}

type Summary struct {
	BySeverity map[string]int `json:"by_severity,omitempty"`
	Total      int            `json:"total"`
}
