package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type storageMetadata struct {
	ContractVersion string `json:"contract_version"`
	Revision        int    `json:"revision"`
	BaselineID      string `json:"baseline_id"`
}

func readStorageMetadata(source string) (storageMetadata, error) {
	var metadata storageMetadata
	payload, err := os.ReadFile(filepath.Join(source, "metadata.json"))
	if err != nil {
		return metadata, err
	}
	decoder := json.NewDecoder(bytes.NewReader(payload))
	start, err := decoder.Token()
	if err != nil || start != json.Delim('{') {
		return metadata, fmt.Errorf("storage metadata must be an object")
	}
	seen := map[string]bool{}
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return metadata, err
		}
		key, ok := token.(string)
		if !ok || seen[key] {
			return metadata, fmt.Errorf("duplicate storage metadata key")
		}
		seen[key] = true
		var target any
		switch key {
		case "contract_version":
			target = &metadata.ContractVersion
		case "revision":
			target = &metadata.Revision
		case "baseline_id":
			target = &metadata.BaselineID
		default:
			return metadata, fmt.Errorf("unknown storage metadata key")
		}
		if err := decoder.Decode(target); err != nil {
			return metadata, err
		}
	}
	if _, err := decoder.Token(); err != nil {
		return metadata, err
	}
	if len(seen) != 3 {
		return metadata, fmt.Errorf("missing storage metadata key")
	}
	if err := decoder.Decode(new(any)); err != io.EOF {
		return metadata, fmt.Errorf("invalid trailing storage metadata")
	}
	if metadata.ContractVersion != "v1" || metadata.Revision < 1 ||
		!regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,127}$`).MatchString(metadata.BaselineID) {
		return metadata, fmt.Errorf("invalid storage metadata")
	}
	return metadata, nil
}
