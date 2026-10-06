package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedNormalizationBehavior(t *testing.T) {
	dir := t.TempDir()
	sourceDir := filepath.Join(dir, "registry")
	outDir := filepath.Join(dir, "v1", "enums")
	for _, path := range []string{sourceDir, outDir} {
		if err := os.MkdirAll(path, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	fixturePath := filepath.Join("testdata", "enum-normalization.json")
	fixture, err := os.ReadFile(fixturePath)
	if err != nil {
		t.Fatal(err)
	}
	var decoded struct {
		Registry []json.RawMessage `json:"registry"`
	}
	if err := json.Unmarshal(fixture, &decoded); err != nil {
		t.Fatal(err)
	}
	for i, document := range decoded.Registry {
		var definition enumDef
		if err := json.Unmarshal(document, &definition); err != nil {
			t.Fatal(err)
		}
		path := filepath.Join(sourceDir, definition.Name+".json")
		if err := os.WriteFile(path, document, 0o600); err != nil {
			t.Fatal(err)
		}
		if err := generateEnum(path, outDir); err != nil {
			t.Fatalf("registry %d: %v", i, err)
		}
	}
	if err := generateNormalization(sourceDir, outDir, "neutral-registry-fixture"); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		filepath.Join("..", "..", "v1", "enums", "normalization.go"): "normalization.go",
		filepath.Join("testdata", "normalization_test.go"):           "normalization_test.go",
		fixturePath: "enum-normalization.json",
	}
	for source, destination := range files {
		data, err := os.ReadFile(source)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(outDir, destination), data, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module github.com/traust-security/traust-sdk/go\n\ngo 1.26.0\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	command := exec.Command("go", "test", "./v1/enums", "-count=1", "-v")
	command.Dir = dir
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("generated reader behavior failed: %v\n%s", err, output)
	}
}

func TestMalformedDeprecationCannotBecomeDrop(t *testing.T) {
	for _, metadata := range []string{`{}`, `{"replaced_by":null}`} {
		t.Run(metadata, func(t *testing.T) {
			dir := t.TempDir()
			writeEnumDef(t, dir, `{"name":"colour","values":["blue"],"deprecated":{"blue":`+metadata+`}}`)
			err := generateNormalization(dir, dir, "neutral-registry-fixture")
			if err == nil {
				t.Fatalf("missing list was interpreted as a drop: %v", err)
			}
		})
	}
}
