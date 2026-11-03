package ops

import (
	"slices"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestPackCommandRunsOnlyVocabularyArguments(t *testing.T) {
	asset := func(lines ...string) string {
		return "pack: p\nassets:\n  a:\n    " + strings.Join(lines, "\n    ") + "\n"
	}
	for _, tc := range []struct {
		name string
		pack string
		ok   bool
	}{
		{"run_id as a whole argument", asset(`run: [tool, "{run_id}"]`), true},
		{"worker as a whole argument", asset(`run: [tool, "{worker}"]`), true},
		{"rule as a whole argument", asset(`run: [tool, "{rule}"]`), true},
		{"repo_url as a whole argument", asset(`run: [tool, "{repo_url}"]`), true},
		{"role with a model role", asset(`run: [tool, "{role}"]`, "model: {role: audit}"), true},
		{"role without a model role", asset(`run: [tool, "{role}"]`), false},
		{"model role without {role}", asset("run: [tool]", "model: {role: audit}"), true},
		{"no placeholder", asset("run: [tool, --all]"), true},
		{"repository partitions", asset("run: [tool]", "partitions: dynamic:repo"), true},
		{"typed params", asset(`run: [tool, "{repo_url}"]`, "params: {repo_url: HttpsRepoUrl}"), true},
		{"mismatched param type", asset(`run: [tool, "{repo_url}"]`, "params: {repo_url: LocalRepoUrl}"), false},
		{"declared type without the parameter", asset("run: [tool]", "params: {repo_url: HttpsRepoUrl}"), true},
		{"forward-declared param", asset("run: [tool]", "params: {tokens: int}"), true},
		{"other placeholder", asset(`run: [tool, "{repo}"]`), false},
		{"embedded placeholder", asset(`run: [tool, "run-{run_id}"]`), false},
		{"placeholder prefix", asset(`run: [tool, "{x}y"]`), false},
		{"placeholder suffix", asset(`run: [tool, "a{b}"]`), false},
		{"stray brace", asset(`run: [tool, "{"]`), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pack, err := LoadPack([]byte(tc.pack))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := pack.command("a"); (err == nil) != tc.ok {
				t.Fatalf("command err=%v, want ok=%v", err, tc.ok)
			}
		})
	}
	pack, err := LoadPack([]byte(asset("run: [tool]")))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pack.command("missing"); err == nil {
		t.Fatal("an asset the pack doesn't define has a command")
	}
}

func TestCommandFillsTheParametersTheOrderSupplies(t *testing.T) {
	id := uuid.New()
	values := map[string]string{
		paramRunID: id.String(), paramWorker: "runner-1", paramRole: "audit",
		paramRule: "rule-1", paramRepoURL: "https://github.com/example/app",
	}
	got, err := command{"tool", paramRunID, paramWorker, paramRole, paramRule, paramRepoURL, "--flag"}.fill(values)
	if err != nil || !slices.Equal(got, []string{"tool", id.String(), "runner-1", "audit", "rule-1", "https://github.com/example/app", "--flag"}) {
		t.Fatalf("resolved argv=%v err=%v", got, err)
	}
	if _, err := (command{"tool", paramRule}).fill(map[string]string{paramRunID: id.String()}); err == nil {
		t.Fatal("a parameter the order cannot supply was filled")
	}
}

func TestPackPinsOrdersAndListsItsAssets(t *testing.T) {
	raw := []byte("pack: p\nassets:\n  b:\n    run: [tool]\n  a:\n    run: [tool]\n")
	pack, err := LoadPack(raw)
	if err != nil {
		t.Fatal(err)
	}
	pin := pack.Pin()
	if pin.Name != "pack:p" || pin.Digest != pack.digest || len(pin.Digest) != 64 {
		t.Fatalf("pin %+v", pin)
	}
	if got := pack.Assets(); !slices.Equal(got, []string{"a", "b"}) {
		t.Fatalf("assets %v", got)
	}
}
