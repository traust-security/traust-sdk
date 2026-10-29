package ops

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v6"
	"gopkg.in/yaml.v3"
)

//go:embed pack.schema.json
var packSchema []byte

type packAsset struct {
	Run        []string          `json:"run"`
	Params     map[string]string `json:"params"`
	Partitions string            `json:"partitions"`
	Model      *struct {
		Role string `json:"role"`
	} `json:"model"`
}

// Pack is a validated, pinned workload declaration.
type Pack struct {
	name   string
	digest string
	assets map[string]packAsset
}

// LoadPack validates the installed pack against its pinned Contracts schema.
func LoadPack(raw []byte) (*Pack, error) {
	var yamlDoc any
	if err := yaml.Unmarshal(raw, &yamlDoc); err != nil {
		return nil, fmt.Errorf("ops: parse pack YAML: %w", err)
	}
	encoded, err := json.Marshal(yamlDoc)
	if err != nil {
		return nil, fmt.Errorf("ops: convert pack to JSON: %w", err)
	}
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat()
	schemaDoc, err := jsonschema.UnmarshalJSON(bytes.NewReader(packSchema))
	if err != nil {
		return nil, fmt.Errorf("ops: parse pinned pack schema: %w", err)
	}
	const schemaID = "https://example.com/traust-contracts/config/v1/pack.schema.json"
	if err := compiler.AddResource(schemaID, schemaDoc); err != nil {
		return nil, fmt.Errorf("ops: load pinned pack schema: %w", err)
	}
	schema, err := compiler.Compile(schemaID)
	if err != nil {
		return nil, fmt.Errorf("ops: compile pinned pack schema: %w", err)
	}
	instance, err := jsonschema.UnmarshalJSON(bytes.NewReader(encoded))
	if err != nil {
		return nil, fmt.Errorf("ops: parse pack JSON: %w", err)
	}
	if err := schema.Validate(instance); err != nil {
		return nil, fmt.Errorf("ops: invalid pack: %w", err)
	}
	var doc struct {
		Pack   string               `json:"pack"`
		Assets map[string]packAsset `json:"assets"`
	}
	if err := json.Unmarshal(encoded, &doc); err != nil {
		return nil, fmt.Errorf("ops: decode pack: %w", err)
	}
	digest := sha256.Sum256(raw)
	return &Pack{name: doc.Pack, digest: hex.EncodeToString(digest[:]), assets: doc.Assets}, nil
}
