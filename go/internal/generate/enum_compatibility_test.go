package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedEnumPreservesDeclaredLegacySymbols(t *testing.T) {
	for _, values := range [][]string{{"secure-code", "secure_code"}, {"secure_code", "secure-code"}} {
		t.Run(values[0], func(t *testing.T) {
			dir := t.TempDir()
			outDir := filepath.Join(dir, "generated")
			if err := os.Mkdir(outDir, 0o700); err != nil {
				t.Fatal(err)
			}
			definition := enumDef{
				Name:   "assessment_kind",
				Values: values,
				Deprecated: map[string]normalizationMetadata{
					"secure-code": {ReplacedBy: []normalizationPair{{Enum: "assessment_kind", Value: "secure_code"}}},
				},
			}
			body, err := json.Marshal(definition)
			if err != nil {
				t.Fatal(err)
			}
			if err := generateEnum(writeEnumDef(t, dir, string(body)), outDir, "test-ref"); err != nil {
				t.Fatal(err)
			}
			files := map[string]string{
				"go.mod": "module classifiercompat.test\n\ngo 1.26.0\n",
				"main.go": `package main
import (
    "encoding/json"
    "os"
    generated "classifiercompat.test/generated"
)
func main() {
    result := struct {
        Legacy generated.AssessmentKind ` + "`json:\"legacy\"`" + `
        Canonical generated.AssessmentKind ` + "`json:\"canonical\"`" + `
        Values []generated.AssessmentKind ` + "`json:\"values\"`" + `
    }{generated.AssessmentKindSecureCode, generated.AssessmentKindSecureCodeCanonical, generated.AssessmentKindValues()}
    if err := json.NewEncoder(os.Stdout).Encode(result); err != nil { panic(err) }
}
`,
			}
			for name, content := range files {
				if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600); err != nil {
					t.Fatal(err)
				}
			}
			command := exec.Command("go", "run", ".")
			command.Dir = dir
			output, err := command.CombinedOutput()
			if err != nil {
				t.Fatalf("generated consumer: %v\n%s", err, output)
			}
			var result struct {
				Legacy    string   `json:"legacy"`
				Canonical string   `json:"canonical"`
				Values    []string `json:"values"`
			}
			if err := json.Unmarshal(output, &result); err != nil {
				t.Fatal(err)
			}
			if result.Legacy != "secure-code" || result.Canonical != "secure_code" {
				t.Fatalf("public symbol values changed: %+v", result)
			}
			seen := make(map[string]bool)
			for _, value := range result.Values {
				seen[value] = true
			}
			if len(result.Values) != 2 || !seen["secure-code"] || !seen["secure_code"] {
				t.Fatalf("historical value was lost: %v", result.Values)
			}
		})
	}
}

func TestGeneratedEnumRejectsCanonicalSymbolCollision(t *testing.T) {
	dir := t.TempDir()
	source := writeEnumDef(t, dir, `{"name":"assessment_kind","values":["secure_code_canonical","secure-code","secure_code"],"deprecated":{"secure-code":{"replaced_by":[{"enum":"assessment_kind","value":"secure_code"}]}}}`)
	if err := generateEnum(source, dir, "test-ref"); err == nil {
		t.Fatal("ambiguous canonical symbol was accepted")
	}
	if _, err := os.Stat(filepath.Join(dir, "assessment_kind.go")); !os.IsNotExist(err) {
		t.Fatalf("ambiguous enum produced an output file: %v", err)
	}
}
