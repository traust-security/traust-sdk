// Command calls generates v1/traust's calls.gen.go files from the oapi-codegen client: one `(T, error)`
// call per operation and one `…All` iterator per paged listing, so consumers never decode
// responses by hand. Operations land in the package their OpenAPI tag maps to (tagPackage);
// health and identity stay in the root beside the transport. Each domain package gets its own
// Client over the root's, with the models referenced as traust.X. It first renames the raw
// client's constructor to NewRawClient, leaving NewClient for the package's own Client.
// Usage: go run ./internal/generate/calls <client.gen.go> <openapi.json> <v1/traust dir>
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var route = regexp.MustCompile(`operationPath := fmt\.Sprintf\("([^"]+)"`)

// The root package's import path, the one every domain client wraps.
const rootImport = "github.com/traust-security/traust-sdk/go/v1/traust"

// tagPackage maps every OpenAPI tag to the package its operations generate into; "" is the
// root, which keeps only health and identity. A tag missing here stops generation: every
// route must land somewhere.
var tagPackage = map[string]string{
	"health":      "",
	"identity":    "",
	"portfolio":   "inventory",
	"registry":    "inventory",
	"ops":         "runs",
	"findings":    "findings",
	"assessments": "findings",
	"artifacts":   "findings",
	"validation":  "findings",
	"remediation": "remediation",
	"exposure":    "exposure",
	"ledger":      "ledger",
	"dashboards":  "insight",
	"reports":     "insight",
}

// packageDoc is each domain package's doc line.
var packageDoc = map[string]string{
	"inventory":   "Package inventory carries core's portfolio and registry routes: the products, repositories, refs and owners you have.",
	"runs":        "Package runs carries core's ops routes: submitting orders and following the work that runs them.",
	"findings":    "Package findings carries core's findings, assessments, artifacts and validation routes.",
	"remediation": "Package remediation carries core's remediation routes: efforts, work items and their verification.",
	"exposure":    "Package exposure carries core's exposure routes: advisories, packages and the reach of both.",
	"ledger":      "Package ledger carries core's ledger routes: layers, events and their verification.",
	"insight":     "Package insight carries core's dashboards and reports routes: counted rows and built report files.",
}

type param struct{ Name, Type string }

type call struct {
	Name, Method, Path string
	Params             []param
	Result             string // JSON200 type; empty for a raw-bytes response
	Accepted           bool   // core answers 202 with the same body as 200: whichever arrived
	Item, Cursor       string // set for a paged listing
	PageParam          string
	PageValue          bool
	Empty              bool // core documents 204 No Content: the call returns nil, nil
	packageName        string
}

func (c call) Signature() string {
	parts := []string{"ctx context.Context"}
	for _, p := range c.Params {
		parts = append(parts, p.Name+" "+p.Type)
	}
	return strings.Join(parts, ", ")
}

// PathArgs renders the call's path parameters for an error message: the resource a
// failed call concerns, so "CorrectRepositoryUrl(3f2a...)" names the repository. Empty
// when the operation addresses no path parameter.
func (c call) PathArgs() string {
	var names []string
	for _, p := range c.Params {
		if strings.HasPrefix(p.Type, "*") || strings.Contains(p.Type, "Body") {
			continue
		}
		if !strings.Contains(c.Path, "{"+p.Name+"}") {
			continue
		}
		names = append(names, "%s")
	}
	if len(names) == 0 {
		return ""
	}
	return "(" + strings.Join(names, ", ") + ")"
}

// CallPrefix is the call's name with the path parameters it concerns
// ("CorrectRepositoryUrl(%s)"), so a failed call names the resource it was for
// (Qodo #77).
func (c call) CallPrefix() string {
	return c.Name + c.PathArgs()
}

// WrapFormat is the error wrap's format string: CallPrefix with the %w cause.
func (c call) WrapFormat() string {
	return c.CallPrefix() + ": %w"
}

// WrapArgs renders the wrap's format arguments with `last` (the error) last:
// "repositoryId, err" - the format names the resource before the %w cause.
func (c call) WrapArgs(last string) string {
	names := []string{}
	for _, p := range c.Params {
		if strings.HasPrefix(p.Type, "*") || strings.Contains(p.Type, "Body") {
			continue
		}
		if !strings.Contains(c.Path, "{"+p.Name+"}") {
			continue
		}
		names = append(names, p.Name)
	}
	return strings.Join(append(names, last), ", ")
}

func (c call) Args() string {
	names := []string{"ctx"}
	for _, p := range c.Params {
		names = append(names, p.Name)
	}
	return strings.Join(names, ", ")
}

func (c call) PageArgs() string {
	names := []string{"ctx"}
	for _, p := range c.Params {
		if p.Name == c.PageParam {
			if c.PageValue {
				names = append(names, "page")
			} else {
				names = append(names, "&page")
			}
		} else {
			names = append(names, p.Name)
		}
	}
	return strings.Join(names, ", ")
}

func (c call) List() bool { return strings.HasPrefix(c.Result, "[]") }

func (c call) Returns() string {
	switch {
	case c.Result == "":
		return "[]byte"
	case c.List():
		return c.Result
	}
	return "*" + c.Result
}

func (c call) PageParamType() string {
	for _, p := range c.Params {
		if p.Name == c.PageParam {
			return strings.TrimPrefix(p.Type, "*")
		}
	}
	return ""
}

// output is one package's calls.gen.go: the package, its calls, and how that package
// spells the root's types and helpers ("" in the root itself, "traust." in a domain).
type output struct {
	Package  string
	Doc      string
	Domain   bool
	Calls    []call
	Iterator bool
	qualify  func(string) string // a root helper: Refusal, Decode, Parse…Response, Pages
	types    func(string) string // a type the root declares: RunOrder, *…Params
}

// RootImport keeps the template's import line in sync with rootImport.
func (output) RootImport() string { return rootImport }

// Core is how a call reaches the raw client: the root holds it, a domain reaches it through core.
func (o output) Core() string {
	if o.Domain {
		return "core.Raw()"
	}
	return "raw"
}

const outTemplate = `// Code generated by internal/generate/calls; DO NOT EDIT.

{{if .Domain}}// {{.Doc}}
{{end}}package {{.Package}}

import (
	"cmp"
	"context"
	"fmt"
	"io"
{{- if .Iterator}}
	"iter"
{{- end}}
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
{{- if .Domain}}
	"{{.RootImport}}"
{{- end}}
)

var (
	_ openapi_types.UUID
	_ = http.StatusNoContent
	_ = io.EOF
	_ = cmp.Compare[int]
)
{{if .Domain}}
// Client is {{.Package}}'s calls over one core client.
type Client struct{ core *traust.Client }

// New returns the {{.Package}} calls for a core client.
func New(core *traust.Client) *Client { return &Client{core: core} }
{{end}}{{range .Calls}}
// {{.Name}} is {{.Method}} {{.Path}}{{if .Result}}: the decoded body, or a *ProblemError when core refuses{{else}}: the exact stored bytes, or a *ProblemError{{end}}.{{if .Empty}}
// It returns nil, nil when core answers 204 No Content.{{end}}
func (c *Client) {{.Name}}({{q .Signature}}) ({{q .Returns}}, error) {
	raw, err := c.{{$.Core}}.{{.Name}}({{.Args}})
	if err != nil {
		return nil, fmt.Errorf("{{.WrapFormat}}", {{.WrapArgs "err"}})
	}
	if err := {{h "Refusal"}}(raw); err != nil {
		return nil, fmt.Errorf("{{.WrapFormat}}", {{.WrapArgs "err"}})
	}
{{- if not .Result}}
	defer func() { _ = raw.Body.Close() }()
	// the stored bytes come only with 200: another success (204) is no file, never empty bytes
	if raw.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("{{.CallPrefix}}: traust-core: HTTP %d, want 200 with the file", {{.WrapArgs "raw.StatusCode"}})
	}
	body, err := io.ReadAll(raw.Body)
	if err != nil {
		return nil, fmt.Errorf("{{.WrapFormat}}", {{.WrapArgs "err"}})
	}
	return body, nil
{{- else}}
	response, err := {{h (printf "Parse%sResponse" .Name)}}(raw)
	if err != nil {
		return nil, fmt.Errorf("{{.WrapFormat}}", {{.WrapArgs "err"}})
	}
{{- if .Empty}}
	if response.StatusCode() == http.StatusNoContent {
		return nil, nil
	}
{{- end}}
{{- if .List}}
	found, err := {{h "Decode"}}(response.HTTPResponse, response.Body, response.JSON200)
	if err != nil {
		return nil, fmt.Errorf("{{.WrapFormat}}", {{.WrapArgs "err"}})
	}
	if found == nil {
		return nil, nil
	}
	return *found, nil
{{- else}}
	found, err := {{h "Decode"}}(response.HTTPResponse, response.Body, {{if .Accepted}}cmp.Or(response.JSON200, response.JSON202){{else}}response.JSON200{{end}})
	if err != nil {
		return nil, fmt.Errorf("{{.WrapFormat}}", {{.WrapArgs "err"}})
	}
	return found, nil
{{- end}}
{{- end}}
}
{{if .Item}}
// {{.Name}}All iterates every item of {{.Name}}, following the cursor page by page; it stops at
// the first error.
func (c *Client) {{.Name}}All({{q .Signature}}) iter.Seq2[{{q .Item}}, error] {
	{{if .PageValue}}page := {{.PageParam}}{{else}}var page {{q .PageParamType}}
	if {{.PageParam}} != nil {
		page = *{{.PageParam}}
	}{{end}}
	return {{h "Pages"}}(func(cursor *string) ([]{{q .Item}}, *string, error) {
		page.{{.Cursor}} = cursor
		found, err := c.{{.Name}}({{.PageArgs}})
		if err != nil {
			return nil, nil, err
		}
		return found.Items, found.Next, nil
	}, page.{{.Cursor}})
}
{{end}}{{end}}`

func main() {
	if len(os.Args) != 4 {
		log.Fatal("usage: calls <client.gen.go> <openapi.json> <v1/traust dir>")
	}
	spec, err := readSpec(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	source = bytes.ReplaceAll(source, []byte("NewClient("), []byte("NewRawClient("))
	if err := os.WriteFile(os.Args[1], source, 0o644); err != nil { // #nosec G306 -- generated source
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, os.Args[1], source, 0)
	if err != nil {
		log.Fatal(err)
	}
	structs, requests, names := index(file)
	groups := map[string][]call{}
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || receiver(fn) != "ClientWithResponses" {
			continue
		}
		name, ok := strings.CutSuffix(fn.Name.Name, "WithResponse")
		if !ok || strings.HasSuffix(name, "WithBody") {
			continue
		}
		c, err := build(fset, source, name, fn, structs, requests[name])
		if err != nil {
			log.Fatal(err)
		}
		c.Empty = spec.empty[c.Name]
		if spec.file[c.Name] {
			c.Result = "" // a file served in several formats: its exact bytes, whatever the format
		}
		pkg, err := spec.packageFor(c.Name)
		if err != nil {
			log.Fatal(err)
		}
		c.packageName = pkg
		groups[pkg] = append(groups[pkg], c)
	}
	for pkg := range groups {
		sort.Slice(groups[pkg], func(i, j int) bool { return groups[pkg][i].Name < groups[pkg][j].Name })
	}
	for pkg, calls := range groups {
		if err := write(os.Args[3], pkg, calls, names); err != nil {
			log.Fatal(err)
		}
	}
}

// spec is the OpenAPI the calls group by: which operations document 204 No Content, which
// serve a file (a 200 with a non-JSON media type) and which tag each carries.
type spec struct {
	empty map[string]bool
	file  map[string]bool
	tags  map[string][]string
}

func readSpec(path string) (*spec, error) {
	raw, err := os.ReadFile(path) // #nosec G304 -- the spec path comes from the Makefile
	if err != nil {
		return nil, err
	}
	var doc struct {
		Paths map[string]map[string]struct {
			OperationID string                     `json:"operationId"`
			Tags        []string                   `json:"tags"`
			Responses   map[string]json.RawMessage `json:"responses"`
		} `json:"paths"`
	}
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	s := &spec{empty: map[string]bool{}, file: map[string]bool{}, tags: map[string][]string{}}
	for _, operations := range doc.Paths {
		for _, op := range operations {
			if op.OperationID == "" {
				continue
			}
			if _, ok := op.Responses["204"]; ok {
				s.empty[camel(op.OperationID)] = true
			}
			var ok struct {
				Content map[string]json.RawMessage `json:"content"`
			}
			if raw, found := op.Responses["200"]; found && json.Unmarshal(raw, &ok) == nil {
				for media := range ok.Content {
					if media != "application/json" {
						s.file[camel(op.OperationID)] = true
					}
				}
			}
			s.tags[camel(op.OperationID)] = dedupe(op.Tags)
		}
	}
	return s, nil
}

// packageFor names the package an operation generates into: the one its tags map to. A
// route without a tag, or one whose tags disagree, stops generation.
func (s *spec) packageFor(name string) (string, error) {
	tags, known := s.tags[name]
	if !known || len(tags) == 0 {
		return "", fmt.Errorf("%s: the OpenAPI carries no tag for it", name)
	}
	pkg, mapped := tagPackage[tags[0]]
	if !mapped {
		return "", fmt.Errorf("%s: tag %q maps to no package in tagPackage", name, tags[0])
	}
	for _, tag := range tags[1:] {
		if tagPackage[tag] != pkg {
			return "", fmt.Errorf("%s: tags %v map to different packages", name, tags)
		}
	}
	return pkg, nil
}

func dedupe(tags []string) []string {
	var seen []string
	for _, tag := range tags {
		if !contains(seen, tag) {
			seen = append(seen, tag)
		}
	}
	return seen
}

func contains(list []string, want string) bool {
	for _, item := range list {
		if item == want {
			return true
		}
	}
	return false
}

// write renders one package's calls.gen.go under dir: the root's beside the transport,
// a domain's in its own subdirectory with the root's types and helpers qualified.
func write(dir, pkg string, calls []call, names map[string]bool) error {
	out := output{Package: "traust", Calls: calls, qualify: func(s string) string { return s }}
	if pkg != "" {
		if packageDoc[pkg] == "" {
			return fmt.Errorf("package %q has no doc in packageDoc", pkg)
		}
		out = output{
			Package: pkg, Doc: packageDoc[pkg], Domain: true, Calls: calls,
			qualify: func(s string) string { return "traust." + s },
			types:   func(s string) string { return qualify(s, names) },
		}
	}
	for _, c := range calls {
		if c.Item != "" {
			out.Iterator = true
		}
	}
	funcs := template.FuncMap{"h": out.qualify, "q": out.types}
	if !out.Domain {
		funcs["q"] = out.qualify
	}
	tmpl := template.Must(template.New("calls").Funcs(funcs).Parse(outTemplate))
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, out); err != nil {
		return err
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formatting %s's calls: %w", pkg, err)
	}
	path := filepath.Join(dir, "calls.gen.go")
	if pkg != "" {
		if err := os.MkdirAll(filepath.Join(dir, pkg), 0o755); err != nil {
			return err
		}
		path = filepath.Join(dir, pkg, "calls.gen.go")
	}
	return os.WriteFile(path, formatted, 0o644) // #nosec G306 -- generated source, world-readable like every source file
}

var ident = regexp.MustCompile(`\.?([A-Za-z_][A-Za-z0-9_]*)`)

// qualify spells a type the root declares as traust.X, leaving selectors
// (openapi_types.UUID) and names the root does not declare alone.
func qualify(typ string, names map[string]bool) string {
	return ident.ReplaceAllStringFunc(typ, func(match string) string {
		name := strings.TrimPrefix(match, ".")
		if strings.HasPrefix(match, ".") || !names[name] {
			return match
		}
		return "traust." + name
	})
}

func camel(id string) string {
	var b strings.Builder
	for _, part := range strings.Split(id, "_") {
		if part != "" {
			b.WriteString(strings.ToUpper(part[:1]) + part[1:])
		}
	}
	return b.String()
}

func index(file *ast.File) (map[string]*ast.StructType, map[string]*ast.FuncDecl, map[string]bool) {
	structs, requests, names := map[string]*ast.StructType{}, map[string]*ast.FuncDecl{}, map[string]bool{}
	aliases := map[string]string{}
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				if t, ok := spec.(*ast.TypeSpec); ok {
					names[t.Name.Name] = true
					if s, ok := t.Type.(*ast.StructType); ok {
						structs[t.Name.Name] = s
					}
					if target, ok := t.Type.(*ast.Ident); ok {
						aliases[t.Name.Name] = target.Name
					}
				}
			}
		case *ast.FuncDecl:
			name, isBuilder := strings.CutPrefix(d.Name.Name, "New")
			if d.Recv != nil || !isBuilder {
				continue
			}
			name, withBody := strings.CutSuffix(name, "RequestWithBody")
			if !withBody {
				if name, isBuilder = strings.CutSuffix(name, "Request"); !isBuilder {
					continue
				}
			}
			if _, seen := requests[name]; withBody || !seen {
				requests[name] = d // the WithBody builder holds the route
			}
		}
	}
	for len(aliases) > 0 {
		resolved := false
		for name, target := range aliases {
			if s := structs[target]; s != nil {
				structs[name] = s
				delete(aliases, name)
				resolved = true
			}
		}
		if !resolved {
			break
		}
	}
	return structs, requests, names
}

func build(fset *token.FileSet, source []byte, name string, fn *ast.FuncDecl, structs map[string]*ast.StructType, request *ast.FuncDecl) (call, error) {
	c := call{Name: name}
	if request == nil {
		return c, fmt.Errorf("%s: no New%sRequest", name, name)
	}
	body := string(source[fset.Position(request.Body.Pos()).Offset:fset.Position(request.Body.End()).Offset])
	path := route.FindStringSubmatch(body)
	if path == nil {
		return c, fmt.Errorf("%s: no operation path", name)
	}
	c.Path = path[1]
	for _, verb := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		if strings.Contains(body, strconv.Quote(verb)) {
			c.Method = verb
		}
	}
	for _, field := range fn.Type.Params.List[1:] {
		if _, variadic := field.Type.(*ast.Ellipsis); variadic {
			continue
		}
		typ := render(fset, field.Type)
		for _, n := range field.Names {
			c.Params = append(c.Params, param{Name: n.Name, Type: typ})
		}
	}
	for _, p := range c.Params {
		if strings.HasPrefix(p.Type, "*") || strings.Contains(p.Type, "Body") {
			continue
		}
		c.Path = strings.Replace(c.Path, "%s", "{"+p.Name+"}", 1)
	}
	response := structs[name+"Response"]
	if response == nil {
		return c, fmt.Errorf("%s: no %sResponse", name, name)
	}
	for _, f := range response.Fields.List {
		if len(f.Names) == 1 && f.Names[0].Name == "JSON200" {
			c.Result = strings.TrimPrefix(render(fset, f.Type), "*")
		}
		// a bare JSON string is a file core also serves as text: its stored bytes, whatever the format
		if c.Result == "string" {
			c.Result = ""
		}
		if len(f.Names) == 1 && f.Names[0].Name == "JSON202" && strings.TrimPrefix(render(fset, f.Type), "*") == c.Result {
			c.Accepted = true
		}
	}
	page := structs[c.Result]
	if page == nil {
		return c, nil
	}
	hasNext := false
	for _, f := range page.Fields.List {
		if len(f.Names) == 1 && f.Names[0].Name == "Items" && strings.HasPrefix(render(fset, f.Type), "[]") {
			c.Item = strings.TrimPrefix(render(fset, f.Type), "[]")
		}
		if len(f.Names) == 1 && f.Names[0].Name == "Next" && render(fset, f.Type) == "*string" {
			hasNext = true
		}
	}
	if c.Item == "" || !hasNext {
		c.Item = ""
		return c, nil
	}
	for _, p := range c.Params {
		params := structs[strings.TrimPrefix(p.Type, "*")]
		if params == nil {
			continue
		}
		for _, f := range params.Fields.List {
			if len(f.Names) == 1 && (f.Names[0].Name == "Cursor" || f.Names[0].Name == "After") && render(fset, f.Type) == "*string" {
				c.Cursor, c.PageParam = f.Names[0].Name, p.Name
				c.PageValue = !strings.HasPrefix(p.Type, "*")
			}
		}
	}
	if c.Cursor == "" {
		c.Item = ""
	}
	return c, nil
}

func receiver(fn *ast.FuncDecl) string {
	if star, ok := fn.Recv.List[0].Type.(*ast.StarExpr); ok {
		if id, ok := star.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	return ""
}

func render(fset *token.FileSet, expr ast.Expr) string {
	var buf bytes.Buffer
	_ = printer.Fprint(&buf, fset, expr)
	return buf.String()
}
