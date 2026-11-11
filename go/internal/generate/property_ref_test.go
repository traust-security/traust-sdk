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
  "$defs":{"adapter_finding":{"type":"object","required":["severity","locations","cvss","dependency","disposition"],"additionalProperties":false,"properties":{
    "severity":{"$ref":"report.schema.json#/$defs/finding/properties/severity"},
    "locations":{"$ref":"report.schema.json#/$defs/finding/properties/locations"},
    "cvss":{"$ref":"report.schema.json#/$defs/finding/properties/cvss"},
    "dependency":{"$ref":"report.schema.json#/$defs/finding/properties/dependency"},
    "disposition":{"$ref":"report.schema.json#/$defs/finding/properties/disposition"}
  }}}
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
  "github.com/traust-security/traust-sdk/go/v1/enums"
  "github.com/traust-security/traust-sdk/go/v1/types"
)
func requireSeverity(value enums.Severity) bool { return value == enums.SeverityHigh }
func requireValidity(value enums.Validity) bool { return value == enums.ValidityConfirmed }
func main() {
  artifact, err := types.ParseAdapterResultArtifact([]byte("{\"findings\":[{\"severity\":\"high\",\"locations\":[{\"path\":\"source.go\"}],\"cvss\":{\"score\":3.1},\"dependency\":{\"module\":\"original-module\"},\"disposition\":{\"validity\":\"confirmed\"}}]}"))
  if err != nil { panic(err) }
  value, err := artifact.Value()
  if err != nil { panic(err) }
  if !requireSeverity(value.Findings[0].Severity) || !requireValidity(value.Findings[0].Disposition.Validity) { panic("referenced enum changed") }
  if value.Findings[0].Locations[0].Path != "source.go" || value.Findings[0].Cvss.Score != 3.1 || value.Findings[0].Dependency.Module != "original-module" { panic("referenced observations lost") }
  value.Findings[0].Cvss.Score = 7.2
  value.Findings[0].Dependency.Module = "updated-module"
  encoded, err := types.EncodeAdapterResultArtifact(value)
  if err != nil { panic(err) }
  value.Findings[0].Severity = enums.Severity("urgent")
  _, invalidErr := types.EncodeAdapterResultArtifact(value)
  result := struct { Payload json.RawMessage; InvalidRejected bool }{encoded.Payload(), invalidErr != nil}
  if err := json.NewEncoder(os.Stdout).Encode(result); err != nil { panic(err) }
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
		Payload struct {
			Findings []struct {
				Cvss       struct{ Score float64 }
				Dependency struct{ Module string }
			}
		}
		InvalidRejected bool
	}
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatal(err)
	}
	if result.Payload.Findings[0].Cvss.Score != 7.2 || result.Payload.Findings[0].Dependency.Module != "updated-module" {
		t.Fatalf("referenced observations did not survive typed mutation: %s", output)
	}
	if !result.InvalidRejected {
		t.Fatal("invalid enum passed the generated artifact gate")
	}
}
