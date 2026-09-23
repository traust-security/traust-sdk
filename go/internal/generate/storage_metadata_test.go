package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStorageMetadataRejectsInvalidInput(t *testing.T) {
	for _, payload := range []string{
		`[]`, `{}`, `{"contract_version":"v1","revision":0,"baseline_id":"test"}`,
		`{"CONTRACT_VERSION":"v1","revision":1,"baseline_id":"test"}`,
		`{"contract_version":"v1","revision":1,"revision":null,"baseline_id":"test"}`,
		`{"contract_version":"v1","revision":true,"baseline_id":"test"}`,
		`{"contract_version":"v1","revision":1,"baseline_id":""}`,
		`{"contract_version":"v2","revision":1,"baseline_id":"test"}`,
		`{"contract_version":"v1","revision":1,"baseline_id":"test","extra":1}`,
		`{"contract_version":"v1","revision":1,"baseline_id":"test"} {}`,
	} {
		t.Run(payload, func(t *testing.T) {
			directory := t.TempDir()
			if err := os.WriteFile(filepath.Join(directory, "metadata.json"), []byte(payload), 0o600); err != nil {
				t.Fatal(err)
			}
			if _, err := readStorageMetadata(directory); err == nil {
				t.Fatal("invalid metadata accepted")
			}
		})
	}
	if _, err := readStorageMetadata(t.TempDir()); err == nil {
		t.Fatal("missing metadata accepted")
	}
}
