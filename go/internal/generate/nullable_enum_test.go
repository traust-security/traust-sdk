package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedNullableEnumArtifactRoundTrip(t *testing.T) {
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
	registry := `{"name":"effort","values":["s","m"]}`
	source := writeEnumDef(t, registryDir, registry)
	if err := generateEnum(source, enumDir, "nullable-fixture"); err != nil {
		t.Fatal(err)
	}
	schema := `{
    "$schema":"https://json-schema.org/draft/2020-12/schema",
    "$id":"https://example.com/traust-contracts/schemas/v1/nullable.schema.json",
    "type":"object",
    "required":["effort","reference_effort"],
    "additionalProperties":false,
    "properties":{
        "effort":{"enum":["s","m",null]},
        "reference_effort":{"$ref":"#/$defs/holder/properties/effort"}
    },
    "$defs":{
        "effort":{"type":["string","null"],"enum":["s","m",null]},
        "holder":{"type":"object","properties":{"effort":{"$ref":"#/$defs/effort"}}}
    }
}`
	if err := os.WriteFile(filepath.Join(schemaDir, "nullable.schema.json"), []byte(schema), 0o600); err != nil {
		t.Fatal(err)
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
	const enumsPackage = "github.com/traust-security/traust-sdk/go/v1/enums"
	if err := GenerateTypes(schemas, mappings, enumsPackage, typesDir, "nullable-fixture"); err != nil {
		t.Fatal(err)
	}
	if err := GenerateSchemasEmbed(schemaDir, filepath.Join(validateDir, "schemas_gen.go"), "nullable-fixture"); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		filepath.Join("..", "..", "go.mod"):                        "go.mod",
		filepath.Join("..", "..", "go.sum"):                        "go.sum",
		filepath.Join("..", "..", "v1", "types", "artifact.go"):    filepath.Join("v1", "types", "artifact.go"),
		filepath.Join("..", "..", "v1", "validate", "validate.go"): filepath.Join("v1", "validate", "validate.go"),
	}
	for source, destination := range files {
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
func main() {
    artifact, err := types.ParseNullableArtifact([]byte("{\"effort\":null,\"reference_effort\":null}"))
    if err != nil { panic(err) }
    value, err := artifact.Value()
    if err != nil { panic(err) }
    restored, err := types.EncodeNullableArtifact(value)
    if err != nil { panic(err) }
    var nulls map[string]any
    if err := json.Unmarshal(restored.Payload(), &nulls); err != nil { panic(err) }
    size := enums.EffortS
    value.Effort = &size
    value.ReferenceEffort = &size
    known, err := types.EncodeNullableArtifact(value)
    if err != nil { panic(err) }
    var sizes map[string]any
    if err := json.Unmarshal(known.Payload(), &sizes); err != nil { panic(err) }
    invalid := enums.Effort("urgent")
    value.Effort = &invalid
    _, invalidErr := types.EncodeNullableArtifact(value)
    result := struct {
        Nulls map[string]any ` + "`json:\"nulls\"`" + `
        Sizes map[string]any ` + "`json:\"sizes\"`" + `
        InvalidRejected bool ` + "`json:\"invalid_rejected\"`" + `
    }{nulls, sizes, invalidErr != nil}
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
		t.Fatalf("generated nullable consumer: %v\n%s", err, output)
	}
	var result struct {
		Nulls           map[string]any `json:"nulls"`
		Sizes           map[string]any `json:"sizes"`
		InvalidRejected bool           `json:"invalid_rejected"`
	}
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"effort", "reference_effort"} {
		if value, present := result.Nulls[field]; !present || value != nil {
			t.Fatalf("required unknown size was lost: %v", result.Nulls)
		}
		if result.Sizes[field] != "s" {
			t.Fatalf("known size changed: %v", result.Sizes)
		}
	}
	if !result.InvalidRejected {
		t.Fatal("invalid enum escaped the actual artifact write gate")
	}
}
