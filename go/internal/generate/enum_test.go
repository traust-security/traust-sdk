package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestIdentSafePascalCase(t *testing.T) {
	cases := map[string]string{
		"source+runtime":    "SourceRuntime",
		"repo:maintenance":  "RepoMaintenance",
		"not-ready":         "NotReady",
		"needs_manual_test": "NeedsManualTest",
		"E0":                "E0",
	}
	for in, want := range cases {
		if got := toPascalCase(identSafe(in)); got != want {
			t.Errorf("toPascalCase(identSafe(%q)) = %q, want %q", in, got, want)
		}
	}
}

func writeEnumDef(t *testing.T, dir, body string) string {
	t.Helper()
	path := filepath.Join(dir, "def.json")
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestGenerateEnumSanitisesValues(t *testing.T) {
	dir := t.TempDir()
	src := writeEnumDef(t, dir, `{"name": "assessment_basis", "values": ["source", "sbom-only", "source+runtime"]}`)
	if err := generateEnum(src, dir, "test-ref"); err != nil {
		t.Fatalf("generateEnum: %v", err)
	}
	out, err := os.ReadFile(filepath.Join(dir, "assessment_basis.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `AssessmentBasisSourceRuntime AssessmentBasis = "source+runtime"`) {
		t.Fatalf("sanitised constant missing:\n%s", out)
	}
}

func TestGenerateEnumRejectsConstantCollision(t *testing.T) {
	dir := t.TempDir()
	src := writeEnumDef(t, dir, `{"name": "clash", "values": ["a+b", "a-b"]}`)
	err := generateEnum(src, dir, "test-ref")
	if err == nil || !strings.Contains(err.Error(), "both map to constant") {
		t.Fatalf("want collision error, got %v", err)
	}
}

func TestGenerateEnumKeepsTheFirstCasesPublishedNames(t *testing.T) {
	dir := t.TempDir()
	src := writeEnumDef(t, dir, `{"name": "priority", "values": ["p0", "p1", "P0", "P1"]}`)
	if err := generateEnum(src, dir, "test-ref"); err != nil {
		t.Fatalf("generateEnum: %v", err)
	}
	out, err := os.ReadFile(filepath.Join(dir, "priority.go"))
	if err != nil {
		t.Fatal(err)
	}
	// lowercase came first: its constants keep their names and values, uppercase is added
	for _, want := range []string{
		`PriorityP0      Priority = "p0"`,
		`PriorityP1      Priority = "p1"`,
		`PriorityP0Upper Priority = "P0"`,
		`PriorityP1Upper Priority = "P1"`,
	} {
		if !strings.Contains(strings.Join(strings.Fields(string(out)), " "), strings.Join(strings.Fields(want), " ")) {
			t.Fatalf("missing %q:\n%s", want, out)
		}
	}
}

func TestGenerateEnumSuffixesLowercaseCaseVariants(t *testing.T) {
	dir := t.TempDir()
	src := writeEnumDef(t, dir, `{"name": "effort", "values": ["S", "M", "s", "m", "xl"]}`)
	if err := generateEnum(src, dir, "test-ref"); err != nil {
		t.Fatalf("generateEnum: %v", err)
	}
	out, err := os.ReadFile(filepath.Join(dir, "effort.go"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`EffortS      Effort = "S"`,
		`EffortSLower Effort = "s"`,
		`EffortMLower Effort = "m"`,
		`EffortXlLower Effort = "xl"`,
	} {
		if !strings.Contains(strings.Join(strings.Fields(string(out)), " "), strings.Join(strings.Fields(want), " ")) {
			t.Fatalf("missing %q:\n%s", want, out)
		}
	}
}
