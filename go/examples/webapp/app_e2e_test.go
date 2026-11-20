//go:build e2e && (darwin || linux)

package main

import (
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
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
	token := polled(t, send(h, http.MethodPost, "/runs", user, url.Values{"asset": {"smoke"}}), "/runs/")

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

// core's report builder (its reports role, which make e2e starts) builds the person's report;
// the row polls until it is ready, then the file downloads.
func TestE2EAPersonBuildsAndDownloadsAReport(t *testing.T) {
	coreURL := os.Getenv("TRAUST_E2E_URL")
	person, err := os.ReadFile(os.Getenv("TRAUST_E2E_TOKEN_FILE"))
	if coreURL == "" || err != nil {
		t.Fatalf("TRAUST_E2E_URL and TRAUST_E2E_TOKEN_FILE are required; run make e2e (%v)", err)
	}
	machine, err := traust.NewClient(traust.Config{BaseURL: coreURL, Tokens: traust.FileTokenSource{Path: os.Getenv("TRAUST_E2E_RUNNER_TOKEN_FILE")}})
	if err != nil {
		t.Fatal(err)
	}
	h := (&app{core: machine}).routes()
	user := strings.TrimSpace(string(person))

	token := polled(t, send(h, http.MethodPost, "/reports", user, url.Values{"report": {"findings-rollup"}}), "/reports/")
	row := ""
	for range 60 {
		if row = send(h, http.MethodGet, "/reports/"+token, user, nil).Body.String(); !strings.Contains(row, "hx-trigger") {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if !strings.Contains(row, string(traust.ReportStatusReady)) {
		t.Fatalf("report: %s, want it ready", row)
	}
	if file := send(h, http.MethodGet, "/reports/"+token+"/file", user, nil); file.Code != http.StatusOK || file.Body.Len() == 0 {
		t.Fatalf("download: %d with %d bytes, want the report", file.Code, file.Body.Len())
	}
}

// polled is the token in a started row's poll path.
func polled(t *testing.T, started *httptest.ResponseRecorder, path string) string {
	t.Helper()
	_, rest, found := strings.Cut(started.Body.String(), `hx-get="`+path)
	token, _, closed := strings.Cut(rest, `"`)
	if started.Code != http.StatusOK || !found || !closed {
		t.Fatalf("start: %d %s, want a row that polls %s", started.Code, started.Body, path)
	}
	return token
}
