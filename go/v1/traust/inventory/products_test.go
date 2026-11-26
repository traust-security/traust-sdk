package inventory_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/inventory"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func products(client *inventory.Client, limit int) ([]traust.Product, error) {
	var found []traust.Product
	for p, err := range client.ListProductsAll(context.Background(), &traust.ListProductsParams{Limit: &limit}) {
		if err != nil {
			return found, err
		}
		found = append(found, p)
	}
	return found, nil
}

func TestListProductsAllFollowsCursorAndReportsProblems(t *testing.T) {
	dir := t.TempDir()
	tokenPath := filepath.Join(dir, "token")
	writeFile(t, tokenPath, "t")

	strp := func(s string) *string { return &s }
	page := func(items []traust.Product, next *string) traust.PageProduct {
		return traust.PageProduct{Items: items, Next: next}
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer t" {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(traust.Problem{Type: "about:blank", Title: "Unauthorized", Status: http.StatusUnauthorized, Detail: "a bearer token is required"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("cursor") == "" {
			_ = json.NewEncoder(w).Encode(page([]traust.Product{{Slug: "a"}, {Slug: "b"}}, strp("c2")))
			return
		}
		_ = json.NewEncoder(w).Encode(page([]traust.Product{{Slug: "c"}}, nil))
	}))
	defer server.Close()

	client, err := traust.NewClient(traust.Config{BaseURL: server.URL, Tokens: traust.FileTokenSource{Path: tokenPath}})
	if err != nil {
		t.Fatal(err)
	}
	found, err := products(inventory.New(client), 2)
	if err != nil {
		t.Fatal(err)
	}
	var slugs []string
	for _, p := range found {
		slugs = append(slugs, p.Slug)
	}
	if len(slugs) != 3 || slugs[0] != "a" || slugs[2] != "c" {
		t.Fatalf("walked %v, want [a b c]", slugs)
	}
}
