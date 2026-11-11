package ops

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"gopkg.in/yaml.v3"
)

// The schema core validates packs with, generated from core at CORE_REF (make api-generate).
//
//go:embed pack.schema.json
var packSchemaJSON []byte

var packSchema = sync.OnceValues(func() (*jsonschema.Schema, error) {
	doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(packSchemaJSON))
	if err != nil {
		return nil, err
	}
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat()
	if err := compiler.AddResource("pack.schema.json", doc); err != nil {
		return nil, err
	}
	return compiler.Compile("pack.schema.json")
})

// Pack is a workload pack (assets.yaml), validated and pinned by the digest of its bytes.
type Pack struct {
	name   string
	digest string
	assets map[string]packAsset
}

type packAsset struct {
	Run []string `json:"run"`
	// Model is the asset's model role, what {role} fills; nil when the asset declares none.
	Model *struct {
		Role string `json:"role"`
	} `json:"model"`
	// Partitions is "dynamic:repo" when the asset's orders are one per repository; empty
	// when the asset runs fleet-wide.
	Partitions string `json:"partitions"`
	// Params declares the typed params the order carries for this asset's command, name to
	// value type (e.g. repo_url: HttpsRepoUrl); command checks what the runner fills against
	// the declaration.
	Params map[string]string `json:"params"`
}

// command is an asset's pinned argv. A whole argument may name a parameter the order supplies.
type command []string

// The parameters a command can name, each a whole argument: the order's id and routing rule,
// the claiming worker, the asset's model role and its repository's URL.
const (
	paramRunID   = "{run_id}"
	paramWorker  = "{worker}"
	paramRole    = "{role}"
	paramRule    = "{rule}"
	paramRepoURL = "{repo_url}"
)

// paramName is a parameter's name in a pack's params declaration (no braces).
func paramName(param string) string { return strings.Trim(param, "{}") }

// httpsRepoURLType is the value type the runner fills {repo_url} with: core's HttpsRepoUrl,
// the order's repository read through core and checked against the runner's forge hosts.
const httpsRepoURLType = "HttpsRepoUrl"

func isParam(arg string) bool {
	switch arg {
	case paramRunID, paramWorker, paramRole, paramRule, paramRepoURL:
		return true
	}
	return false
}

// LoadPack validates a pack against the schema core loads packs with.
func LoadPack(raw []byte) (*Pack, error) {
	var doc any
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("ops: parse pack: %w", err)
	}
	encoded, err := json.Marshal(doc)
	if err != nil {
		return nil, fmt.Errorf("ops: parse pack: %w", err)
	}
	instance, err := jsonschema.UnmarshalJSON(bytes.NewReader(encoded))
	if err != nil {
		return nil, fmt.Errorf("ops: parse pack: %w", err)
	}
	schema, err := packSchema()
	if err != nil {
		return nil, fmt.Errorf("ops: pack schema: %w", err)
	}
	if err := schema.Validate(instance); err != nil {
		return nil, fmt.Errorf("ops: invalid pack: %w", err)
	}
	var pack struct {
		Name   string               `json:"pack"`
		Assets map[string]packAsset `json:"assets"`
	}
	if err := json.Unmarshal(encoded, &pack); err != nil {
		return nil, fmt.Errorf("ops: decode pack: %w", err)
	}
	digest := sha256.Sum256(raw)
	return &Pack{name: pack.Name, digest: hex.EncodeToString(digest[:]), assets: pack.Assets}, nil
}

// ref is the name core records the pack's digest under in an order's manifest.
func (p *Pack) ref() string { return "pack:" + p.name }

// Pin is the pack version an order must run under: pass it in Submit's config so runners with
// this exact pack, and only those, run the order.
func (p *Pack) Pin() traust.ConfigDigest { return traust.ConfigDigest{Name: p.ref(), Digest: p.digest} }

// Assets are the assets the pack declares, sorted.
func (p *Pack) Assets() []string { return slices.Sorted(maps.Keys(p.assets)) }

// role is the asset's model role, what {role} fills; empty when the asset declares none.
func (p *Pack) role(name string) string {
	if asset, ok := p.assets[name]; ok && asset.Model != nil {
		return asset.Model.Role
	}
	return ""
}

// partitioned reports whether the asset's orders are one per repository (partitions:
// dynamic:repo): every such order is repository-bound, whatever its command names.
func (p *Pack) partitioned(name string) bool {
	asset, ok := p.assets[name]
	return ok && asset.Partitions == "dynamic:repo"
}

// command is the asset's argv when a runner can fill every parameter it names: each parameter
// a whole argument, an argument naming {role} backed by the asset's model role, and a declared
// param type agreeing with the value this runner fills.
func (p *Pack) command(name string) (command, error) {
	asset, ok := p.assets[name]
	if !ok || len(asset.Run) == 0 {
		return nil, fmt.Errorf("ops: pack %s has no command for asset %q", p.name, name)
	}
	for _, arg := range asset.Run {
		if !isParam(arg) && strings.ContainsAny(arg, "{}") {
			return nil, fmt.Errorf("ops: asset %q has an unresolvable argument %q", name, arg)
		}
	}
	if slices.Contains(asset.Run, paramRole) && asset.Model == nil {
		return nil, fmt.Errorf("ops: asset %q names %s without a model role", name, paramRole)
	}
	// A declared type must agree with the value this runner fills: {repo_url} is core's
	// HttpsRepoUrl, the order's repository read through core and checked against the runner's
	// forge hosts. The other parameters have no declared-type vocabulary yet; a declaration
	// for a parameter the command does not name is the pack's forward contract, not this
	// runner's to check.
	if declared := asset.Params[paramName(paramRepoURL)]; slices.Contains(asset.Run, paramRepoURL) && declared != "" && declared != httpsRepoURLType {
		return nil, fmt.Errorf("ops: asset %q declares %s as %s; the runner fills %s", name, paramRepoURL, declared, httpsRepoURLType)
	}
	return asset.Run, nil
}

// names reports whether the command names the parameter.
func (c command) names(parameter string) bool { return slices.Contains(c, parameter) }

// fill replaces each whole "{param}" argument with its value; a parameter the order cannot
// supply is an error, never a partial substitution.
func (c command) fill(values map[string]string) ([]string, error) {
	argv := make([]string, len(c))
	for i, arg := range c {
		if !isParam(arg) {
			argv[i] = arg
			continue
		}
		value, ok := values[arg]
		if !ok {
			return nil, &unresolvedParam{reason: traust.RunOrderReasonBadRequest, err: fmt.Errorf("%s: no value for this order", arg)}
		}
		argv[i] = value
	}
	return argv, nil
}
