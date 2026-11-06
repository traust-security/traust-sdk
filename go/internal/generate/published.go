package main

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// publishedShapes is what the committed types package already exports. Regeneration keeps
// those shapes, so a contracts change never renames or retypes a type consumers build on.
type publishedShapes struct {
	owner   map[string]string            // type name -> schema filename
	fields  map[string]map[string]string // struct name -> Go field -> Go type
	current map[string]string            // struct name -> this run's field block, normalized
	aliases map[string]string            // alias name -> the struct it names
}

var (
	publishedStruct = regexp.MustCompile(`(?ms)^type (\w+) struct \{\n(.*?)^\}`)
	publishedAlias  = regexp.MustCompile(`(?m)^type (\w+) = (\w+)$`)
	publishedField  = regexp.MustCompile(`^\s*(\w+)\s+(\S+)\s+` + "`")
)

func readPublished(outDir string) (*publishedShapes, error) {
	shapes := &publishedShapes{
		owner:   make(map[string]string),
		fields:  make(map[string]map[string]string),
		current: make(map[string]string),
		aliases: make(map[string]string),
	}
	entries, err := os.ReadDir(outDir)
	if os.IsNotExist(err) {
		return shapes, nil
	}
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") {
			continue
		}
		content, err := os.ReadFile(filepath.Join(outDir, name))
		if err != nil {
			return nil, err
		}
		if !bytes.HasPrefix(content, []byte("// Code generated from traust-contracts")) {
			continue
		}
		schema := strings.TrimSuffix(name, ".go") + ".schema.json"
		for _, m := range publishedStruct.FindAllSubmatch(content, -1) {
			structName, body := string(m[1]), string(m[2])
			shapes.owner[structName] = schema
			fields := make(map[string]string)
			for line := range strings.SplitSeq(body, "\n") {
				if f := publishedField.FindStringSubmatch(line); f != nil {
					fields[f[1]] = f[2]
				}
			}
			shapes.fields[structName] = fields
		}
		for _, m := range publishedAlias.FindAllSubmatch(content, -1) {
			shapes.owner[string(m[1])] = schema
			shapes.aliases[string(m[1])] = string(m[2])
		}
	}
	return shapes, nil
}

func normalizedBody(body string) string {
	return strings.Join(strings.Fields(body), " ")
}

// aliasFor is the published struct, from another schema, that a new struct matches field for
// field as generated now: the new name becomes an alias, so either type's values assign to the
// other.
func (p *publishedShapes) aliasFor(name, schema string) string {
	want, ok := p.current[name]
	if !ok {
		return ""
	}
	if target, aliased := p.aliases[name]; aliased {
		if p.current[target] == want {
			return target
		}
		return ""
	}
	if _, published := p.owner[name]; published {
		return ""
	}
	candidates := make([]string, 0, len(p.current))
	for candidate := range p.current {
		candidates = append(candidates, candidate)
	}
	sort.Strings(candidates)
	for _, candidate := range candidates {
		owner, published := p.owner[candidate]
		if published && owner != schema && p.current[candidate] == want {
			return candidate
		}
	}
	return ""
}

// keepsInterface reports a field already published as interface{}: a nullable section the
// generator now types concretely keeps the shape consumers already build on.
func (p *publishedShapes) keepsInterface(structName, field string) bool {
	return p.fields[structName][field] == "interface{}"
}
