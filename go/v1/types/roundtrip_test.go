package types_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/enums"
	"github.com/traust-security/traust-sdk/go/v1/types"
	"github.com/traust-security/traust-sdk/go/v1/validate"
)

func TestReportRoundTrip(t *testing.T) {
	original := minimalReport()

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	if err := validate.Validate("report", data); err != nil {
		t.Fatalf("validate report schema: %v", err)
	}

	var roundTripped types.Report
	if err := json.Unmarshal(data, &roundTripped); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if roundTripped.Title != original.Title {
		t.Errorf("title: got %q, want %q", roundTripped.Title, original.Title)
	}
	if roundTripped.Metadata.Date != original.Metadata.Date {
		t.Errorf("metadata.date: got %q, want %q", roundTripped.Metadata.Date, original.Metadata.Date)
	}
	if roundTripped.Metadata.Scope != original.Metadata.Scope {
		t.Errorf("metadata.scope: got %q, want %q", roundTripped.Metadata.Scope, original.Metadata.Scope)
	}
	if roundTripped.ExecutiveSummary.Prose != original.ExecutiveSummary.Prose {
		t.Errorf("executive_summary.prose: got %q, want %q", roundTripped.ExecutiveSummary.Prose, original.ExecutiveSummary.Prose)
	}
	if len(roundTripped.SeverityCriteria) != len(original.SeverityCriteria) {
		t.Errorf("severity_criteria length: got %d, want %d", len(roundTripped.SeverityCriteria), len(original.SeverityCriteria))
	}
	if len(roundTripped.FindingsSummary) != len(original.FindingsSummary) {
		t.Errorf("findings_summary length: got %d, want %d", len(roundTripped.FindingsSummary), len(original.FindingsSummary))
	}
	if len(roundTripped.RemediationRoadmap) != len(original.RemediationRoadmap) {
		t.Errorf("remediation_roadmap length: got %d, want %d", len(roundTripped.RemediationRoadmap), len(original.RemediationRoadmap))
	}
}

func TestLayerRoundTrip(t *testing.T) {
	original := minimalLayer()

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	if err := validate.Validate("layer", data); err != nil {
		t.Fatalf("validate layer schema: %v", err)
	}

	var roundTripped types.Layer
	if err := json.Unmarshal(data, &roundTripped); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if roundTripped.Metadata.AuditReport != original.Metadata.AuditReport {
		t.Errorf("metadata.audit_report: got %q, want %q", roundTripped.Metadata.AuditReport, original.Metadata.AuditReport)
	}
	if roundTripped.Metadata.Repository != original.Metadata.Repository {
		t.Errorf("metadata.repository: got %q, want %q", roundTripped.Metadata.Repository, original.Metadata.Repository)
	}
}

func TestTriageRoundTrip(t *testing.T) {
	original := minimalTriage()

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	if err := validate.Validate("triage", data); err != nil {
		t.Fatalf("validate triage schema: %v", err)
	}

	var roundTripped types.Triage
	if err := json.Unmarshal(data, &roundTripped); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if roundTripped.TriageCompleted != original.TriageCompleted {
		t.Errorf("triage_completed: got %q, want %q", roundTripped.TriageCompleted, original.TriageCompleted)
	}
	if roundTripped.TriageContext.Repo != original.TriageContext.Repo {
		t.Errorf("triage_context.repo: got %q, want %q", roundTripped.TriageContext.Repo, original.TriageContext.Repo)
	}
}

func minimalReport() types.Report {
	longDefinition := "Exploitable vulnerability with severe impact on confidentiality, integrity, or availability of the system."
	return types.Report{
		Title: "Security Audit Report",
		Metadata: types.ReportMetadata{
			Date:  "2026-08-12",
			Scope: "Full source code review of the application repository and its dependencies.",
		},
		ExecutiveSummary: types.ExecutiveSummary{
			Prose: "This security audit assessed the target application for common vulnerability classes. No critical issues were identified during this review cycle.",
			SeverityCounts: types.ExecutiveSummarySeverityCounts{
				Critical:      0,
				High:          0,
				Medium:        0,
				Low:           0,
				Informational: 0,
			},
		},
		SeverityCriteria: []types.SeverityCriterion{
			{Level: enums.SeverityCritical, Definition: longDefinition},
			{Level: enums.SeverityHigh, Definition: longDefinition},
			{Level: enums.SeverityMedium, Definition: longDefinition},
			{Level: enums.SeverityLow, Definition: longDefinition},
		},
		Findings: []types.ReportFinding{},
		FindingsSummary: []types.SeverityCountEntry{
			{Severity: enums.SeverityCritical, Count: 0, FindingIds: []string{}},
			{Severity: enums.SeverityHigh, Count: 0, FindingIds: []string{}},
			{Severity: enums.SeverityMedium, Count: 0, FindingIds: []string{}},
			{Severity: enums.SeverityLow, Count: 0, FindingIds: []string{}},
		},
		RemediationRoadmap: []types.RoadmapItem{
			{
				Priority:  "low",
				Action:    "Continue periodic security reviews and dependency updates.",
				Addresses: []string{"general-hygiene"},
			},
		},
	}
}

func minimalLayer() types.Layer {
	return types.Layer{
		Metadata: types.LayerMetadata{
			AuditReport:    "example-security-audit.json",
			Repository:     "https://github.com/example/repo",
			Created:        "2026-08-12T10:00:00Z",
			HarnessVersion: "0.2.0",
		},
		Events:      []types.Event{},
		NeedsReview: []types.ReviewItem{},
	}
}

func minimalTriage() types.Triage {
	return types.Triage{
		TriageCompleted: "2026-08-12",
		TriageContext: types.TriageContext{
			Environment:     "local development environment with full source access",
			VotesPerFinding: 1,
			Repo:            "https://github.com/example/repo",
			HarnessVersion:  "0.2.0",
		},
		Summary: types.TriageSummary{
			InputCount:     0,
			TruePositives:  0,
			Hardening:      0,
			FalsePositives: 0,
			Undetermined:   0,
			Duplicates:     0,
			BySeverity: types.TriageSummaryBySeverity{
				Critical: 0,
				High:     0,
				Medium:   0,
				Low:      0,
			},
		},
		Findings: []types.TriageFinding{},
	}
}

// A report's not_affected section survives a Go decode and re-encode, and the
// re-encoded bytes still pass the contracts schema: a Go client never drops the
// eliminated matches or their rationale (remediation plan step 12).
func TestReportNotAffectedRoundTrip(t *testing.T) {
	ptr := func(s string) *string { return &s }
	classification := func(c enums.ImpactClassification) *enums.ImpactClassification { return &c }
	ecosystem := enums.EcosystemGo
	level := func(l enums.DependencyEvidenceLevel) *enums.DependencyEvidenceLevel { return &l }
	cases := []struct {
		name  string
		entry types.NotAffectedEntry
	}{
		{
			name: "impact row: version not in range",
			entry: types.NotAffectedEntry{
				Dependency: types.ReportFindingDependency{
					Advisory:         "CVE-2024-0001",
					Module:           "golang.org/x/net",
					Ecosystem:        &ecosystem,
					InstalledVersion: ptr("v0.33.0"),
					Classification:   classification(enums.ImpactClassificationVersionNotInRange),
					EvidenceLevel:    level(enums.DependencyEvidenceLevelManifest),
					ImpactArtifact:   ptr("impact row: widget@abcdef0, go, golang.org/x/net, v0.33.0"),
				},
				Rationale: "At or past the fix in every source's range.",
			},
		},
		{
			name: "vendor reference, optional dependency fields unset",
			entry: types.NotAffectedEntry{
				Dependency: types.ReportFindingDependency{Advisory: "CVE-2024-0002", Module: "openssl"},
				Rationale:  "Red Hat states this product is not affected.",
				Reference:  ptr("RHSA-2024:0001"),
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			report := minimalReport()
			report.NotAffected = []types.NotAffectedEntry{tc.entry}
			data, err := json.Marshal(report)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			if err := validate.Validate("report", data); err != nil {
				t.Fatalf("validate report schema: %v", err)
			}
			var back types.Report
			if err := json.Unmarshal(data, &back); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			again, err := json.Marshal(back)
			if err != nil {
				t.Fatalf("re-marshal: %v", err)
			}
			if err := validate.Validate("report", again); err != nil {
				t.Fatalf("validate re-encoded report: %v", err)
			}
			// every field survives, the reference and impact fields included
			if len(back.NotAffected) != 1 || !reflect.DeepEqual(back.NotAffected[0], tc.entry) {
				t.Fatalf("not_affected changed on round-trip:\n got %+v\nwant %+v", back.NotAffected, tc.entry)
			}
		})
	}
}

// An advisory-backed vuln-scan finding's cvss and dependency blocks
// (traust-contracts#190) survive a Go decode and re-encode, field for field,
// and the re-encoded document still passes the contract.
func TestVulnFindingsAdvisoryFieldsRoundTrip(t *testing.T) {
	doc := []byte(`{
  "target": "example/repo",
  "scanned_at": "2026-01-01T00:00:00Z",
  "focus_areas": ["dependencies"],
  "metadata": {"repo": "repo", "repo_slug": "REPO", "scanned_ref": "abcdef0",
    "harness_version": "0.1.0", "baseline": null, "baseline_findings": 0},
  "findings": [{
    "id": "REPO-abcdef0-001", "file": "go.mod", "line": 3, "category": "supply-chain",
    "severity": "critical", "confidence": 0.9, "title": "Vulnerable golang.org/x/net",
    "description": "golang.org/x/net v0.20.0 is in the advisory's affected range.",
    "recommendation": "Upgrade golang.org/x/net past the fixed version.",
    "cvss": {"vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H", "score": 9.8},
    "dependency": {"advisory": "CVE-2024-0001", "module": "golang.org/x/net", "ecosystem": "go",
      "installed_version": "v0.20.0", "classification": "likely_affected",
      "evidence_level": "manifest", "impact_artifact": "impact row: repo@abcdef0"}
  }],
  "summary": {"total": 1, "critical": 1, "high": 0, "medium": 0, "low": 0,
    "informational": 0, "known": 0, "low_confidence": 0}
}`)
	if err := validate.Validate("vuln-findings", doc); err != nil {
		t.Fatalf("validate vuln-findings: %v", err)
	}
	var first types.VulnFindings
	if err := json.Unmarshal(doc, &first); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	again, err := json.Marshal(first)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if err := validate.Validate("vuln-findings", again); err != nil {
		t.Fatalf("validate re-encoded vuln-findings: %v", err)
	}
	var second types.VulnFindings
	if err := json.Unmarshal(again, &second); err != nil {
		t.Fatalf("unmarshal re-encoded: %v", err)
	}
	f := second.Findings[0]
	if f.Cvss == nil || f.Dependency == nil || !reflect.DeepEqual(first.Findings[0], f) {
		t.Fatalf("advisory fields changed on round-trip:\n got %+v\nwant %+v", f, first.Findings[0])
	}
}
