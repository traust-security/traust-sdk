package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedCrossSchemaPropertyReferencesRoundTrip(t *testing.T) {
	dir := t.TempDir()
	registryDir := filepath.Join(dir, "registry")
	schemaDir := filepath.Join(dir, "schemas")
	enumDir := filepath.Join(dir, "v1", "enums")
	typesDir := filepath.Join(dir, "v1", "types")
	validateDir := filepath.Join(dir, "v1", "validate")
	for _, path := range []string{registryDir, schemaDir, enumDir, validateDir} {
		if err := os.MkdirAll(path, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	for name, definition := range map[string]string{
		"severity.json": `{"name":"severity","values":["high","low"]}`,
		"validity.json": `{"name":"validity","values":["confirmed","false_positive"]}`,
	} {
		source := filepath.Join(registryDir, name)
		if err := os.WriteFile(source, []byte(definition), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := generateEnum(source, enumDir, "property-ref-fixture"); err != nil {
			t.Fatal(err)
		}
	}
	sources := map[string]string{
		"report.schema.json": `{
  "$schema":"https://json-schema.org/draft/2020-12/schema",
  "$id":"https://example.test/report.schema.json",
  "type":"object","properties":{"findings":{"type":"array","items":{"$ref":"#/$defs/finding"}}},
  "$defs":{
    "severity":{"type":"string","enum":["high","low"]},
    "validity":{"type":"string","enum":["confirmed","false_positive"]},
    "disposition":{"type":"object","required":["validity"],"properties":{"validity":{"$ref":"#/$defs/validity"}}},
    "location":{"type":"object","required":["path"],"properties":{"path":{"type":"string"}}},
    "finding":{"type":"object","required":["severity"],"properties":{
      "severity":{"$ref":"#/$defs/severity"},
      "locations":{"type":"array","items":{"$ref":"#/$defs/location"}},
      "cvss":{"type":"object","required":["score"],"properties":{"score":{"type":"number"}}},
      "dependency":{"type":"object","required":["module"],"properties":{"module":{"type":"string"}}},
      "disposition":{"$ref":"#/$defs/disposition"}
    }}
  }
}`,
		"adapter-result.schema.json": `{
  "$schema":"https://json-schema.org/draft/2020-12/schema",
  "$id":"https://example.test/adapter-result.schema.json",
  "type":"object","required":["findings"],"additionalProperties":false,
  "properties":{"findings":{"type":"array","items":{"$ref":"#/$defs/adapter_finding"}}},
  "$defs":{
    "disposition":{"type":"object","properties":{"description":{"type":"string"}}},
    "adapter_finding":{"type":"object","required":["severity","locations"],"additionalProperties":false,"properties":{
      "severity":{"$ref":"report.schema.json#/$defs/finding/properties/severity"},
      "locations":{"$ref":"report.schema.json#/$defs/finding/properties/locations"},
      "cvss":{"$ref":"report.schema.json#/$defs/finding/properties/cvss"},
      "dependency":{"$ref":"report.schema.json#/$defs/finding/properties/dependency"},
      "disposition":{"$ref":"report.schema.json#/$defs/finding/properties/disposition"}
    }}
  }
}`,
	}
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
	mappings, err := LoadEnumMappings(registryDir)
	if err != nil {
		t.Fatal(err)
	}
	if err := GenerateTypes(schemas, mappings, "github.com/traust-security/traust-sdk/go/v1/enums", typesDir, "property-ref-fixture"); err != nil {
		t.Fatal(err)
	}
	if err := GenerateSchemasEmbed(schemaDir, filepath.Join(validateDir, "schemas_gen.go"), "property-ref-fixture"); err != nil {
		t.Fatal(err)
	}
	for source, destination := range map[string]string{
		filepath.Join("..", "..", "go.mod"):                        "go.mod",
		filepath.Join("..", "..", "go.sum"):                        "go.sum",
		filepath.Join("..", "..", "v1", "types", "artifact.go"):    filepath.Join("v1", "types", "artifact.go"),
		filepath.Join("..", "..", "v1", "validate", "validate.go"): filepath.Join("v1", "validate", "validate.go"),
	} {
		data, err := os.ReadFile(source)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, destination), data, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	consumer := `package main
import (
  "encoding/json"
  "os"
  "strings"
  "github.com/traust-security/traust-sdk/go/v1/enums"
  "github.com/traust-security/traust-sdk/go/v1/types"
)
func requireSeverity(value enums.Severity) bool { return value == enums.SeverityHigh }
func requireValidity(value enums.Validity) bool { return value == enums.ValidityConfirmed }
func main() {
  payload := []byte("{\"findings\":[{\"severity\":\"high\",\"locations\":[{\"path\":\"source.go\"}],\"cvss\":{\"score\":3.1},\"dependency\":{\"module\":\"original-module\"},\"disposition\":{\"validity\":\"confirmed\"}}]}")
  artifact, err := types.ParseAdapterResultArtifact(payload)
  if err != nil { panic(err) }
  value, err := artifact.Value()
  if err != nil { panic(err) }
  if !requireSeverity(value.Findings[0].Severity) || !requireValidity(value.Findings[0].Disposition.Validity) { panic("referenced enum changed") }
  if value.Findings[0].Locations[0].Path != "source.go" || value.Findings[0].Cvss.Score != 3.1 || value.Findings[0].Dependency.Module != "original-module" { panic("referenced observations lost") }
  absentArtifact, err := types.ParseAdapterResultArtifact([]byte("{\"findings\":[{\"severity\":\"low\",\"locations\":[]}]}"))
  if err != nil { panic(err) }
  absent, err := absentArtifact.Value()
  if err != nil { panic(err) }
  if absent.Findings[0].Cvss != nil || absent.Findings[0].Dependency != nil || absent.Findings[0].Disposition != nil { panic("optional referenced metadata lost pointer semantics") }
  value.Findings[0].Cvss.Score = 7.2
  value.Findings[0].Dependency.Module = "updated-module"
  value.Findings[0].Disposition.Validity = enums.ValidityFalsePositive
  value.Findings[0].Locations[0].Path = "updated.go"
  encoded, err := types.EncodeAdapterResultArtifact(value)
  if err != nil { panic(err) }
  roundTrip, err := encoded.Value()
  if err != nil { panic(err) }
  if roundTrip.Findings[0].Disposition.Validity != enums.ValidityFalsePositive { panic("nested disposition mutation lost") }
  invalid := []byte(strings.Replace(string(payload), "\"high\"", "\"urgent\"", 1))
  if _, err := types.ParseAdapterResultArtifact(invalid); err == nil { panic("invalid severity passed the parse gate") }
  value.Findings[0].Severity = enums.Severity("urgent")
  if _, err := types.EncodeAdapterResultArtifact(value); err == nil { panic("invalid severity passed the encode gate") }
  value.Findings[0].Severity = enums.SeverityHigh
  value.Findings[0].Disposition.Validity = enums.Validity("unknown")
  if _, err := types.EncodeAdapterResultArtifact(value); err == nil { panic("invalid nested validity passed the encode gate") }
  if err := json.NewEncoder(os.Stdout).Encode(json.RawMessage(encoded.Payload())); err != nil { panic(err) }
}
`
	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(consumer), 0o600); err != nil {
		t.Fatal(err)
	}
	command := exec.Command("go", "run", ".")
	command.Dir = dir
	command.Env = append(os.Environ(), "GOPROXY=off")
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("generated property-reference consumer: %v\n%s", err, output)
	}
	var result struct {
		Findings []struct {
			Severity    string
			Locations   []struct{ Path string }
			Cvss        struct{ Score float64 }
			Dependency  struct{ Module string }
			Disposition struct{ Validity string }
		}
	}
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatal(err)
	}
	if len(result.Findings) != 1 || len(result.Findings[0].Locations) != 1 {
		t.Fatalf("referenced finding or location lost: %s", output)
	}
	finding := result.Findings[0]
	if finding.Severity != "high" || finding.Locations[0].Path != "updated.go" || finding.Cvss.Score != 7.2 || finding.Dependency.Module != "updated-module" || finding.Disposition.Validity != "false_positive" {
		t.Fatalf("referenced metadata did not survive typed mutation: %s", output)
	}
}

func TestPropertyReferencePointers(t *testing.T) {
	cvss := &SchemaNode{Type: "object", Properties: map[string]*SchemaNode{"score": {Type: "number"}}}
	location := &SchemaNode{Type: "object", Properties: map[string]*SchemaNode{"path": {Type: "string"}}}
	dependency := &SchemaNode{Type: "object", Properties: map[string]*SchemaNode{"module": {Type: "string"}}}
	sf := &SchemaFile{
		Filename: "report.schema.json",
		Defs: map[string]*SchemaNode{
			"finding": {
				Type: "object",
				Properties: map[string]*SchemaNode{
					"cvss":       cvss,
					"locations":  {Type: "array", Items: location},
					"dependency": {Type: "object", AdditionalProperties: dependency},
				},
			},
			"finding/escaped~": {Properties: map[string]*SchemaNode{"cvss/escaped~": cvss}},
		},
	}
	schemas := map[string]*SchemaFile{sf.Filename: sf}
	registry := newTypeRegistry()
	registry.build(schemas)
	for _, tc := range []struct {
		pointer string
		node    *SchemaNode
		name    string
	}{
		{"#/$defs/finding/properties/cvss", cvss, "FindingCvss"},
		{"#/$defs/finding/properties/locations/items", location, "FindingLocationsItem"},
		{"#/definitions/finding/properties/dependency/additionalProperties", dependency, "FindingDependencyEntry"},
		{"#/$defs/finding~1escaped~0/properties/cvss~1escaped~0", cvss, ""},
	} {
		t.Run(tc.pointer, func(t *testing.T) {
			node, source, err := lookupRef(tc.pointer, sf.Filename, schemas)
			if err != nil {
				t.Fatal(err)
			}
			if node != tc.node || source != sf.Filename {
				t.Fatalf("resolved pointer to %p in %q, want %p in %q", node, source, tc.node, sf.Filename)
			}
			if tc.name != "" {
				if name := registry.structNameForRef(tc.pointer, sf.Filename); name != tc.name {
					t.Fatalf("generated name %q, want %q", name, tc.name)
				}
			}
		})
	}
	for _, ref := range []string{
		"#/$defs/finding/properties",
		"#/$defs/finding/properties/missing",
		"#/$defs/finding/properties/cvss/items",
		"#/$defs/finding/properties/cvss/unsupported",
	} {
		t.Run(ref, func(t *testing.T) {
			if _, _, err := lookupRef(ref, sf.Filename, schemas); err == nil {
				t.Fatal("invalid property pointer was accepted")
			}
		})
	}
}
