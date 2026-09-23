package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func TestGenerateOperationsUsesTypedArtifacts(t *testing.T) {
	storage := writeMinimalCanonicalStorage(t)
	schemas := t.TempDir()
	definitions := map[string]string{
		"triage":        `{}`,
		"vuln-findings": `{}`,
		"report":        `{"type":"object","properties":{"title":{"type":"string"}}}`,
	}
	for name, schema := range definitions {
		if err := os.WriteFile(filepath.Join(schemas, name+".schema.json"), []byte(schema), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	outputDir := t.TempDir()
	output := filepath.Join(outputDir, "operations.go")
	vocabulary := newProjectionVocabulary()
	if err := generateOperations(storage, schemas, output, "test-ref", vocabulary); err != nil {
		t.Fatal(err)
	}
	identifiers := filepath.Join(outputDir, "projection_identifiers.go")
	if err := generateProjectionIdentifiers(identifiers, "test-ref", vocabulary); err != nil {
		t.Fatal(err)
	}
	generated, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	text := string(generated)
	for _, want := range []string{
		"type SaveTriageInput struct",
		"Binding  Binding",
		"types.Artifact[types.Triage]",
		"bindingRequirements{subject: true, run: true, layer: false}",
		"GetTriage(ctx context.Context, bindingID string)",
		"SaveVulnFindings(ctx context.Context, input SaveVulnFindingsInput)",
		"func (s *sqlStore) projectReport",
		"return s.projectReportFindings(ctx, conn, state, value)",
		"state.bindingID",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("missing %q", want)
		}
	}
	for _, stale := range []string{"SaveInput", "map[string]any", "schemaNameFor", "projectionText(value.Title)"} {
		if strings.Contains(text, stale) {
			t.Errorf("generated operation surface contains %q", stale)
		}
	}
	if !strings.Contains(text, "title:") || !strings.Contains(text, "value.Title") {
		t.Error("generated projector does not assign typed text directly")
	}
	if strings.Contains(text, `projectionError("`) {
		t.Error("generated projector contains raw projection identifiers")
	}
	if !strings.Contains(text, "projectionError(projectionReport, projectionFieldRow, err)") {
		t.Error("generated projector does not use typed projection identifiers")
	}
	identifierSource, err := os.ReadFile(identifiers)
	if err != nil {
		t.Fatal(err)
	}
	identifierText := string(identifierSource)
	if !strings.Contains(identifierText, "projectionReport") || !strings.Contains(identifierText, `"report"`) {
		t.Error("generated projection vocabulary is missing report")
	}
	if strings.Contains(identifierText, "projectionFieldTitle") {
		t.Error("generated projection vocabulary includes an unused field")
	}
}

func TestGenerateSQLDeterministicAndDialectTripwire(t *testing.T) {
	source := writeMinimalCanonicalStorage(t)
	firstDir := t.TempDir()
	secondDir := t.TempDir()
	firstDDL := filepath.Join(firstDir, "schema.go")
	firstQueries := filepath.Join(firstDir, "queries.go")
	secondDDL := filepath.Join(secondDir, "schema.go")
	secondQueries := filepath.Join(secondDir, "queries.go")
	if err := generateSQL(source, firstDDL, firstQueries, "test-ref"); err != nil {
		t.Fatal(err)
	}
	if err := generateSQL(source, secondDDL, secondQueries, "test-ref"); err != nil {
		t.Fatal(err)
	}
	ddlSource, err := os.ReadFile(firstDDL)
	if err != nil {
		t.Fatal(err)
	}
	ddlText := string(ddlSource)
	for _, expected := range []string{`storageFormatVersion = "v1"`, `contractRevision = 7`, `storageBaselineID = "test-baseline"`} {
		if !strings.Contains(ddlText, expected) {
			t.Errorf("generated metadata missing %s", expected)
		}
	}
	namespaceAt := strings.Index(ddlText, "CREATE SCHEMA IF NOT EXISTS traust_storage")
	tableAt := strings.Index(ddlText, "CREATE TABLE IF NOT EXISTS traust_storage.report")
	if namespaceAt < 0 || tableAt < 0 || namespaceAt > tableAt {
		t.Error("generated PostgreSQL bootstrap does not create the storage namespace before tables")
	}
	if strings.Contains(ddlText, "projectionName") || strings.Contains(ddlText, "projectionField") {
		t.Error("generated schema contains projection identifiers")
	}

	for _, pair := range [][2]string{{firstDDL, secondDDL}, {firstQueries, secondQueries}} {
		a, err := os.ReadFile(pair[0])
		if err != nil {
			t.Fatal(err)
		}
		b, err := os.ReadFile(pair[1])
		if err != nil {
			t.Fatal(err)
		}
		if string(a) != string(b) {
			t.Fatal("generation is not deterministic")
		}
	}

	query := filepath.Join(source, "sqlite", "queries", "report.upsert.sql")
	data, err := os.ReadFile(query)
	if err != nil {
		t.Fatal(err)
	}
	data = []byte(strings.Replace(string(data), ":artifact_digest", ":wrong_digest", 1))
	if err := os.WriteFile(query, data, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := generateSQL(source, secondDDL, secondQueries, "test-ref"); err == nil || !strings.Contains(err.Error(), "placeholders differ") {
		t.Fatalf("expected dialect consistency failure, got %v", err)
	}
}

func TestGenerateSQLRejectsUnexpectedOneSidedQuery(t *testing.T) {
	source := writeMinimalCanonicalStorage(t)
	if err := os.Remove(filepath.Join(source, "sqlite", "queries", "report.upsert.sql")); err != nil {
		t.Fatal(err)
	}
	err := generateSQL(
		source,
		filepath.Join(t.TempDir(), "schema.go"),
		filepath.Join(t.TempDir(), "queries.go"),
		"test-ref",
	)
	if err == nil || !strings.Contains(err.Error(), "missing sqlite dialect") {
		t.Fatalf("expected missing SQLite dialect failure, got %v", err)
	}
}

func TestGeneratedStatementMethodsMatchCanonicalQueries(t *testing.T) {
	source := writeMinimalCanonicalStorage(t)
	outputDir := t.TempDir()
	ddlOutput := filepath.Join(outputDir, "schema.go")
	output := filepath.Join(outputDir, "queries.go")
	if err := generateSQL(source, ddlOutput, output, "test-ref"); err != nil {
		t.Fatal(err)
	}

	want := canonicalQueryMethods(t, source)
	file, err := parser.ParseFile(token.NewFileSet(), output, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Recv == nil {
			continue
		}
		if function.Name.IsExported() {
			t.Errorf("generated statement method is exported: %s", function.Name.Name)
		}
		receiver, ok := function.Recv.List[0].Type.(*ast.Ident)
		if ok && receiver.Name == "queries" {
			got = append(got, function.Name.Name)
		}
	}
	sort.Strings(got)
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("statement/query mismatch: methods=%v queries=%v", got, want)
	}

	generated, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	text := string(generated)
	for _, stale := range []string{"func exec", " args()", "type ReportUpsertParams"} {
		if strings.Contains(text, stale) {
			t.Errorf("generated statement surface contains stale shape %q", stale)
		}
	}
	if !strings.Contains(text, "func (q queries) reportUpsert") || !strings.Contains(text, "q.dialect == dialectPostgres") {
		t.Error("generated method does not select a frozen dialect literal")
	}
}

func canonicalQueryMethods(t *testing.T, source string) []string {
	t.Helper()
	set := map[string]bool{}
	for _, dialect := range []string{"postgres", "sqlite"} {
		entries, err := os.ReadDir(filepath.Join(source, dialect, "queries"))
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
				set[sqlIdent(strings.TrimSuffix(entry.Name(), ".sql"))] = true
			}
		}
	}
	methods := make([]string, 0, len(set))
	for method := range set {
		methods = append(methods, method)
	}
	sort.Strings(methods)
	return methods
}

func writeMinimalCanonicalStorage(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "metadata.json"), []byte(`{"contract_version":"v1","revision":7,"baseline_id":"test-baseline"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	profiles := `{"version":1,"artifacts":{"triage":{"class":"run-bound","required":["subject_id","run_id"],"projection":"report"},"vuln-findings":{"class":"run-bound","required":["subject_id","run_id"],"projection":"report"},"report":{"class":"run-bound","required":["subject_id","run_id"],"projection":"report"}}}`
	if err := os.WriteFile(filepath.Join(root, "profiles.json"), []byte(profiles), 0o600); err != nil {
		t.Fatal(err)
	}
	ddl := map[string]string{
		"postgres": "CREATE TABLE IF NOT EXISTS traust_storage.report (\n    binding_id TEXT NOT NULL,\n    artifact_digest TEXT NOT NULL,\n    title TEXT NOT NULL\n);\n",
		"sqlite":   "CREATE TABLE report (\n    binding_id TEXT NOT NULL,\n    artifact_digest TEXT NOT NULL,\n    title TEXT NOT NULL\n);\n",
	}
	queries := map[string]string{
		"postgres": "INSERT INTO traust_storage.report (binding_id, artifact_digest, title) VALUES (%(binding_id)s, %(artifact_digest)s, %(title)s);\n",
		"sqlite":   "INSERT INTO report (binding_id, artifact_digest, title) VALUES (:binding_id, :artifact_digest, :title);\n",
	}
	if err := os.MkdirAll(filepath.Join(root, "postgres"), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(
		filepath.Join(root, "postgres", "namespace.sql"),
		[]byte("CREATE SCHEMA IF NOT EXISTS traust_storage;\n"),
		0o600,
	); err != nil {
		t.Fatal(err)
	}
	for dialect, query := range queries {
		for _, section := range []string{"schema", "queries", "views"} {
			if err := os.MkdirAll(filepath.Join(root, dialect, section), 0o700); err != nil {
				t.Fatal(err)
			}
		}
		if err := os.WriteFile(filepath.Join(root, dialect, "schema", "report.sql"), []byte(ddl[dialect]), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, dialect, "queries", "report.upsert.sql"), []byte(query), 0o600); err != nil {
			t.Fatal(err)
		}
		if dialect == "postgres" {
			if err := os.WriteFile(filepath.Join(root, dialect, "queries", "artifact.lock.sql"), []byte("SELECT pg_advisory_xact_lock(1);\n"), 0o600); err != nil {
				t.Fatal(err)
			}
		}
	}
	return root
}
