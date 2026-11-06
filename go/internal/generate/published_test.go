package main

import (
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

const publishedReport = `{
  "$schema":"https://json-schema.org/draft/2020-12/schema",
  "$id":"https://example.test/report.schema.json",
  "type":"object","properties":{
    "findings":{"type":"array","items":{"$ref":"#/$defs/finding"}},
    "origin":{"oneOf":[{"$ref":"#/$defs/location"},{"type":"null"}]}
  },
  "$defs":{
    "location":{"type":"object","required":["path"],"properties":{"path":{"type":"string"}}},
    "finding":{"type":"object","required":["title"],"properties":{
      "title":{"type":"string"},
      "locations":{"type":"array","items":{"$ref":"#/$defs/location"}},
      "note":{"oneOf":[{"$ref":"#/$defs/location"},{"type":"null"}]}
    }}
  }
}`

// The later schemas sort first and claim the published report's names.
var contesting = map[string]string{
	"adapter-result.schema.json": `{
  "$schema":"https://json-schema.org/draft/2020-12/schema",
  "$id":"https://example.test/adapter-result.schema.json",
  "type":"object","required":["findings"],
  "properties":{"findings":{"type":"array","items":{"$ref":"#/$defs/adapter_finding"}}},
  "$defs":{"adapter_finding":{"type":"object","required":["title"],"properties":{
    "title":{"type":"string"},
    "locations":{"type":"array","items":{"$ref":"report.schema.json#/$defs/location"}},
    "note":{"oneOf":[{"$ref":"report.schema.json#/$defs/location"},{"type":"null"}]}
  }}}
}`,
	"code-facts.schema.json": `{
  "$schema":"https://json-schema.org/draft/2020-12/schema",
  "$id":"https://example.test/code-facts.schema.json",
  "type":"object","required":["symbols","coverage"],
  "properties":{
    "symbols":{"type":["array","null"],"items":{"$ref":"#/$defs/location"}},
    "coverage":{"oneOf":[{"$ref":"#/$defs/coverage"},{"type":"null"}]}
  },
  "$defs":{
    "location":{"type":"object","required":["line"],"properties":{"line":{"type":"integer"}}},
    "coverage":{"type":"object","required":["lines"],"properties":{"lines":{"type":"integer"}}}
  }
}`,
}

func generateFixture(t *testing.T, sources map[string]string, typesDir string) {
	t.Helper()
	schemaDir := t.TempDir()
	for name, content := range sources {
		if err := os.WriteFile(filepath.Join(schemaDir, name), []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	schemas, err := LoadSchemas(schemaDir)
	if err != nil {
		t.Fatal(err)
	}
	if err := ResolveRefs(schemas); err != nil {
		t.Fatal(err)
	}
	if err := GenerateTypes(schemas, nil, "example.test/enums", typesDir, "published-fixture"); err != nil {
		t.Fatal(err)
	}
}

func generatedSource(t *testing.T, typesDir, file string) string {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(typesDir, file))
	if err != nil {
		t.Fatal(err)
	}
	return strings.Join(strings.Fields(string(content)), " ")
}

func TestRegenerationKeepsPublishedShapes(t *testing.T) {
	typesDir := t.TempDir()
	generateFixture(t, map[string]string{"report.schema.json": publishedReport}, typesDir)
	// a field published before the generator typed nullable sections
	report := filepath.Join(typesDir, "report.go")
	published, err := os.ReadFile(report)
	if err != nil {
		t.Fatal(err)
	}
	legacy := regexp.MustCompile(`Origin\s+\*Location`).ReplaceAllString(string(published), "Origin interface{}")
	if legacy == string(published) {
		t.Fatalf("fixture lost its nullable field:\n%s", published)
	}
	if err := os.WriteFile(report, []byte(legacy), 0o600); err != nil {
		t.Fatal(err)
	}

	sources := map[string]string{"report.schema.json": publishedReport}
	maps.Copy(sources, contesting)
	generateFixture(t, sources, typesDir)

	for file, wants := range map[string][]string{
		"report.go": {
			"type Location struct", // the published name stays with its schema
			"Locations []Location",
			"Origin interface{}", // a published shape is never retyped
		},
		"adapter-result.go": {
			"type AdapterFinding = Finding", // identical to a published struct: an alias
		},
		"code-facts.go": {
			"type CodeFactsLocation struct", // the newer claimant takes a prefixed name
			"Symbols []CodeFactsLocation",   // a nullable section keeps its concrete type
			"Coverage *Coverage",
		},
	} {
		source := generatedSource(t, typesDir, file)
		for _, want := range wants {
			if !strings.Contains(source, strings.Join(strings.Fields(want), " ")) {
				t.Fatalf("%s: missing %q:\n%s", file, want, source)
			}
		}
	}

	// regenerating over its own output changes nothing
	before := map[string]string{}
	for _, file := range []string{"report.go", "adapter-result.go", "code-facts.go"} {
		before[file] = generatedSource(t, typesDir, file)
	}
	generateFixture(t, sources, typesDir)
	for file, was := range before {
		if now := generatedSource(t, typesDir, file); now != was {
			t.Fatalf("%s changed on a second regeneration:\nwas %s\nnow %s", file, was, now)
		}
	}
}
