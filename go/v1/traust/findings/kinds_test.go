package findings

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// kebab spells a camel-cased schema name the way core spells its kind: CveReplayArtifact's
// contract is the kind cve-replay.
func kebab(camel string) string {
	var b strings.Builder
	for i, r := range camel {
		if r >= 'A' && r <= 'Z' && i > 0 {
			b.WriteByte('-')
		}
		b.WriteRune(r)
	}
	return strings.ToLower(b.String())
}

// The generated kind map covers every artifact contract type: a generated type whose name ends
// in Artifact and whose kind the enum names is a contract, and each one must map to its kind —
// a contract core adds reaches every consumer's Decode, never only the bytes.
func TestTheKindMapCoversEveryArtifactContractType(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "../client.gen.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	declared := map[string]bool{}
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gen.Specs {
			if typeSpec, ok := spec.(*ast.TypeSpec); ok {
				declared[typeSpec.Name.Name] = true
			}
		}
	}
	found := map[string]traust.ArtifactKind{}
	for typ, kind := range kindsByType {
		if !declared[strings.TrimPrefix(typ.String(), "traust.")] {
			t.Errorf("the kind map covers %s, which the generated client no longer declares", typ)
		}
		found[typ.String()] = kind
	}
	for name := range declared {
		prefix, ok := strings.CutSuffix(name, "Artifact")
		if !ok {
			continue
		}
		kind := kebab(prefix)
		if !artifactKinds[traust.ArtifactKind(kind)] {
			continue // a nested schema of another contract, no kind of its own
		}
		if got, covered := found["traust."+name]; !covered {
			t.Errorf("%s publishes the kind %s, but the kind map covers no entry for it", name, kind)
		} else if string(got) != kind {
			t.Errorf("the kind map decodes %s as %s, want %s", name, got, kind)
		}
		delete(found, "traust."+name)
	}
	for typ := range found {
		t.Errorf("the kind map covers %s, which names no artifact contract", typ)
	}
}
