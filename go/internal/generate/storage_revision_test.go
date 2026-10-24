package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStorageRevisionComesFromPinnedSource(t *testing.T) {
	for _, tc := range []struct {
		name, declaration, want string
	}{
		{"merged-main", "REVISION = 6\n", "const contractRevision = 6"},
		{"later-migration", "# Old revision: 6\nREVISION = 7\n", "const contractRevision = 7"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			source, metadata := storageRevisionFixture(t)
			if err := os.WriteFile(metadata, []byte(tc.declaration), 0o600); err != nil {
				t.Fatal(err)
			}
			output := filepath.Join(t.TempDir(), "revision.go")
			if err := generateStorageRevision(source, output, "immutable-test-ref"); err != nil {
				t.Fatal(err)
			}
			data, err := os.ReadFile(output)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(data), tc.want) || !strings.Contains(string(data), "traust-contracts immutable-test-ref") {
				t.Fatalf("storage revision did not follow its pinned source: %s", data)
			}
		})
	}
}

func TestStorageRevisionRejectsMissingOrMalformedMetadata(t *testing.T) {
	for name, declaration := range map[string]string{
		"missing-assignment": "CONTRACT_VERSION = 'v1'\n",
		"zero":               "REVISION = 0\n",
		"boolean":            "REVISION = True\n",
		"string":             "REVISION = '7'\n",
		"multiple":           "REVISION = 6\nREVISION = 7\n",
		"expression":         "REVISION = 6 + 1\n",
	} {
		t.Run(name, func(t *testing.T) {
			source, metadata := storageRevisionFixture(t)
			if err := os.WriteFile(metadata, []byte(declaration), 0o600); err != nil {
				t.Fatal(err)
			}
			if err := generateStorageRevision(source, filepath.Join(t.TempDir(), "revision.go"), "test-ref"); err == nil {
				t.Fatal("malformed source metadata silently selected a storage revision")
			}
		})
	}
	t.Run("missing-file", func(t *testing.T) {
		source, _ := storageRevisionFixture(t)
		if err := generateStorageRevision(source, filepath.Join(t.TempDir(), "revision.go"), "test-ref"); err == nil {
			t.Fatal("missing source metadata silently selected a storage revision")
		}
	})
}

func storageRevisionFixture(t *testing.T) (source, metadata string) {
	t.Helper()
	root := t.TempDir()
	source = filepath.Join(root, "storage", "v1")
	metadata = filepath.Join(root, "src", "traust_contracts", "v1", "storage", "sql.py")
	for _, path := range []string{source, filepath.Dir(metadata)} {
		if err := os.MkdirAll(path, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	return source, metadata
}
