//go:build e2e && (darwin || linux)

package main

import (
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// The app's client is a machine identity, which core refuses to submit for; a person signs in and
// starts a run through the app, a runner finishes it, and the app's row stops polling.
// `make e2e` starts the server.
func TestE2EAPersonStartsARunThatTheRunnerFinishes(t *testing.T) {
	coreURL := os.Getenv("TRAUST_E2E_URL")
	person, err := os.ReadFile(os.Getenv("TRAUST_E2E_TOKEN_FILE"))
	if coreURL == "" || err != nil {
		t.Fatalf("TRAUST_E2E_URL and TRAUST_E2E_TOKEN_FILE are required; run make e2e (%v)", err)
	}
	machine, err := traust.NewClient(traust.Config{BaseURL: coreURL, Tokens: traust.FileTokenSource{Path: os.Getenv("TRAUST_E2E_RUNNER_TOKEN_FILE")}})
	if err != nil {
		t.Fatal(err)
	}
	pack, err := ops.LoadPack(packYAML)
	if err != nil {
		t.Fatal(err)
	}
	submit := traust.SubmitRequest{Order: traust.OrderRequest{Asset: "smoke", Priority: 1}, Config: &[]traust.ConfigDigest{pack.Pin()}}
	if _, err := machine.Submit(t.Context(), submit); !errors.Is(err, traust.ErrForbidden) {
		t.Fatalf("a machine's own submit: %v, want ErrForbidden", err)
	}
	h := (&app{machine, pack.Pin()}).routes()
	user := strings.TrimSpace(string(person))

	if w := send(h, http.MethodPost, "/sign-in", "", url.Values{"token": {user}}); w.Code != http.StatusSeeOther {
		t.Fatalf("sign in: %d %s", w.Code, w.Body)
	}
	started := send(h, http.MethodPost, "/runs", user, url.Values{"asset": {"smoke"}})
	token, _, found := strings.Cut(strings.TrimPrefix(started.Body.String(), `<li id="run-`), `"`)
	if started.Code != http.StatusOK || !found {
		t.Fatalf("start: %d %s", started.Code, started.Body)
	}

	runner, err := ops.NewRunner(ops.RunnerConfig{
		Client: machine, Pack: pack, Executor: ops.ProcessExecutor{}, Assets: pack.Assets(),
		Logger: slog.New(slog.NewTextHandler(os.Stderr, nil)), LeaseSeconds: 30, PollInterval: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := runner.RunOnce(t.Context()); err != nil {
		t.Fatal(err)
	}

	done := send(h, http.MethodGet, "/runs/"+token, user, nil).Body.String()
	if !strings.Contains(done, string(traust.RunOrderStateSucceeded)) || strings.Contains(done, "hx-trigger") {
		t.Fatalf("finished: %s, want succeeded and polling stopped", done)
	}
}
