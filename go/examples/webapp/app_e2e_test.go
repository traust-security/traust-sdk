//go:build e2e && (darwin || linux)

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// The repository the e2e harness's forge stand-in answers for; make webapp onboards real ones.
const e2eRepository = "https://e2e-forge.test/e2e/fixture"

type e2e struct {
	app    http.Handler
	user   string
	worker *ops.Runner
}

// newE2E is the app on the core make e2e started: its client a machine identity, which core
// refuses to act for a person, and its worker claiming orders from the stand-in forge.
func newE2E(t *testing.T) e2e {
	t.Helper()
	person, err := os.ReadFile(os.Getenv("TRAUST_E2E_TOKEN_FILE"))
	if os.Getenv("TRAUST_E2E_URL") == "" || err != nil {
		t.Fatalf("TRAUST_E2E_URL and TRAUST_E2E_TOKEN_FILE are required; run make e2e (%v)", err)
	}
	machine, err := traust.NewClient(traust.Config{BaseURL: os.Getenv("TRAUST_E2E_URL"), Tokens: traust.FileTokenSource{Path: os.Getenv("TRAUST_E2E_RUNNER_TOKEN_FILE")}})
	if err != nil {
		t.Fatal(err)
	}
	pack, err := ops.LoadPack(packYAML)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := machine.Submit(t.Context(), traust.SubmitRequest{Order: traust.OrderRequest{Asset: "scan"}}); !errors.Is(err, traust.ErrForbidden) {
		t.Fatalf("a machine's own submit: %v, want ErrForbidden", err)
	}
	runner, err := worker(machine, pack, []string{"e2e-forge.test"})
	if err != nil {
		t.Fatal(err)
	}
	user := strings.TrimSpace(string(person))
	h := (&app{machine, pack.Pin()}).routes()
	if w := send(h, http.MethodPost, "/sign-in", "", url.Values{"token": {user}}); w.Code != http.StatusSeeOther {
		t.Fatalf("sign in: %d %s", w.Code, w.Body)
	}
	return e2e{h, user, runner}
}

func (e e2e) get(t *testing.T, path string) string {
	t.Helper()
	w := send(e.app, http.MethodGet, path, e.user, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("GET %s: %d %s", path, w.Code, w.Body)
	}
	return w.Body.String()
}

func TestE2EOnboardScanAndDecide(t *testing.T) {
	e := newE2E(t)
	form := url.Values{"product": {"webapp-e2e"}, "label": {"Webapp e2e"}, "repositories": {e2eRepository}}
	if w := send(e.app, http.MethodPost, "/onboard", e.user, form); w.Header().Get("HX-Redirect") != "/" {
		t.Fatalf("onboard: %d %s", w.Code, w.Body)
	}
	repo := regexp.MustCompile(`href="/repositories/([0-9a-f-]+)">` + regexp.QuoteMeta(e2eRepository)).FindStringSubmatch(e.get(t, "/"))
	if repo == nil {
		t.Fatal("home lists no onboarded repository")
	}

	started := send(e.app, http.MethodPost, "/runs", e.user, url.Values{"asset": {"scan"}, "repository_id": {repo[1]}})
	token := polled(t, started, "/runs/")
	if _, err := e.worker.RunOnce(t.Context()); err != nil {
		t.Fatal(err)
	}
	if done := e.get(t, "/runs/"+token); !strings.Contains(done, string(traust.RunOrderStateSucceeded)) {
		t.Fatalf("scan: %s, want succeeded", done)
	}

	if runs := e.get(t, "/repositories/"+repo[1]+"/runs"); !strings.Contains(runs, "scan: succeeded, requested by ") {
		t.Fatalf("run history: %s, want the person's scan", runs)
	}
	findings := e.get(t, "/repositories/"+repo[1]+"/findings")
	ids := regexp.MustCompile(`href="/findings/([0-9a-f-]+)">Fixture: `).FindAllStringSubmatch(findings, -1)
	if len(ids) != 3 {
		t.Fatalf("findings: %s, want the fixture scan's three", findings)
	}
	decided := send(e.app, http.MethodPost, "/findings/"+ids[0][1]+"/decide", e.user, url.Values{"validity": {"false_positive"}, "rationale": {"a fixture finding"}})
	if decided.Code != http.StatusOK {
		t.Fatalf("decide: %d %s", decided.Code, decided.Body)
	}
	if detail := e.get(t, "/findings/"+ids[0][1]); !strings.Contains(detail, "false_positive") {
		t.Fatalf("finding after the decision: %s, want it false_positive", detail)
	}
	if home := e.get(t, "/"); !strings.Contains(home, "<dt>") {
		t.Fatalf("home: the dashboard counts nothing after a scan")
	}
}

// core's report builder (its reports role, which make e2e starts) builds the person's report;
// the row polls until it is ready, then the file downloads.
func TestE2EBuildAndDownloadAReport(t *testing.T) {
	e := newE2E(t)
	token := polled(t, send(e.app, http.MethodPost, "/reports", e.user, url.Values{"report": {"findings-rollup"}}), "/reports/")
	row := ""
	for range 60 {
		if row = e.get(t, "/reports/"+token); !strings.Contains(row, "hx-trigger") {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if !strings.Contains(row, string(traust.ReportStatusReady)) {
		t.Fatalf("report: %s, want it ready", row)
	}
	if file := send(e.app, http.MethodGet, "/reports/"+token+"/file", e.user, nil); file.Code != http.StatusOK || file.Body.Len() == 0 {
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
