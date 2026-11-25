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

	"github.com/traust-security/traust-sdk/go/v1/traust"
	coreruns "github.com/traust-security/traust-sdk/go/v1/traust/runs"
)

// A real public repository the e2e onboards; make webapp onboards any public one.
const e2eRepository = "https://github.com/octocat/Hello-World"

type e2e struct {
	app  http.Handler
	user string
}

// newE2E is the app on the box make e2e started, holding only core's URL: every page calls as
// the signed-in person; the box's claim worker runs what the person submits.
func newE2E(t *testing.T) e2e {
	t.Helper()
	person, err := os.ReadFile(os.Getenv("TRAUST_E2E_TOKEN_FILE"))
	if os.Getenv("TRAUST_E2E_URL") == "" || err != nil {
		t.Fatalf("TRAUST_E2E_URL and TRAUST_E2E_TOKEN_FILE are required; run make e2e (%v)", err)
	}
	app, err := traust.NewClient(traust.Config{BaseURL: os.Getenv("TRAUST_E2E_URL")})
	if err != nil {
		t.Fatal(err)
	}
	// A machine identity may not submit a person's work: core refuses it, whoever holds it.
	machine, err := traust.NewClient(traust.Config{BaseURL: os.Getenv("TRAUST_E2E_URL"), Tokens: traust.FileTokenSource{Path: os.Getenv("TRAUST_E2E_RUNNER_TOKEN_FILE")}})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := coreruns.New(machine).Submit(t.Context(), traust.SubmitRequest{Order: traust.OrderRequest{Asset: "full_audit"}}); !errors.Is(err, traust.ErrForbidden) {
		t.Fatalf("a machine's own submit: %v, want ErrForbidden", err)
	}
	user := strings.TrimSpace(string(person))
	h := routes(app, nil)
	if w := send(h, http.MethodPost, "/sign-in", "", url.Values{"token": {user}}); w.Code != http.StatusSeeOther {
		t.Fatalf("sign in: %d %s", w.Code, w.Body)
	}
	return e2e{h, user}
}

func (e e2e) get(t *testing.T, path string) string {
	t.Helper()
	w := send(e.app, http.MethodGet, path, e.user, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("GET %s: %d %s", path, w.Code, w.Body)
	}
	return w.Body.String()
}

func TestE2EOnboardScanTriageAndReview(t *testing.T) {
	e := newE2E(t)
	form := url.Values{"product": {"webapp-e2e"}, "label": {"Webapp e2e"}, "repositories": {e2eRepository}}
	if w := send(e.app, http.MethodPost, "/onboard", e.user, form); w.Header().Get("HX-Redirect") != "/" {
		t.Fatalf("onboard: %d %s", w.Code, w.Body)
	}
	repo := regexp.MustCompile(`href="/repositories/([0-9a-f-]+)">` + regexp.QuoteMeta(e2eRepository)).FindStringSubmatch(e.get(t, "/"))
	if repo == nil {
		t.Fatal("home lists no onboarded repository")
	}
	// the person's audit is admission-held: the engine's tick holds it for an operator, the
	// admin releases it through the app, and only then does the claim worker run it
	audit := polled(t, send(e.app, http.MethodPost, "/runs", e.user, url.Values{"asset": {"full_audit"}, "repository_id": {repo[1]}}), "/runs/")
	e.held(t, audit)
	if page := e.get(t, "/approvals"); !strings.Contains(page, "full_audit") {
		t.Fatalf("approvals: %s, want the held audit", page)
	}
	admin := tokenFile(t, "TRAUST_E2E_ADMIN_TOKEN_FILE")
	if w := send(e.app, http.MethodPost, "/approvals/"+audit, admin, nil); w.Code != http.StatusOK {
		t.Fatalf("the admin's approval: %d %s", w.Code, w.Body)
	}
	if done := e.finished(t, audit); !strings.Contains(done, string(traust.RunOrderStateSucceeded)) {
		t.Fatalf("full_audit: %s, want succeeded", done)
	}
	// sci_triage carries no approval gate, so admission runs it without a person's release
	triage := polled(t, send(e.app, http.MethodPost, "/runs", e.user, url.Values{"asset": {"sci_triage"}, "repository_id": {repo[1]}}), "/runs/")
	if done := e.finished(t, triage); !strings.Contains(done, string(traust.RunOrderStateSucceeded)) {
		t.Fatalf("sci_triage: %s, want succeeded", done)
	}
	if runs := e.get(t, "/repositories/"+repo[1]+"/runs"); !strings.Contains(runs, "sci_triage: succeeded, requested by ") {
		t.Fatalf("run history: %s, want the person's triage", runs)
	}
	findings := e.get(t, "/repositories/"+repo[1]+"/findings")
	if n := len(regexp.MustCompile(`href="/findings/[0-9a-f-]+">Fixture: `).FindAllString(findings, -1)); n != 3 {
		t.Fatalf("findings: %s, want the fixture scan's three", findings)
	}
	if !strings.Contains(findings, "· confirmed") {
		t.Fatalf("findings: %s, want the triage's true positive confirmed", findings)
	}

	// the triage's other verdicts wait on a person: settle one, then restate it
	cases := regexp.MustCompile(`href="/findings/([0-9a-f-]+)">[^<]*</a>\s*needs_review`).FindAllStringSubmatch(e.get(t, "/review"), -1)
	if len(cases) == 0 {
		t.Fatal("review: nothing waits on the person after a triage")
	}
	finding := "/findings/" + cases[0][1]
	if w := send(e.app, http.MethodPost, finding+"/decide", e.user, url.Values{"validity": {"false_positive"}, "rationale": {"a fixture finding"}}); w.Code != http.StatusOK {
		t.Fatalf("decide: %d %s", w.Code, w.Body)
	}
	latest := regexp.MustCompile(`hx-post="/decisions/([0-9a-f-]+)/restate"`).FindStringSubmatch(e.get(t, finding))
	if latest == nil {
		t.Fatal("finding: no decision to restate after deciding")
	}
	restate := url.Values{"validity": {"confirmed"}, "reason": {"operator_error"}, "ticket": {"SEC-1"}, "rationale": {"decided on the wrong finding"}}
	if w := send(e.app, http.MethodPost, "/decisions/"+latest[1]+"/restate", e.user, restate); w.Code != http.StatusForbidden {
		t.Fatalf("a person's restatement: %d, want 403 (restating is an admin's correction)", w.Code)
	}
	if w := send(e.app, http.MethodPost, "/decisions/"+latest[1]+"/restate", admin, restate); w.Code != http.StatusOK {
		t.Fatalf("the admin's restatement: %d %s", w.Code, w.Body)
	}
	if detail := e.get(t, finding); !strings.Contains(detail, "restates an earlier decision") {
		t.Fatalf("finding after the restatement: %s, want it in the history", detail)
	}
	if home := e.get(t, "/"); !strings.Contains(home, "Findings by state") || !strings.Contains(home, "<dt>") {
		t.Fatal("home: the dashboard counts nothing after a scan and a triage")
	}

	// the decisions are the repository's ledger: anyone verifies it, only an admin checkpoints it
	ledger := e.get(t, "/repositories/"+repo[1]+"/ledger")
	layer := regexp.MustCompile(`hx-post="/layers/([0-9a-f-]+)/verify"`).FindStringSubmatch(ledger)
	if layer == nil {
		t.Fatalf("ledger: %s, want the repository's layer", ledger)
	}
	if v := send(e.app, http.MethodPost, "/layers/"+layer[1]+"/verify", e.user, nil).Body.String(); !strings.Contains(v, "verified") {
		t.Fatalf("verify: %s, want the layer verified", v)
	}
	if w := send(e.app, http.MethodPost, "/repositories/"+repo[1]+"/checkpoint", e.user, nil); w.Code != http.StatusForbidden {
		t.Fatalf("a person's checkpoint: %d, want 403", w.Code)
	}
	if w := send(e.app, http.MethodPost, "/repositories/"+repo[1]+"/checkpoint", strings.TrimSpace(string(admin)), nil); w.Code != http.StatusOK {
		t.Fatalf("the admin's checkpoint: %d %s", w.Code, w.Body)
	}
}

// A person's submit of an asset the pack does not carry never runs: the engine's admission
// withholds it, and the run's page names the reason core recorded.
func TestE2EUnknownAssetIsWithheld(t *testing.T) {
	e := newE2E(t)
	token := polled(t, send(e.app, http.MethodPost, "/runs", e.user, url.Values{"asset": {"no_such_asset"}}), "/runs/")
	if done := e.finished(t, token); !strings.Contains(done, string(traust.RunOrderStateWithheld)) {
		t.Fatalf("no_such_asset: %s, want withheld", done)
	}
	if detail := e.get(t, "/runs/"+token+"/detail"); !strings.Contains(detail, string(traust.RunOrderReasonAssetNotInPack)) {
		t.Fatalf("no_such_asset: %s, want the asset-not-in-pack reason", detail)
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

// finished polls a run's row until the box's claim worker has finished the order.
func (e e2e) finished(t *testing.T, token string) string {
	t.Helper()
	row := ""
	for range 240 {
		if row = e.get(t, "/runs/"+token); !strings.Contains(row, "hx-trigger") {
			return row
		}
		time.Sleep(time.Second)
	}
	t.Fatalf("run %s: %s, still running after 4m", token, row)
	return row
}

// held polls a run's row until the engine's admission holds the order for a person; the row
// keeps polling while it waits, so only its words say it is held.
func (e e2e) held(t *testing.T, token string) string {
	t.Helper()
	row := ""
	for range 120 {
		if row = e.get(t, "/runs/"+token); strings.Contains(row, string(traust.RunOrderStateAwaitingApproval)) {
			return row
		}
		time.Sleep(time.Second)
	}
	t.Fatalf("run %s: %s, still not held for approval after 2m", token, row)
	return row
}

// tokenFile reads one of make e2e's minted tokens, whole and trimmed.
func tokenFile(t *testing.T, name string) string {
	t.Helper()
	raw, err := os.ReadFile(os.Getenv(name))
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(raw))
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
