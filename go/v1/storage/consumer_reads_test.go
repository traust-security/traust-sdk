package storage

// consumer_reads_test.go — the reads the interaction-boundary audit
// (traust-foundry#97, class 3) retires behind this package: the
// subject-scoped binding lookup, the projection-completeness count, and
// the baseline accessor. Same package as the reads so the SQLite leg of
// the storage suite covers them; the Postgres integration leg runs the
// package's shared checks.

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/types"
)

func documentFindingCount(t *testing.T, payload []byte) int {
	t.Helper()
	var document struct {
		Findings []json.RawMessage `json:"findings"`
	}
	if err := json.Unmarshal(payload, &document); err != nil {
		t.Fatal(err)
	}
	return len(document.Findings)
}

func TestFindBindingBySubjectResolvesCompanionArtifact(t *testing.T) {
	checkFindBindingBySubject(t, openTestStorage(t))
}

func TestFindBindingBySubjectAbsentAndInvalid(t *testing.T) {
	checkFindBindingBySubjectAbsentAndInvalid(t, openTestStorage(t))
}

func TestFindBindingBySubjectStaysInsideItsScope(t *testing.T) {
	checkFindBindingBySubjectStaysInsideItsScope(t, openTestStorage(t))
}

func TestCountProjectionRowsMatchesFindings(t *testing.T) {
	checkCountProjectionRows(t, openTestStorage(t))
}

func TestStorageRevisionFailsClosed(t *testing.T) {
	checkStorageRevisionFailsClosed(t, openTestStorage(t))
}

// The checks below run under both dialects: the SQLite legs above and the
// PostgreSQL legs in postgres_test.go, so the dialect-specific query text
// (the traust_storage prefixes, the information_schema shape read) is
// exercised on the engine that serves it.

func checkFindBindingBySubject(t *testing.T, client *Client) {
	t.Helper()
	ctx := context.Background()
	report, err := client.SaveReport(ctx, SaveReportInput{Binding: bindingFor("report"), References: []string{"s3://results/consumer-reads/report.json"}, Artifact: mustParseArtifact(t, detailReport(t), types.ParseReportArtifact)})
	if err != nil {
		t.Fatal(err)
	}
	threat, err := client.saveNamed(ctx, "threat-model", ratedThreatModel(t), bindingFor("threat-model"))
	if err != nil {
		t.Fatal(err)
	}
	const subject, run = "sci:inventory-item:42", "sci:scan-result:7"

	record, err := client.FindBindingBySubject(ctx, "local", subject, run, "threat-model")
	if err != nil {
		t.Fatal(err)
	}
	if record.BindingID != threat.BindingID || record.ArtifactName != "threat-model" || record.Digest != threat.Digest {
		t.Fatalf("threat-model binding = %+v, want the saved companion", record)
	}
	if record.ByteSize <= 0 {
		t.Fatalf("threat-model binding = %+v, want the evidence size", record)
	}
	if record.Binding.SubjectID == nil || *record.Binding.SubjectID != subject || record.Binding.RunID == nil || *record.Binding.RunID != run {
		t.Fatalf("threat-model binding context = %+v, want the requested subject and run", record.Binding)
	}

	reportRecord, err := client.FindBindingBySubject(ctx, "local", subject, run, "report")
	if err != nil {
		t.Fatal(err)
	}
	if reportRecord.BindingID != report.BindingID || reportRecord.ArtifactName != "report" {
		t.Fatalf("report binding = %+v, want the saved report", reportRecord)
	}
	if len(reportRecord.References) != 1 || reportRecord.References[0] != "s3://results/consumer-reads/report.json" {
		t.Fatalf("report references = %v, want the registered location", reportRecord.References)
	}
}

func checkFindBindingBySubjectAbsentAndInvalid(t *testing.T, client *Client) {
	t.Helper()
	ctx := context.Background()
	// The check owns its binding context so it stays absent under the
	// shared PostgreSQL client the other checks already populated.
	const subject, run = "sci:inventory-item:absent", "sci:scan-result:absent"
	absent := Binding{ScopeID: "local", SubjectID: stringPointer(subject), RunID: stringPointer(run)}
	if _, err := client.SaveReport(ctx, SaveReportInput{Binding: absent, References: []string{"s3://results/consumer-reads/absent/report.json"}, Artifact: mustParseArtifact(t, detailReport(t), types.ParseReportArtifact)}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.FindBindingBySubject(ctx, absent.ScopeID, subject, run, "threat-model"); !errors.Is(err, ErrBindingNotFound) {
		t.Fatalf("absent companion = %v, want ErrBindingNotFound", err)
	}
	for _, input := range [][4]string{
		{"", "sci:inventory-item:42", "sci:scan-result:7", "threat-model"},
		{"local", "", "sci:scan-result:7", "threat-model"},
		{"local", "sci:inventory-item:42", "", "threat-model"},
		{"local", "sci:inventory-item:42", "sci:scan-result:7", ""},
		{"bad\x00scope", "sci:inventory-item:42", "sci:scan-result:7", "threat-model"},
		{"bad\x00subject", "sci:inventory-item:42", "sci:scan-result:7", "threat-model"},
	} {
		if _, err := client.FindBindingBySubject(ctx, input[0], input[1], input[2], input[3]); !errors.Is(err, ErrInvalidIdentifier) {
			t.Fatalf("input %q = %v, want ErrInvalidIdentifier", input, err)
		}
	}
}

// checkFindBindingBySubjectStaysInsideItsScope proves the lookup cannot
// cross scope boundaries: two scopes reusing the caller-owned subject and
// run, the other scope's companion newer, and each scope still resolves
// its own binding.
func checkFindBindingBySubjectStaysInsideItsScope(t *testing.T, client *Client) {
	t.Helper()
	ctx := context.Background()
	const subject, run = "sci:inventory-item:42", "sci:scan-result:7"
	local, err := client.saveNamed(ctx, "threat-model", ratedThreatModel(t), Binding{ScopeID: "local", SubjectID: stringPointer(subject), RunID: stringPointer(run)})
	if err != nil {
		t.Fatal(err)
	}
	other, err := client.saveNamed(ctx, "threat-model", ratedThreatModel(t), Binding{ScopeID: "other", SubjectID: stringPointer(subject), RunID: stringPointer(run)})
	if err != nil {
		t.Fatal(err)
	}
	// The other scope's companion is definitively newer, so an unscoped
	// newest-wins lookup would steal it.
	bump := "UPDATE artifact_binding SET bound_at = datetime(bound_at, '+1 hour') WHERE binding_id = ?"
	if client.store.dialect == dialectPostgres {
		bump = "UPDATE traust_storage.artifact_binding SET bound_at = bound_at + interval '1 hour' WHERE binding_id = $1"
	}
	if _, err := sqlDB(client).ExecContext(ctx, bump, other.BindingID); err != nil {
		t.Fatal(err)
	}
	record, err := client.FindBindingBySubject(ctx, "local", subject, run, "threat-model")
	if err != nil || record.BindingID != local.BindingID {
		t.Fatalf("local lookup = (%s, %v), want its own binding %s", record.BindingID, err, local.BindingID)
	}
	record, err = client.FindBindingBySubject(ctx, "other", subject, run, "threat-model")
	if err != nil || record.BindingID != other.BindingID {
		t.Fatalf("other lookup = (%s, %v), want its own binding %s", record.BindingID, err, other.BindingID)
	}
	if _, err := client.FindBindingBySubject(ctx, "missing-scope", subject, run, "threat-model"); !errors.Is(err, ErrBindingNotFound) {
		t.Fatalf("absent scope = %v, want ErrBindingNotFound", err)
	}
}

func checkCountProjectionRows(t *testing.T, client *Client) {
	t.Helper()
	ctx := context.Background()
	binding := bindingFor("report")
	reportPayload := detailReport(t)
	if _, err := client.SaveReport(ctx, SaveReportInput{Binding: binding, Artifact: mustParseArtifact(t, reportPayload, types.ParseReportArtifact)}); err != nil {
		t.Fatal(err)
	}
	triagePayload := detailTriage(t, "2026-01-03", []any{
		map[string]any{"id": "f101", "title": "Confirmed", "orig_id": "FIND-001", "verdict": "true_positive", "severity": "high"},
		map[string]any{"id": "f102", "title": "Refuted", "orig_id": "FIND-003", "verdict": "false_positive", "severity": "low"},
	})
	if _, err := client.SaveTriage(ctx, SaveTriageInput{Binding: binding, Artifact: mustParseArtifact(t, triagePayload, types.ParseTriageArtifact)}); err != nil {
		t.Fatal(err)
	}
	// The report and the triage are distinct bindings of one subject/run:
	// resolve each through its own name, then count its projection.
	reportRecord, err := client.FindBindingBySubject(ctx, binding.ScopeID, *binding.SubjectID, *binding.RunID, "report")
	if err != nil {
		t.Fatal(err)
	}
	triageRecord, err := client.FindBindingBySubject(ctx, binding.ScopeID, *binding.SubjectID, *binding.RunID, "triage")
	if err != nil {
		t.Fatal(err)
	}
	reportCount, err := client.CountProjectionRows(ctx, reportRecord.BindingID, "report")
	if err != nil {
		t.Fatal(err)
	}
	if want := int64(documentFindingCount(t, reportPayload)); reportCount != want {
		t.Fatalf("report projection rows = %d, want %d", reportCount, want)
	}
	triageCount, err := client.CountProjectionRows(ctx, triageRecord.BindingID, "triage")
	if err != nil {
		t.Fatal(err)
	}
	if want := int64(documentFindingCount(t, triagePayload)); triageCount != want {
		t.Fatalf("triage projection rows = %d, want %d", triageCount, want)
	}
	if _, err := client.CountProjectionRows(ctx, reportRecord.BindingID, "threat-model"); !errors.Is(err, ErrUnknownArtifact) {
		t.Fatalf("unprojected kind = %v, want ErrUnknownArtifact", err)
	}
	if _, err := client.CountProjectionRows(ctx, "not-a-digest", "report"); !errors.Is(err, ErrBindingNotFound) {
		t.Fatalf("invalid binding id = %v, want ErrBindingNotFound", err)
	}
}

func checkStorageRevisionFailsClosed(t *testing.T, client *Client) {
	t.Helper()
	ctx := context.Background()
	baseline, err := client.StorageRevision(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if baseline.ContractVersion != storageFormatVersion || baseline.Revision != contractRevision {
		t.Fatalf("baseline = %+v, want the stamped %s/%d", baseline, storageFormatVersion, contractRevision)
	}
	meta := "traust_storage_meta"
	if client.store.dialect == dialectPostgres {
		meta = "traust_storage.traust_storage_meta"
	}
	if _, err := sqlDB(client).ExecContext(ctx, "UPDATE "+meta+" SET revision = $1 WHERE id = 1", contractRevision-1); err != nil {
		t.Fatal(err)
	}
	if _, err := client.StorageRevision(ctx); !errors.Is(err, ErrIncompatibleRevision) {
		t.Fatalf("revision mismatch = %v, want ErrIncompatibleRevision", err)
	}
	if _, err := sqlDB(client).ExecContext(ctx, "UPDATE "+meta+" SET revision = $1 WHERE id = 1", contractRevision); err != nil {
		t.Fatal(err)
	}
	// A schema that claims the revision but not the shape it implies: the
	// evidence table lost its size column (and grew an inline payload one,
	// the pre-revision layout).
	evidence := "artifact_evidence"
	if client.store.dialect == dialectPostgres {
		evidence = "traust_storage.artifact_evidence"
	}
	if _, err := sqlDB(client).ExecContext(ctx, "ALTER TABLE "+evidence+" RENAME COLUMN byte_size TO size"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.StorageRevision(ctx); !errors.Is(err, ErrIncompatibleDatabase) {
		t.Fatalf("evidence shape mismatch = %v, want ErrIncompatibleDatabase", err)
	}
	if _, err := sqlDB(client).ExecContext(ctx, "ALTER TABLE "+evidence+" RENAME COLUMN size TO byte_size"); err != nil {
		t.Fatal(err)
	}
	if _, err := sqlDB(client).ExecContext(ctx, "DELETE FROM "+meta+" WHERE id = 1"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.StorageRevision(ctx); !errors.Is(err, ErrNotFound) {
		t.Fatalf("absent baseline = %v, want ErrNotFound", err)
	}
}
