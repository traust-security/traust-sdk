package storage

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/types"
)

// detailReport returns a report whose findings cover every exposure class:
// one open with a cvss score, effective severity and a triage verdict, one
// bare (untriaged, no disposition), and one each of false_positive, closed
// and hardening.
func detailReport(t *testing.T) []byte {
	t.Helper()
	return replaceJSONField(t, sampleArtifacts(t)["report"], func(document map[string]any) {
		document["findings"] = []any{
			map[string]any{
				"id": "FIND-001", "title": "Open disposed finding", "severity": "high",
				"description": strings.Repeat("Synthetic evidence. ", 4), "remediation": strings.Repeat("Apply a correction. ", 2),
				"fingerprint": strings.Repeat("a", 64), "fingerprint_algo": "v1",
				"validation_status": "confirmed", "category": "injection", "cwes": []any{"CWE-79"},
				"locations":          []any{map[string]any{"path": "test.go", "lines": "1-2"}},
				"cvss":               map[string]any{"score": 8.1, "vector": "CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:U/C:H/I:H/A:H"},
				"effective_severity": "critical",
				"disposition": map[string]any{
					"validity": "confirmed", "resolution": "fix_in_progress", "assurance": "execution_proven",
					"last_updated": "2026-01-02T00:00:00Z", "events": []any{},
				},
			},
			map[string]any{
				"id": "FIND-002", "title": "Bare finding", "severity": "low",
				"description": strings.Repeat("Synthetic baseline. ", 4), "remediation": strings.Repeat("Review the baseline. ", 2),
				"fingerprint": strings.Repeat("b", 64), "cwes": []any{"CWE-200"}, "locations": []any{map[string]any{"path": "bare.go"}},
			},
			map[string]any{
				"id": "FIND-003", "title": "Refuted finding", "severity": "medium",
				"description": strings.Repeat("Synthetic refutation. ", 4), "remediation": strings.Repeat("Close it. ", 4),
				"fingerprint": strings.Repeat("c", 64), "cwes": []any{"CWE-89"}, "locations": []any{map[string]any{"path": "refuted.go"}},
				"disposition": map[string]any{
					"validity": "false_positive", "resolution": "open",
					"last_updated": "2026-01-02T00:00:00Z", "events": []any{},
				},
			},
			map[string]any{
				"id": "FIND-004", "title": "Resolved finding", "severity": "high",
				"description": strings.Repeat("Synthetic resolution. ", 4), "remediation": strings.Repeat("Already fixed. ", 4),
				"fingerprint": strings.Repeat("d", 64), "cwes": []any{"CWE-22"}, "locations": []any{map[string]any{"path": "resolved.go"}},
				"disposition": map[string]any{
					"validity": "confirmed", "resolution": "resolved",
					"last_updated": "2026-01-02T00:00:00Z", "events": []any{},
				},
			},
			map[string]any{
				"id": "FIND-005", "title": "Hardening finding", "severity": "low",
				"description": strings.Repeat("Synthetic posture debt. ", 4), "remediation": strings.Repeat("Harden it. ", 4),
				"fingerprint": strings.Repeat("e", 64), "cwes": []any{"CWE-693"}, "locations": []any{map[string]any{"path": "hardening.go"}},
				"disposition": map[string]any{
					"validity": "hardening", "resolution": "open",
					"last_updated": "2026-01-02T00:00:00Z", "events": []any{},
				},
			},
		}
	})
}

func detailTriage(t *testing.T, completed string, findings []any) []byte {
	t.Helper()
	document := map[string]any{
		"triage_completed": completed,
		"triage_context": map[string]any{
			"environment":       "Synthetic test environment",
			"repo":              "repo",
			"harness_version":   "0.1.0",
			"votes_per_finding": 1,
		},
		"summary": map[string]any{
			"input_count": len(findings), "true_positives": 0, "hardening": 0,
			"false_positives": 0, "undetermined": 0, "duplicates": 0,
			"by_severity": map[string]any{"critical": 0, "high": 0, "medium": 0, "low": 0},
		},
		"findings": findings,
	}
	encoded, err := json.Marshal(document)
	if err != nil {
		t.Fatal(err)
	}
	return encoded
}

func checkCurrentFindingDetail(t *testing.T, client *Client) {
	t.Helper()
	ctx := context.Background()
	const scope = "go-detail-test"
	binding := runBinding(scope, nil)
	reportArtifact := mustParseArtifact(t, detailReport(t), types.ParseReportArtifact)
	if _, err := client.SaveReport(ctx, SaveReportInput{Binding: binding, Artifact: reportArtifact}); err != nil {
		t.Fatal(err)
	}
	triageArtifact := mustParseArtifact(t, detailTriage(t, "2026-01-03", []any{
		map[string]any{
			"id": "f101", "title": "Confirmed authorization issue", "orig_id": "FIND-001",
			"verdict": "true_positive", "severity": "critical", "rationale": "Confirmed with a second account.",
		},
		map[string]any{
			"id": "f102", "title": "Refuted", "orig_id": "FIND-003",
			"verdict": "false_positive", "severity": "medium",
		},
	}), types.ParseTriageArtifact)
	if _, err := client.SaveTriage(ctx, SaveTriageInput{Binding: binding, Artifact: triageArtifact}); err != nil {
		t.Fatal(err)
	}

	rows, err := client.QueryCurrentFindingDetail(ctx, []string{scope})
	if err != nil {
		t.Fatal(err)
	}
	byID := make(map[string]CurrentFindingDetailRow, len(rows))
	for _, row := range rows {
		byID[row.FindingID] = row
	}
	if len(byID) != 5 {
		t.Fatalf("current_finding_detail returned %d rows, want 5: %+v", len(rows), rows)
	}

	open := byID["FIND-001"]
	if open.ExposureClass != "open" {
		t.Errorf("FIND-001 exposure_class = %q, want open", open.ExposureClass)
	}
	if open.CVSSScore == nil || *open.CVSSScore != 8.1 {
		t.Errorf("FIND-001 cvss_score = %v, want 8.1", open.CVSSScore)
	}
	if open.EffectiveSeverity == nil || *open.EffectiveSeverity != "critical" {
		t.Errorf("FIND-001 effective_severity = %v, want critical", open.EffectiveSeverity)
	}
	if open.Category == nil || *open.Category != "injection" {
		t.Errorf("FIND-001 category = %v, want injection", open.Category)
	}
	if open.TriageVerdict == nil || *open.TriageVerdict != "true_positive" {
		t.Errorf("FIND-001 triage_verdict = %v, want true_positive", open.TriageVerdict)
	}
	if open.TriageSeverity == nil || *open.TriageSeverity != "critical" {
		t.Errorf("FIND-001 triage_severity = %v, want critical", open.TriageSeverity)
	}

	bare := byID["FIND-002"]
	if bare.ExposureClass != "open" {
		t.Errorf("FIND-002 exposure_class = %q, want open (no disposition)", bare.ExposureClass)
	}
	if bare.TriageVerdict != nil {
		t.Errorf("FIND-002 triage_verdict = %v, want nil (untriaged)", bare.TriageVerdict)
	}
	if bare.CVSSScore != nil {
		t.Errorf("FIND-002 cvss_score = %v, want nil", bare.CVSSScore)
	}

	if got := byID["FIND-003"].ExposureClass; got != "false_positive" {
		t.Errorf("FIND-003 exposure_class = %q, want false_positive", got)
	}
	if verdict := byID["FIND-003"].TriageVerdict; verdict == nil || *verdict != "false_positive" {
		t.Errorf("FIND-003 triage_verdict = %v, want false_positive", verdict)
	}
	if got := byID["FIND-004"].ExposureClass; got != "closed" {
		t.Errorf("FIND-004 exposure_class = %q, want closed", got)
	}
	if got := byID["FIND-005"].ExposureClass; got != "hardening" {
		t.Errorf("FIND-005 exposure_class = %q, want hardening", got)
	}

	// Scope isolation: another scope reads nothing.
	other, err := client.QueryCurrentFindingDetail(ctx, []string{"other-scope"})
	if err != nil {
		t.Fatal(err)
	}
	if len(other) != 0 {
		t.Fatalf("other scope returned %d rows, want 0", len(other))
	}
}

func checkCurrentFindingDetailNewestTriage(t *testing.T, client *Client) {
	t.Helper()
	ctx := context.Background()
	const scope = "go-detail-newest"
	binding := runBinding(scope, nil)
	reportArtifact := mustParseArtifact(t, detailReport(t), types.ParseReportArtifact)
	if _, err := client.SaveReport(ctx, SaveReportInput{Binding: binding, Artifact: reportArtifact}); err != nil {
		t.Fatal(err)
	}
	first := mustParseArtifact(t, detailTriage(t, "2026-01-03", []any{
		map[string]any{
			"id": "f301", "title": "First adjudication", "orig_id": "FIND-001",
			"verdict": "true_positive", "severity": "critical",
		},
	}), types.ParseTriageArtifact)
	if _, err := client.SaveTriage(ctx, SaveTriageInput{Binding: binding, Artifact: first}); err != nil {
		t.Fatal(err)
	}
	second := mustParseArtifact(t, detailTriage(t, "2026-01-04", []any{
		map[string]any{
			"id": "f302", "title": "Re-adjudication", "orig_id": "FIND-001",
			"verdict": "undetermined", "severity": "low",
		},
	}), types.ParseTriageArtifact)
	if _, err := client.SaveTriage(ctx, SaveTriageInput{Binding: binding, Artifact: second}); err != nil {
		t.Fatal(err)
	}

	rows, err := client.QueryCurrentFindingDetail(ctx, []string{scope})
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 5 {
		t.Fatalf("a second triage fanned findings out: %d rows, want 5", len(rows))
	}
	byID := make(map[string]CurrentFindingDetailRow, len(rows))
	for _, row := range rows {
		byID[row.FindingID] = row
	}
	if verdict := byID["FIND-001"].TriageVerdict; verdict == nil || *verdict != "undetermined" {
		t.Errorf("FIND-001 triage_verdict = %v, want undetermined (newest triage wins)", verdict)
	}
	// The newest triage is the context's triage: a finding it does not cover
	// reads as untriaged rather than mixing verdicts across generations.
	if verdict := byID["FIND-003"].TriageVerdict; verdict != nil {
		t.Errorf("FIND-003 triage_verdict = %v, want nil (uncovered by the newest triage)", verdict)
	}
}

func TestCurrentFindingDetail(t *testing.T) {
	checkCurrentFindingDetail(t, openTestStorage(t))
}

func TestCurrentFindingDetailNewestTriage(t *testing.T) {
	checkCurrentFindingDetailNewestTriage(t, openTestStorage(t))
}

func TestPostgresCurrentFindingDetail(t *testing.T) {
	client, _ := openPostgresStorage(t)
	checkCurrentFindingDetail(t, client)
}

func TestPostgresCurrentFindingDetailNewestTriage(t *testing.T) {
	client, _ := openPostgresStorage(t)
	checkCurrentFindingDetailNewestTriage(t, client)
}
