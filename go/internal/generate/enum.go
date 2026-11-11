package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// lowerSuffix names the lowercase spellings in an enum that also carries
// historical uppercase ones (e.g. "S" and "s"), keeping existing names stable.
const lowerSuffix = "Lower"

type enumDef struct {
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Values      []string                         `json:"values"`
	Deprecated  map[string]normalizationMetadata `json:"deprecated"`
}

func generateEnum(srcPath, outDir, contractsRef string) error {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	var def enumDef
	if err := json.Unmarshal(data, &def); err != nil {
		return fmt.Errorf("parsing %s: %w", srcPath, err)
	}

	typeName := toPascalCase(identSafe(def.Name))
	outFile := filepath.Join(outDir, toSnakeCase(def.Name)+".go")

	caseVariants := hasCaseVariants(def.Values)
	constNames := make([]string, len(def.Values))
	seen := make(map[string]int, len(def.Values))
	for i, val := range def.Values {
		name := typeName + toPascalCase(identSafe(val))
		if caseVariants && val == strings.ToLower(val) {
			name += lowerSuffix
		}
		if previous, ok := seen[name]; ok {
			oldValue := def.Values[previous]
			canonicalName := name + "Canonical"
			if def.replaces(oldValue, val) {
				if _, exists := seen[canonicalName]; exists {
					return fmt.Errorf("%s: canonical constant %s already exists", srcPath, canonicalName)
				}
				seen[canonicalName] = i
				constNames[i] = canonicalName
				continue
			}
			if !def.replaces(val, oldValue) {
				return fmt.Errorf("%s: values %q and %q both map to constant %s", srcPath, oldValue, val, name)
			}
			if _, exists := seen[canonicalName]; exists {
				return fmt.Errorf("%s: canonical constant %s already exists", srcPath, canonicalName)
			}
			constNames[previous] = canonicalName
			seen[canonicalName] = previous
		}
		seen[name] = i
		constNames[i] = name
	}

	var b strings.Builder
	fmt.Fprintf(&b, "// Code generated from traust-contracts %s. DO NOT EDIT.\n\n", contractsRef)
	b.WriteString("package enums\n\n")

	if def.Description != "" {
		fmt.Fprintf(&b, "// %s — %s\n", typeName, def.Description)
	}
	fmt.Fprintf(&b, "type %s string\n\n", typeName)

	fmt.Fprint(&b, "const (\n")
	for i, val := range def.Values {
		fmt.Fprintf(&b, "\t%s %s = %q\n", constNames[i], typeName, val)
	}
	fmt.Fprint(&b, ")\n\n")

	fmt.Fprintf(&b, "// %sValues returns all valid %s values.\n", typeName, typeName)
	fmt.Fprintf(&b, "func %sValues() []%s {\n", typeName, typeName)
	fmt.Fprintf(&b, "\treturn []%s{\n", typeName)
	for _, name := range constNames {
		fmt.Fprintf(&b, "\t\t%s,\n", name)
	}
	fmt.Fprint(&b, "\t}\n")
	fmt.Fprint(&b, "}\n")

	return formatAndWrite(outFile, []byte(b.String()))
}

func (def enumDef) replaces(legacy, canonical string) bool {
	if _, deprecated := def.Deprecated[canonical]; deprecated {
		return false
	}
	for _, target := range def.Deprecated[legacy].ReplacedBy {
		if target.Enum == def.Name && target.Value == canonical {
			return true
		}
	}
	return false
}

func hasCaseVariants(values []string) bool {
	spellings := make(map[string]string, len(values))
	for _, val := range values {
		folded := strings.ToLower(val)
		if prev, ok := spellings[folded]; ok && prev != val {
			return true
		}
		spellings[folded] = val
	}
	return false
}

// identSafe turns every rune that cannot appear in a Go identifier into a word
// separator, so a registry value such as "source+runtime" yields SourceRuntime
// instead of invalid Go. Letters, digits and the separators toPascalCase already
// understands pass through unchanged, so existing identifiers do not move.
func identSafe(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || strings.ContainsRune("_- :", r) {
			return r
		}
		return '_'
	}, s)
}

func toPascalCase(s string) string {
	var result strings.Builder
	capitalize := true
	for _, r := range s {
		if r == '_' || r == '-' || r == ' ' || r == ':' {
			capitalize = true
			continue
		}
		if capitalize {
			result.WriteRune(unicode.ToUpper(r))
			capitalize = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func toSnakeCase(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "-", "_"), " ", "_")
}
