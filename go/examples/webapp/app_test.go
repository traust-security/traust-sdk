package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

func newApp(t *testing.T, core *traustest.Server) http.Handler {
	t.Helper()
	pack, err := ops.LoadPack(packYAML)
	if err != nil {
		t.Fatal(err)
	}
	return routes(core.Client(), pack.Pin())
}

// send makes one request, signed in when token is set.
func send(h http.Handler, method, path, token string, form url.Values) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if token != "" {
		r.AddCookie(&http.Cookie{Name: web.SessionCookie, Value: token})
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

// onlyFor answers requests carrying token, and refuses every other caller, as core would.
func onlyFor(token string, answer http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+token {
			traustest.Problem(http.StatusUnauthorized, "unknown caller").ServeHTTP(w, r)
			return
		}
		answer.ServeHTTP(w, r)
	})
}

// anAssessment is a recorded assessment of kind, as core would answer one.
func anAssessment(t *testing.T, id, ref uuid.UUID, kind traust.AssessmentKind, status string) traust.Assessment {
	t.Helper()
	a := traust.Assessment{
		Id: &id, RepositoryRefId: ref, Status: status, Assessor: "webapp-fixture", AssessorVersion: "fixture",
		OccurredAt: time.Date(2026, 2, 3, 4, 5, 6, 0, time.UTC),
	}
	if err := a.Kind.FromAssessmentKind(kind); err != nil {
		t.Fatal(err)
	}
	return a
}

func TestSignInKeepsOnlyATokenCoreAccepts(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	h := newApp(t, core)

	if w := send(h, http.MethodGet, "/", "", nil); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `action="/sign-in"`) {
		t.Fatalf("signed out: %d, want the sign-in form", w.Code)
	}
	refused := send(h, http.MethodPost, "/sign-in", "", url.Values{"token": {"mallory"}})
	if refused.Code != http.StatusUnauthorized || len(refused.Result().Cookies()) != 0 {
		t.Fatalf("a refused token: %d with cookies %v, want 401 and no session", refused.Code, refused.Result().Cookies())
	}
	accepted := send(h, http.MethodPost, "/sign-in", "", url.Values{"token": {"ada"}})
	cookies := accepted.Result().Cookies()
	if accepted.Code != http.StatusSeeOther || len(cookies) != 1 || cookies[0].Value != "ada" || !cookies[0].HttpOnly {
		t.Fatalf("an accepted token: %d %v, want a redirect and an HttpOnly session", accepted.Code, cookies)
	}
}

func TestHomeShowsThePersonAndTheirRepositories(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	repo := uuid.New()
	core.Handle("GET /v1/repositories", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageRepository{
		Items: []traust.Repository{{Id: &repo, Url: "https://forge.example/a"}},
	})))
	high, three := "high", float32(3)
	core.Handle("POST /v1/dashboards:query", onlyFor("ada", traustest.JSON(http.StatusOK, traust.DashboardResults{
		Subject: traust.SubjectFindings,
		Rows:    []traust.DashboardRow{{Groups: map[string]*string{"severity": &high}, Measures: map[string]*float32{"findings": &three}}},
	})))
	body := send(newApp(t, core), http.MethodGet, "/", "ada", nil).Body.String()
	var subjects []traust.Subject
	for _, req := range core.Requests()[2:] {
		var asked traust.DashboardQuery
		if err := json.Unmarshal(req.Body, &asked); err != nil || asked.RepositoryIds == nil || (*asked.RepositoryIds)[0] != repo {
			t.Fatalf("dashboard query %+v %v, want it scoped to the person's repositories", asked, err)
		}
		subjects = append(subjects, asked.Subject)
	}
	if len(subjects) != 3 || subjects[2] != traust.SubjectReviewItems {
		t.Fatalf("dashboard asked %v, want two findings cards and the review queue", subjects)
	}
	for _, want := range []string{"Signed in as Ada", `href="/repositories/` + repo.String() + `"`, "<dt>high</dt><dd>3</dd>", "Review queue by status"} {
		if !strings.Contains(body, want) {
			t.Fatalf("home lacks %q:\n%s", want, body)
		}
	}
}

func TestEachPageAsksWhoIsSignedInAndWatchesCoreHealth(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	core.Handle("GET /v1/repositories", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageRepository{})))
	h := newApp(t, core)

	for _, path := range []string{"/", "/onboard"} {
		page := send(h, http.MethodGet, path, "ada", nil).Body.String()
		for _, want := range []string{`hx-get="/whoami"`, `hx-get="/health"`} {
			if !strings.Contains(page, want) {
				t.Fatalf("%s lacks %q, want the layout asking after the page shows:\n%s", path, want, page)
			}
		}
	}
}

func TestARunPollsUntilCoreFinishesIt(t *testing.T) {
	core := traustest.New(t)
	id := uuid.New()
	finished := time.Now()
	core.Handle("POST /v1/orders", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RunOrder{Id: &id, State: traust.RunOrderStateQueued})))
	core.Handle("GET /v1/orders/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RunOrder{Id: &id, State: traust.RunOrderStateSucceeded, FinishedAt: &finished})))
	h := newApp(t, core)

	started := send(h, http.MethodPost, "/runs", "ada", url.Values{"asset": {"smoke"}})
	if started.Code != http.StatusOK || !strings.Contains(started.Body.String(), `hx-trigger="every 1s"`) {
		t.Fatalf("start: %d %s, want a row that polls", started.Code, started.Body)
	}
	var submitted traust.SubmitRequest
	if err := json.Unmarshal(core.Requests()[0].Body, &submitted); err != nil || submitted.Order.Asset != "smoke" || submitted.Config == nil {
		t.Fatalf("submitted %+v %v, want smoke pinned to the pack", submitted, err)
	}
	done := send(h, http.MethodGet, "/runs/"+id.String(), "ada", nil)
	if !strings.Contains(done.Body.String(), "succeeded") || strings.Contains(done.Body.String(), "hx-trigger") || done.Header().Get("HX-Trigger") != "run-finished" {
		t.Fatalf("finished: %s, want succeeded, polling stopped and the page told", done.Body)
	}
}

func TestOnboardingPreviewsThenOnboards(t *testing.T) {
	core := traustest.New(t)
	core.Handle("POST /v1/onboarding/plan", onlyFor("ada", traustest.JSON(http.StatusOK, traust.OnboardingPlan{
		New: &[]traust.RepositorySpec{{Url: "https://forge.example/org/a"}},
	})))
	core.Handle("POST /v1/onboarding", onlyFor("ada", traustest.JSON(http.StatusOK, traust.OnboardingResult{ProductId: uuid.New()})))
	h := newApp(t, core)
	form := url.Values{"product": {"payments"}, "label": {"Payments"}, "repositories": {"https://forge.example/org/a\n\n https://forge.example/org/b \n"}}

	if preview := send(h, http.MethodPost, "/onboard/plan", "ada", form).Body.String(); !strings.Contains(preview, "register https://forge.example/org/a") || !strings.Contains(preview, `hx-post="/onboard"`) {
		t.Fatalf("preview: %s, want what onboarding will do and the button to do it", preview)
	}
	onboarded := send(h, http.MethodPost, "/onboard", "ada", form)
	var sent traust.OnboardingRequest
	if err := json.Unmarshal(core.Requests()[1].Body, &sent); err != nil || onboarded.Header().Get("HX-Redirect") != "/" {
		t.Fatalf("onboard: %d %v, want a redirect home", onboarded.Code, err)
	}
	if sent.Product != "payments" || *sent.Label != "Payments" || len(*sent.Repositories) != 2 || (*sent.Repositories)[1].Url != "https://forge.example/org/b" || sent.Baseline != nil {
		t.Fatalf("onboarding request %+v, want the form's product, label and two repositories, and no baseline", sent)
	}
}

func TestARepositoryShowsItsBranchAndItsFindingsAreDecided(t *testing.T) {
	core := traustest.New(t)
	repo, ref, found := uuid.New(), uuid.New(), uuid.New()
	branch, high := "main", traust.SeverityHigh
	core.Handle("GET /v1/repositories/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/a"})))
	core.Handle("GET /v1/repositories/{id}/refs/current-default", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{Id: &ref, RepositoryId: repo, Name: &branch, Sha: "4c7f0b9"})))
	view := traust.FindingView{Finding: traust.Finding{Id: &found, RepositoryId: repo, Title: "SQL built from input"}, EffectiveSeverity: &high}
	core.Handle("GET /v1/repositories/{id}/findings", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageFindingView{Items: []traust.FindingView{view}})))
	core.Handle("GET /v1/findings/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, view)))
	core.Handle("POST /v1/findings/{id}/decisions", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Decision{})))
	decision, actor, why := uuid.New(), "triage", "a fixture verdict"
	var verdict traust.Decision_Validity
	if err := verdict.FromValidity(traust.ValidityTruePositive); err != nil {
		t.Fatal(err)
	}
	core.Handle("GET /v1/findings/{id}/decisions", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Decision{{Id: &decision, FindingId: found, Actor: &actor, Rationale: &why, Validity: &verdict}})))
	core.Handle("POST /v1/decisions/{id}/restatement", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Decision{})))
	core.Handle("GET /v1/orders", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageRunOrder{
		Items: []traust.RunOrder{{Asset: "scan", State: traust.RunOrderStateSucceeded, RequestedBy: "engine"}},
	})))
	h := newApp(t, core)
	if runs := send(h, http.MethodGet, "/repositories/"+repo.String()+"/runs", "ada", nil).Body.String(); !strings.Contains(runs, "scan: succeeded, requested by engine") {
		t.Fatalf("run history: %s, want the engine's scan", runs)
	}
	if q := core.Requests()[0].Query; q.Get("repository_id") != repo.String() {
		t.Fatalf("run history asked for %v, want the repository's runs", q)
	}

	page := send(h, http.MethodGet, "/repositories/"+repo.String(), "ada", nil).Body.String()
	if !strings.Contains(page, "Default branch main at 4c7f0b9") || !strings.Contains(page, `repository_id&#34;:&#34;`+repo.String()) {
		t.Fatalf("repository: %s, want its branch and a scan button for it", page)
	}
	list := send(h, http.MethodGet, "/repositories/"+repo.String()+"/findings", "ada", nil).Body.String()
	if !strings.Contains(list, "SQL built from input") || !strings.Contains(list, `<span class="sev high">high</span> · undecided`) {
		t.Fatalf("findings: %s", list)
	}
	detail := send(h, http.MethodGet, "/findings/"+found.String(), "ada", nil).Body.String()
	if !strings.Contains(detail, "true_positive by triage: a fixture verdict") || !strings.Contains(detail, `hx-post="/decisions/`+decision.String()+`/restate"`) {
		t.Fatalf("finding: %s, want its decisions and a restate form for the latest", detail)
	}
	restated := send(h, http.MethodPost, "/decisions/"+decision.String()+"/restate", "ada", url.Values{"validity": {"confirmed"}, "reason": {"operator_error"}, "ticket": {"SEC-1"}, "rationale": {"wrong verdict"}})
	var restatement traust.RestatementBody
	if err := json.Unmarshal(core.Requests()[len(core.Requests())-1].Body, &restatement); err != nil || restated.Code != http.StatusOK ||
		restatement.Reason != traust.OperatorError || *restatement.Validity != traust.ValidityConfirmed || restatement.Ticket != "SEC-1" || restatement.Rationale != "wrong verdict" {
		t.Fatalf("restate: %d %+v %v", restated.Code, restatement, err)
	}
	decided := send(h, http.MethodPost, "/findings/"+found.String()+"/decide", "ada", url.Values{"validity": {"false_positive"}, "rationale": {"test fixture"}})
	var body traust.DecisionBody
	last := core.Requests()[len(core.Requests())-1]
	if err := json.Unmarshal(last.Body, &body); err != nil || decided.Code != http.StatusOK || *body.Validity != traust.ValidityFalsePositive || body.Rationale != "test fixture" {
		t.Fatalf("decide: %d %+v %v, want the person's validity and rationale", decided.Code, body, err)
	}
}

func TestTheReviewQueueListsWhatWaitsOnThePerson(t *testing.T) {
	core := traustest.New(t)
	found := uuid.New()
	core.Handle("GET /v1/findings/review-queue", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageReviewCase{Items: []traust.ReviewCase{{
		FindingId: found, Kind: traust.AwaitingSignoff, View: &traust.FindingView{Finding: traust.Finding{Id: &found, Title: "SQL built from input"}},
	}}})))
	page := send(newApp(t, core), http.MethodGet, "/review", "ada", nil).Body.String()
	if !strings.Contains(page, `href="/findings/`+found.String()+`">SQL built from input</a>`) || !strings.Contains(page, "awaiting_signoff") {
		t.Fatalf("review: %s", page)
	}
}

func TestTheLedgerVerifiesAndOnlyAnAdminCheckpoints(t *testing.T) {
	core := traustest.New(t)
	repo, layer := uuid.New(), uuid.New()
	line, size, root := "main", 4, "abcdef0123456789"
	core.Handle("GET /v1/repositories/{id}/ledger/layers", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Layer{{Id: &layer, RepositoryId: repo, Line: &line, MerkleSize: &size, MerkleRoot: &root}})))
	passed := true
	core.Handle("GET /v1/ledger/layers/{id}/verification", onlyFor("ada", traustest.JSON(http.StatusOK, traust.VerifyResult{LayerId: layer, Passed: &passed})))
	core.Handle("POST /v1/repositories/{id}/ledger/checkpoint", traustest.Problem(http.StatusForbidden, "a ledger checkpoint is an admin's"))
	h := newApp(t, core)

	page := send(h, http.MethodGet, "/repositories/"+repo.String()+"/ledger", "ada", nil).Body.String()
	if !strings.Contains(page, "main layer: 4 events, root abcdef012345,") || !strings.Contains(page, "unsigned") {
		t.Fatalf("ledger: %s", page)
	}
	if v := send(h, http.MethodPost, "/layers/"+layer.String()+"/verify", "ada", nil).Body.String(); !strings.Contains(v, "verified") {
		t.Fatalf("verify: %s", v)
	}
	if w := send(h, http.MethodPost, "/repositories/"+repo.String()+"/checkpoint", "ada", nil); w.Code != http.StatusForbidden || !strings.Contains(w.Body.String(), "an admin&#39;s") {
		t.Fatalf("a person's checkpoint: %d %s, want core's refusal shown", w.Code, w.Body)
	}
}

// eventsPage answers the layer's events as core would: the first page without a cursor, the
// second with one, and nothing after that.
func eventsPage(t *testing.T, layer uuid.UUID, first traust.PageEvent) http.Handler {
	t.Helper()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("cursor") == "" {
			traustest.JSON(http.StatusOK, first).ServeHTTP(w, r)
			return
		}
		traustest.JSON(http.StatusOK, traust.PageEvent{Items: []traust.Event{{Seq: 2, LayerId: layer, DecisionId: uuid.New(), EventPayload: "a severity restated"}}}).ServeHTTP(w, r)
	})
}

func TestALayerPageShowsItsSealAndItsEventsAPageAtATime(t *testing.T) {
	core := traustest.New(t)
	repo, layer, decided := uuid.New(), uuid.New(), uuid.New()
	line, epoch, size, root := "main", 3, 4, "abcdef0123456789"
	signature, method, format := "f00d", "ed25519", 1
	sealed := time.Unix(0, 0).UTC()
	core.Handle("GET /v1/ledger/layers/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Layer{
		Id: &layer, RepositoryId: repo, Line: &line, MerkleEpoch: &epoch, MerkleSize: &size, MerkleRoot: &root,
		MerkleRootSignature: &signature, MerkleSigningMethod: method, MerkleSignatureFormat: &format, CreatedAt: &sealed,
	})))
	cursor := "next-page"
	core.Handle("GET /v1/ledger/layers/{id}/events", onlyFor("ada", eventsPage(t, layer, traust.PageEvent{
		Items: []traust.Event{{Seq: 1, LayerId: layer, DecisionId: decided, EventPayload: "a decision countersigned", CreatedAt: &sealed}},
		Next:  &cursor,
	})))
	h := newApp(t, core)

	page := send(h, http.MethodGet, "/layers/"+layer.String(), "ada", nil)
	for _, want := range []string{
		"main layer 3 · 4 events · root abcdef012345 · signed by ed25519",
		"#1 a decision countersigned on decision " + decided.String(),
		`hx-get="/layers/` + layer.String() + `/events?cursor=next-page"`,
		`hx-post="/layers/` + layer.String() + `/verify"`,
	} {
		if !strings.Contains(page.Body.String(), want) {
			t.Fatalf("layer page lacks %q:\n%s", want, page.Body)
		}
	}
	more := send(h, http.MethodGet, "/layers/"+layer.String()+"/events?cursor=next-page", "ada", nil)
	if body := more.Body.String(); !strings.Contains(body, "#2 a severity restated") || strings.Contains(body, "More events") {
		t.Fatalf("the next page: %s, want the second page's events and no further page", body)
	}
	var asked string
	for _, req := range core.Requests() {
		if req.Path == "/v1/ledger/layers/"+layer.String()+"/events" && req.Query.Get("cursor") == "next-page" {
			asked = "next-page"
		}
	}
	if asked != "next-page" {
		t.Fatalf("core was asked %v, want the more link's cursor passed through", core.Requests())
	}
}

func TestALayerPageKeepsCoresRefusal(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/ledger/layers/{id}", traustest.Problem(http.StatusNotFound, "no such layer"))
	w := send(newApp(t, core), http.MethodGet, "/layers/"+uuid.New().String(), "ada", nil)
	if w.Code != http.StatusNotFound || !strings.Contains(w.Body.String(), "no such layer") {
		t.Fatalf("a refused layer: %d %s, want core's refusal kept", w.Code, w.Body)
	}
}

func TestAReportIsBuiltThenDownloaded(t *testing.T) {
	core := traustest.New(t)
	id := uuid.New()
	core.Handle("POST /v1/reports", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Report{OrderId: id, Report: traust.FindingsRollup, Status: traust.ReportStatusQueued})))
	core.Handle("GET /v1/reports/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Report{OrderId: id, Report: traust.FindingsRollup, Status: traust.ReportStatusReady})))
	core.Handle("GET /v1/reports/{id}/file", onlyFor("ada", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/markdown")
		_, _ = w.Write([]byte("# Findings rollup\n"))
	})))
	h := newApp(t, core)

	started := send(h, http.MethodPost, "/reports", "ada", url.Values{"report": {"findings-rollup"}})
	if !strings.Contains(started.Body.String(), `hx-get="/reports/`+id.String()+`"`) {
		t.Fatalf("start: %d %s, want a row that polls the build", started.Code, started.Body)
	}
	ready := send(h, http.MethodGet, "/reports/"+id.String(), "ada", nil).Body.String()
	if !strings.Contains(ready, "ready") || !strings.Contains(ready, `href="/reports/`+id.String()+`/file"`) || strings.Contains(ready, "hx-trigger") {
		t.Fatalf("ready: %s, want a download link and polling stopped", ready)
	}
	if file := send(h, http.MethodGet, "/reports/"+id.String()+"/file", "ada", nil); file.Body.String() != "# Findings rollup\n" {
		t.Fatalf("download: %d %q, want core's file", file.Code, file.Body)
	}
}

func TestFindingARepositoryAnswersEachPastedURL(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	core.Handle("GET /v1/repositories", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageRepository{})))
	repo := uuid.New()
	core.Handle("POST /v1/repositories:resolve", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.RepositoryResolution{
		{Url: "https://forge.example/org/a", Repository: &traust.Repository{Id: &repo, Url: "https://forge.example/org/a"}},
		{Url: "https://forge.example/org/missing"},
	})))
	h := newApp(t, core)

	home := send(h, http.MethodGet, "/", "ada", nil).Body.String()
	if !strings.Contains(home, "Find a repository") || !strings.Contains(home, `hx-post="/repositories/resolve"`) {
		t.Fatalf("home lacks the find box:\n%s", home)
	}
	answer := send(h, http.MethodPost, "/repositories/resolve", "ada", url.Values{"urls": {"https://forge.example/org/a\n\n https://forge.example/org/missing \n"}}).Body.String()
	var asked traust.RepositoryResolutionRequest
	if err := json.Unmarshal(core.Requests()[len(core.Requests())-1].Body, &asked); err != nil || len(asked.Urls) != 2 ||
		asked.Urls[0] != "https://forge.example/org/a" || asked.Urls[1] != "https://forge.example/org/missing" {
		t.Fatalf("resolve asked %+v %v, want the two pasted URLs", asked, err)
	}
	if !strings.Contains(answer, `href="/repositories/`+repo.String()) || !strings.Contains(answer, "not in inventory") {
		t.Fatalf("resolve: %s, want the found repository and the honest absence", answer)
	}
}

func TestACorrectedURLShowsTheRepositoryItBecomes(t *testing.T) {
	core := traustest.New(t)
	repo := uuid.New()
	core.Handle("GET /v1/repositories/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/a"})))
	core.Handle("GET /v1/repositories/{id}/refs/current-default", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{RepositoryId: repo, Sha: "4c7f0b9"})))
	core.Handle("PUT /v1/repositories/{id}/url", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/renamed"})))
	h := newApp(t, core)

	page := send(h, http.MethodGet, "/repositories/"+repo.String(), "ada", nil).Body.String()
	if !strings.Contains(page, "Correct URL") || !strings.Contains(page, `hx-post="/repositories/`+repo.String()+`/url"`) ||
		!strings.Contains(page, `hx-post="/repositories/`+repo.String()+`/verify"`) {
		t.Fatalf("repository page lacks the administration actions:\n%s", page)
	}
	answered := send(h, http.MethodPost, "/repositories/"+repo.String()+"/url", "ada", url.Values{"url": {" https://forge.example/org/renamed "}})
	var sent traust.UrlCorrection
	if err := json.Unmarshal(core.Requests()[len(core.Requests())-1].Body, &sent); err != nil || sent.Url != "https://forge.example/org/renamed" {
		t.Fatalf("correction sent %+v %v, want the new URL", sent, err)
	}
	if !strings.Contains(answered.Body.String(), "https://forge.example/org/renamed") {
		t.Fatalf("correction: %s, want the repository's new URL", answered.Body)
	}
}

func TestVerifyingARepositoryRendersCoresAnswer(t *testing.T) {
	core := traustest.New(t)
	repo := uuid.New()
	disposition := traust.RepoUnreachable
	core.Handle("POST /v1/repositories/{id}/verify", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryVerification{
		Facts:      traust.ForgeFacts{Status: traust.ComparisonStatusOk},
		Repository: traust.Repository{Id: &repo, Url: "https://forge.example/org/a", Disposition: &disposition},
	})))
	answer := send(newApp(t, core), http.MethodPost, "/repositories/"+repo.String()+"/verify", "ada", nil).Body.String()
	if !strings.Contains(answer, "ok") || !strings.Contains(answer, "unreachable") {
		t.Fatalf("verify: %s, want the facts' status and the repository's standing", answer)
	}
}

func TestRepositoryAdministrationRefusalKeepsItsStatus(t *testing.T) {
	core := traustest.New(t)
	core.Handle("PUT /v1/repositories/{id}/url", traustest.Problem(http.StatusForbidden, "a correction is an admin's"))
	core.Handle("POST /v1/repositories/{id}/verify", traustest.Problem(http.StatusForbidden, "a verification is an admin's"))
	h := newApp(t, core)
	repo := uuid.NewString()
	for _, call := range [][2]string{
		{http.MethodPost, "/repositories/" + repo + "/url"},
		{http.MethodPost, "/repositories/" + repo + "/verify"},
	} {
		form := url.Values{"url": {"https://forge.example/org/b"}}
		if w := send(h, call[0], call[1], "ada", form); w.Code != http.StatusForbidden || !strings.Contains(w.Body.String(), "an admin&#39;s") {
			t.Fatalf("%s %s: %d %s, want core's refusal shown", call[0], call[1], w.Code, w.Body)
		}
	}
}

func TestWorkNeedsASignedInPerson(t *testing.T) {
	core := traustest.New(t)
	h := newApp(t, core)
	for _, route := range [][2]string{
		{http.MethodPost, "/runs"}, {http.MethodGet, "/runs/" + uuid.NewString()},
		{http.MethodPost, "/reports"}, {http.MethodGet, "/reports/" + uuid.NewString()}, {http.MethodGet, "/reports/" + uuid.NewString() + "/file"},
		{http.MethodPost, "/onboard"}, {http.MethodGet, "/repositories/" + uuid.NewString()}, {http.MethodPost, "/findings/" + uuid.NewString() + "/decide"},
		{http.MethodGet, "/review"}, {http.MethodPost, "/decisions/" + uuid.NewString() + "/restate"},
	} {
		if w := send(h, route[0], route[1], "", nil); w.Code != http.StatusUnauthorized {
			t.Fatalf("%s %s: %d, want 401", route[0], route[1], w.Code)
		}
	}
	if n := len(core.Requests()); n != 0 {
		t.Fatalf("core saw %d requests, want none", n)
	}
}

func TestCoreRefusalsKeepTheirStatus(t *testing.T) {
	for core, want := range map[int]int{
		http.StatusForbidden:           http.StatusForbidden,
		http.StatusNotFound:            http.StatusNotFound,
		http.StatusUnprocessableEntity: http.StatusUnprocessableEntity,
		http.StatusInternalServerError: http.StatusBadGateway,
	} {
		t.Run(http.StatusText(core), func(t *testing.T) {
			fake := traustest.New(t)
			fake.Handle("POST /v1/orders", traustest.Problem(core, "refused"))
			if w := send(newApp(t, fake), http.MethodPost, "/runs", "ada", url.Values{"asset": {"smoke"}}); w.Code != want {
				t.Fatalf("core %d: app %d, want %d", core, w.Code, want)
			}
		})
	}
}

func TestTheStylesheetNeedsNoSignIn(t *testing.T) {
	if w := send(newApp(t, traustest.New(t)), http.MethodGet, "/static/app.css", "", nil); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "--frost") {
		t.Fatalf("stylesheet: %d", w.Code)
	}
}

func TestARepositorysRefsAreListedRegisteredAndReDefaulted(t *testing.T) {
	core := traustest.New(t)
	repo, mainRef, releaseRef := uuid.New(), uuid.New(), uuid.New()
	branchKind, tagKind, main, release := "branch", "tag", "main", "v1.0.0"
	isDefault := true
	core.Handle("GET /v1/repositories/{id}/refs", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.RepositoryRef{
		{Id: &mainRef, RepositoryId: repo, Kind: &branchKind, Name: &main, Sha: "4c7f0b9", IsDefault: &isDefault},
		{Id: &releaseRef, RepositoryId: repo, Kind: &tagKind, Name: &release, Sha: "9a1b2c3"},
	})))
	core.Handle("POST /v1/repositories/{id}/refs", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{
		Id: &releaseRef, RepositoryId: repo, Kind: &tagKind, Name: &release, Sha: "5d6e7f8",
	})))
	core.Handle("POST /v1/repositories/{id}/refs/default", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{
		Id: &releaseRef, RepositoryId: repo, Kind: &branchKind, Name: &main, Sha: "9a1b2c3", IsDefault: &isDefault,
	})))
	h := newApp(t, core)

	list := send(h, http.MethodGet, "/repositories/"+repo.String()+"/refs", "ada", nil).Body.String()
	for _, want := range []string{
		`href="/refs/` + mainRef.String() + `">branch main</a> at 4c7f0b9 · default`,
		`href="/refs/` + releaseRef.String() + `">tag ` + release + `</a> at 9a1b2c3`,
		`hx-post="/repositories/` + repo.String() + `/refs/default"`,
		`name="branch" value="main"`, `name="sha" value="4c7f0b9"`, "Make default",
		`hx-post="/repositories/` + repo.String() + `/refs"`, "Register a ref",
	} {
		if !strings.Contains(list, want) {
			t.Fatalf("refs list lacks %q:\n%s", want, list)
		}
	}

	seen := len(core.Requests())
	registered := send(h, http.MethodPost, "/repositories/"+repo.String()+"/refs", "ada", url.Values{"kind": {"tag"}, "name": {release}, "sha": {"5d6e7f8"}})
	var sent traust.RefRegistration
	if err := json.Unmarshal(core.Requests()[seen].Body, &sent); err != nil || registered.Code != http.StatusOK ||
		*sent.Kind != traust.Tag || *sent.Name != release || sent.Sha != "5d6e7f8" {
		t.Fatalf("register: %d %+v %v, want the form's kind, name and sha", registered.Code, sent, err)
	}
	if !strings.Contains(registered.Body.String(), "tag "+release+"</a> at 9a1b2c3") {
		t.Fatalf("register: %s, want the section re-listed with the ref", registered.Body)
	}

	seen = len(core.Requests())
	made := send(h, http.MethodPost, "/repositories/"+repo.String()+"/refs/default", "ada", url.Values{"branch": {"main"}, "sha": {"4c7f0b9"}})
	var naming traust.DefaultRefNaming
	if err := json.Unmarshal(core.Requests()[seen].Body, &naming); err != nil || made.Code != http.StatusOK ||
		naming.Branch != "main" || naming.Sha != "4c7f0b9" {
		t.Fatalf("make default: %d %+v %v, want the branch and its sha", made.Code, naming, err)
	}
	if !strings.Contains(made.Body.String(), "branch main</a> at 4c7f0b9 · default") {
		t.Fatalf("make default: %s, want the section re-listed with the default marked", made.Body)
	}
}

func TestARefPageRendersTheRefItself(t *testing.T) {
	core := traustest.New(t)
	repo, ref := uuid.New(), uuid.New()
	kind, name, isDefault := "branch", "main", true
	core.Handle("GET /v1/refs/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{
		Id: &ref, RepositoryId: repo, Kind: &kind, Name: &name, Sha: "4c7f0b9", IsDefault: &isDefault,
	})))

	page := send(newApp(t, core), http.MethodGet, "/refs/"+ref.String(), "ada", nil).Body.String()
	for _, want := range []string{
		"<h1>branch main</h1>", "branch at 4c7f0b9", "· the repository's default",
		`href="/repositories/` + repo.String() + `"`,
	} {
		if !strings.Contains(page, want) {
			t.Fatalf("ref page lacks %q:\n%s", want, page)
		}
	}
}

func TestRecentRefsSpanTheRepositoriesTheHomePageLists(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	repo, ref := uuid.New(), uuid.New()
	kind, name, isDefault := "branch", "main", true
	core.Handle("GET /v1/repositories", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageRepository{
		Items: []traust.Repository{{Id: &repo, Url: "https://forge.example/org/a"}},
	})))
	core.Handle("POST /v1/dashboards:query", onlyFor("ada", traustest.JSON(http.StatusOK, traust.DashboardResults{})))
	core.Handle("GET /v1/refs", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.RepositoryRef{
		{Id: &ref, RepositoryId: repo, Kind: &kind, Name: &name, Sha: "4c7f0b9", IsDefault: &isDefault},
	})))
	h := newApp(t, core)

	home := send(h, http.MethodGet, "/", "ada", nil).Body.String()
	if !strings.Contains(home, "Recent refs across your repositories") || !strings.Contains(home, `hx-get="/refs"`) {
		t.Fatalf("home lacks the recent refs fragment:\n%s", home)
	}
	fragment := send(h, http.MethodGet, "/refs", "ada", nil)
	if q := core.Requests()[len(core.Requests())-1].Query; q["repository_ids"][0] != repo.String() {
		t.Fatalf("recent refs asked for %v, want them scoped to the person's repositories", q)
	}
	if !strings.Contains(fragment.Body.String(), `href="/refs/`+ref.String()+`">branch main</a> at 4c7f0b9 · default`) {
		t.Fatalf("recent refs: %s, want the ref listed", fragment.Body)
	}
}

func TestARepositorysDecisionsAreListedWithTheRestatedMarked(t *testing.T) {
	core := traustest.New(t)
	repo, found := uuid.New(), uuid.New()
	decision, correction, earlier := uuid.New(), uuid.New(), uuid.New()
	actor, why, ticket := "triage", "a fixture verdict", "SEC-1"
	var verdict traust.Decision_Validity
	if err := verdict.FromValidity(traust.ValidityTruePositive); err != nil {
		t.Fatal(err)
	}
	var reason traust.Decision_Reason
	if err := reason.FromDecisionReason(traust.DecisionReasonOperatorError); err != nil {
		t.Fatal(err)
	}
	core.Handle("GET /v1/repositories/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/a"})))
	core.Handle("GET /v1/repositories/{id}/refs/current-default", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{RepositoryId: repo, Sha: "4c7f0b9"})))
	core.Handle("GET /v1/repositories/{id}/decisions", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Decision{
		{Id: &decision, FindingId: found, Actor: &actor, Rationale: &why, Validity: &verdict, OccurredAt: time.Now()},
		{Id: &correction, FindingId: found, Actor: &actor, Rationale: &why, Validity: &verdict, RestatesId: &earlier, Reason: &reason, Ticket: &ticket, OccurredAt: time.Now()},
	})))
	h := newApp(t, core)

	page := send(h, http.MethodGet, "/repositories/"+repo.String(), "ada", nil).Body.String()
	if !strings.Contains(page, `hx-get="/repositories/`+repo.String()+`/decisions"`) {
		t.Fatalf("repository page lacks the decisions section:\n%s", page)
	}
	list := send(h, http.MethodGet, "/repositories/"+repo.String()+"/decisions", "ada", nil).Body.String()
	for _, want := range []string{
		`href="/decisions/` + decision.String() + `">true_positive</a> by triage`,
		`href="/decisions/` + correction.String() + `">true_positive</a> by triage`,
		"restates an earlier decision",
	} {
		if !strings.Contains(list, want) {
			t.Fatalf("decisions list lacks %q:\n%s", want, list)
		}
	}
}

func TestADecisionPageRendersAndLinksItsFinding(t *testing.T) {
	core := traustest.New(t)
	found, decision, earlier := uuid.New(), uuid.New(), uuid.New()
	actor, why, ticket := "triage", "a fixture verdict", "SEC-1"
	var verdict traust.Decision_Validity
	if err := verdict.FromValidity(traust.ValidityTruePositive); err != nil {
		t.Fatal(err)
	}
	var reason traust.Decision_Reason
	if err := reason.FromDecisionReason(traust.DecisionReasonOperatorError); err != nil {
		t.Fatal(err)
	}
	core.Handle("GET /v1/decisions/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Decision{
		Id: &decision, FindingId: found, Actor: &actor, Rationale: &why, Validity: &verdict,
		RestatesId: &earlier, Reason: &reason, Ticket: &ticket, OccurredAt: time.Now(),
	})))
	core.Handle("GET /v1/findings/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.FindingView{
		Finding: traust.Finding{Id: &found, RepositoryId: uuid.New(), Title: "SQL built from input"},
	})))

	page := send(newApp(t, core), http.MethodGet, "/decisions/"+decision.String(), "ada", nil).Body.String()
	for _, want := range []string{
		"true_positive by triage", "a fixture verdict", "operator_error",
		`href="/findings/` + found.String() + `">SQL built from input</a>`,
		"It restates an earlier decision", "ticket SEC-1",
		`hx-post="/decisions/` + decision.String() + `/restate"`,
	} {
		if !strings.Contains(page, want) {
			t.Fatalf("decision page lacks %q:\n%s", want, page)
		}
	}
}

func TestDecisionRefusalsKeepTheirStatus(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/repositories/{id}/decisions", traustest.Problem(http.StatusForbidden, "a repository's decisions are an admin's"))
	core.Handle("GET /v1/decisions/{id}", traustest.Problem(http.StatusForbidden, "a decision's page is an admin's"))
	h := newApp(t, core)
	for _, path := range []string{"/repositories/" + uuid.NewString() + "/decisions", "/decisions/" + uuid.NewString()} {
		if w := send(h, http.MethodGet, path, "ada", nil); w.Code != http.StatusForbidden || !strings.Contains(w.Body.String(), "an admin&#39;s") {
			t.Fatalf("%s: %d %s, want core's refusal shown", path, w.Code, w.Body)
		}
	}
}

func TestRefAdministrationRefusalKeepsItsStatus(t *testing.T) {
	core := traustest.New(t)
	core.Handle("POST /v1/repositories/{id}/refs", traustest.Problem(http.StatusForbidden, "registering a ref is an admin's"))
	core.Handle("POST /v1/repositories/{id}/refs/default", traustest.Problem(http.StatusForbidden, "naming a default is an admin's"))
	h := newApp(t, core)
	repo := uuid.NewString()
	for _, call := range [][2]string{
		{http.MethodPost, "/repositories/" + repo + "/refs"},
		{http.MethodPost, "/repositories/" + repo + "/refs/default"},
	} {
		form := url.Values{"kind": {"branch"}, "name": {"main"}, "sha": {"4c7f0b9"}, "branch": {"main"}}
		if w := send(h, call[0], call[1], "ada", form); w.Code != http.StatusForbidden || !strings.Contains(w.Body.String(), "an admin&#39;s") {
			t.Fatalf("%s %s: %d %s, want core's refusal shown", call[0], call[1], w.Code, w.Body)
		}
	}
}

func TestTheProductsPageListsCoresCatalogue(t *testing.T) {
	core := traustest.New(t)
	payments, ledger := uuid.New(), uuid.New()
	core.Handle("GET /v1/products", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageProduct{Items: []traust.Product{
		{Id: &payments, Name: "Payments", Slug: "payments"},
		{Id: &ledger, Name: "Ledger", Slug: "ledger"},
	}})))
	page := send(newApp(t, core), http.MethodGet, "/products", "ada", nil).Body.String()
	for _, want := range []string{
		"<h1>Products</h1>",
		`href="/products/` + payments.String() + `">Payments</a> (payments)`,
		`href="/products/` + ledger.String() + `">Ledger</a> (ledger)`,
	} {
		if !strings.Contains(page, want) {
			t.Fatalf("products page lacks %q:\n%s", want, page)
		}
	}
}

func TestAProductPageShowsItselfAndItsRefs(t *testing.T) {
	core := traustest.New(t)
	product, ref, repo := uuid.New(), uuid.New(), uuid.New()
	kind, name, isDefault := "branch", "main", true
	core.Handle("GET /v1/products/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Product{Id: &product, Name: "Payments", Slug: "payments"})))
	core.Handle("GET /v1/products/{id}/refs", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.RepositoryRef{
		{Id: &ref, RepositoryId: repo, Kind: &kind, Name: &name, Sha: "4c7f0b9", IsDefault: &isDefault},
	})))
	page := send(newApp(t, core), http.MethodGet, "/products/"+product.String(), "ada", nil).Body.String()
	for _, want := range []string{
		"<h1>Payments</h1>", "payments",
		`href="/refs/` + ref.String() + `">branch main</a> at 4c7f0b9 · default`,
	} {
		if !strings.Contains(page, want) {
			t.Fatalf("product page lacks %q:\n%s", want, page)
		}
	}
}

func TestARepositoryListsItsProductsAndItsOwners(t *testing.T) {
	core := traustest.New(t)
	repo, product, team := uuid.New(), uuid.New(), uuid.New()
	core.Handle("GET /v1/repositories/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/a"})))
	core.Handle("GET /v1/repositories/{id}/refs/current-default", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{})))
	core.Handle("GET /v1/repositories/{id}/products", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Product{{Id: &product, Name: "Payments", Slug: "payments"}})))
	core.Handle("GET /v1/repositories/{id}/owners", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Team{{Id: &team, Name: "payments"}})))
	h := newApp(t, core)
	page := send(h, http.MethodGet, "/repositories/"+repo.String(), "ada", nil).Body.String()
	for _, want := range []string{
		`hx-get="/repositories/` + repo.String() + `/products"`,
		`hx-get="/repositories/` + repo.String() + `/owners"`,
	} {
		if !strings.Contains(page, want) {
			t.Fatalf("repository page lacks its lazy %q:\n%s", want, page)
		}
	}
	products := send(h, http.MethodGet, "/repositories/"+repo.String()+"/products", "ada", nil).Body.String()
	if !strings.Contains(products, `href="/products/`+product.String()+`">Payments</a>`) {
		t.Fatalf("repository products: %s, want the product listed", products)
	}
	owners := send(h, http.MethodGet, "/repositories/"+repo.String()+"/owners", "ada", nil).Body.String()
	if !strings.Contains(owners, `href="/teams/`+team.String()+`">payments</a>`) {
		t.Fatalf("repository owners: %s, want the team linking to its page", owners)
	}
}

func TestSettingOwnersSendsTheTeams(t *testing.T) {
	core := traustest.New(t)
	repo, team := uuid.New(), uuid.New()
	core.Handle("PUT /v1/repositories/{id}/owners", onlyFor("ada", traustest.JSON(http.StatusOK, traust.SyncReport{})))
	core.Handle("GET /v1/repositories/{id}/owners", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Team{{Id: &team, Name: "payments"}})))
	form := url.Values{"teams": {"payments\n\n ledger "}}
	set := send(newApp(t, core), http.MethodPost, "/repositories/"+repo.String()+"/owners", "ada", form)
	var owners traust.Owners
	if err := json.Unmarshal(core.Requests()[0].Body, &owners); err != nil || set.Code != http.StatusOK ||
		len(owners.Teams) != 2 || owners.Teams[0] != "payments" || owners.Teams[1] != "ledger" {
		t.Fatalf("set owners: %d %+v %v, want the form's two teams", set.Code, owners, err)
	}
	if !strings.Contains(set.Body.String(), `href="/teams/`+team.String()+`">payments</a>`) {
		t.Fatalf("set owners: %s, want the section re-listed with core's answer", set.Body)
	}
}

func TestATeamPageListsItsRepositories(t *testing.T) {
	core := traustest.New(t)
	team, repo := uuid.New(), uuid.New()
	core.Handle("GET /v1/teams/{id}/repositories", onlyFor("ada", traustest.JSON(http.StatusOK, []uuid.UUID{repo})))
	core.Handle("GET /v1/repositories/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/a"})))
	page := send(newApp(t, core), http.MethodGet, "/teams/"+team.String(), "ada", nil).Body.String()
	if !strings.Contains(page, "The team's repositories") ||
		!strings.Contains(page, `href="/repositories/`+repo.String()+`">https://forge.example/org/a</a>`) {
		t.Fatalf("team page: %s, want its repositories listed", page)
	}
}

func TestOwnerAdministrationRefusalKeepsItsStatus(t *testing.T) {
	core := traustest.New(t)
	core.Handle("PUT /v1/repositories/{id}/owners", traustest.Problem(http.StatusForbidden, "setting owners is an admin's"))
	w := send(newApp(t, core), http.MethodPost, "/repositories/"+uuid.NewString()+"/owners", "ada", url.Values{"teams": {"payments"}})
	if w.Code != http.StatusForbidden || !strings.Contains(w.Body.String(), "an admin&#39;s") {
		t.Fatalf("set owners: %d %s, want core's refusal shown", w.Code, w.Body)
	}
}

func TestARepositorysAssessmentsMarkTheLatestScan(t *testing.T) {
	core := traustest.New(t)
	repo, ref, older, latest := uuid.New(), uuid.New(), uuid.New(), uuid.New()
	core.Handle("GET /v1/repositories/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, traust.Repository{Id: &repo, Url: "https://forge.example/org/a"})))
	core.Handle("GET /v1/repositories/{id}/refs/current-default", onlyFor("ada", traustest.JSON(http.StatusOK, traust.RepositoryRef{Id: &ref, RepositoryId: repo})))
	core.Handle("GET /v1/repositories/{id}/assessments", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Assessment{
		anAssessment(t, older, ref, traust.AssessmentKindSecureCode1, "superseded"),
		anAssessment(t, latest, ref, traust.AssessmentKindSecureCode1, "recorded"),
	})))
	core.Handle("GET /v1/repositories/{id}/assessments/latest", onlyFor("ada", traustest.JSON(http.StatusOK,
		anAssessment(t, latest, ref, traust.AssessmentKindSecureCode1, "recorded"))))
	h := newApp(t, core)
	page := send(h, http.MethodGet, "/repositories/"+repo.String(), "ada", nil).Body.String()
	if !strings.Contains(page, "<h2>Assessments</h2>") || !strings.Contains(page, `hx-get="/repositories/`+repo.String()+`/assessments"`) {
		t.Fatalf("repository page lacks the assessments section:\n%s", page)
	}
	section := send(h, http.MethodGet, "/repositories/"+repo.String()+"/assessments", "ada", nil).Body.String()
	if !strings.Contains(section, `href="/assessments/`+older.String()+`"`) || !strings.Contains(section, `href="/assessments/`+latest.String()+`"`) {
		t.Fatalf("assessments: %s, want each assessment linked", section)
	}
	var marked []string
	for _, row := range strings.Split(section, "<li>") {
		if strings.Contains(row, "the latest scan") {
			marked = append(marked, row)
		}
	}
	if len(marked) != 1 || !strings.Contains(marked[0], latest.String()) {
		t.Fatalf("assessments: %s, want only the latest scan marked", section)
	}
	if q := core.Requests()[len(core.Requests())-1].Query["kind"][0]; q != string(traust.AssessmentKindSecureCode1) {
		t.Fatalf("latest asked for kind %q, want the app's scans", q)
	}
}

func TestAnAssessmentPageShowsItsIntakeAndEvidence(t *testing.T) {
	core := traustest.New(t)
	examined, ref, finding, evidence := uuid.New(), uuid.New(), uuid.New(), uuid.New()
	core.Handle("GET /v1/assessments/{id}", onlyFor("ada", traustest.JSON(http.StatusOK,
		anAssessment(t, examined, ref, traust.AssessmentKindSecureCode1, "recorded"))))
	core.Handle("GET /v1/assessments/{id}/intake", onlyFor("ada", traustest.JSON(http.StatusOK, traust.AssessmentIntakeBindings{
		Findings: map[string]uuid.UUID{"ORG-A-4C7F0B9-001": finding},
	})))
	core.Handle("GET /v1/assessments/{id}/artifacts", onlyFor("ada", traustest.JSON(http.StatusOK, []traust.Artifact{{
		Id: &evidence, AssessmentId: &examined, Kind: "report", Digest: "4c7f0b9e2a1c3d4f",
		Uri: "artifacts/report/4c7f0b9e2a1c3d4f.json",
	}})))
	page := send(newApp(t, core), http.MethodGet, "/assessments/"+examined.String(), "ada", nil).Body.String()
	for _, want := range []string{
		"<h1>secure_code</h1>", "recorded · webapp-fixture fixture",
		`href="/refs/` + ref.String() + `"`,
		"<dt>ORG-A-4C7F0B9-001</dt>", `href="/findings/` + finding.String() + `"`,
		`href="/artifacts/` + evidence.String() + `">report</a> · 4c7f0b9e2a1c`,
	} {
		if !strings.Contains(page, want) {
			t.Fatalf("assessment page lacks %q:\n%s", want, page)
		}
	}
}

func TestAllAssessmentsSpanThePersonsRepositories(t *testing.T) {
	core := traustest.New(t)
	ref, scan, triage := uuid.New(), uuid.New(), uuid.New()
	core.Handle("GET /v1/assessments", onlyFor("ada", traustest.JSON(http.StatusOK, traust.PageAssessment{Items: []traust.Assessment{
		anAssessment(t, scan, ref, traust.AssessmentKindSecureCode1, "recorded"),
		anAssessment(t, triage, ref, traust.AssessmentKindTriage, "recorded"),
	}})))
	page := send(newApp(t, core), http.MethodGet, "/assessments", "ada", nil).Body.String()
	if !strings.Contains(page, `href="/assessments/`+scan.String()+`"`) || !strings.Contains(page, "triage") {
		t.Fatalf("all assessments: %s, want the person's assessments listed", page)
	}
}

func TestAnArtifactsEvidenceDownloadsUnderItsName(t *testing.T) {
	evidence, assessment := uuid.New(), uuid.New()
	tests := []struct {
		name        string
		artifact    traust.Artifact
		disposition string
		mediaType   string
	}{
		{
			name: "a report whose uri names json",
			artifact: traust.Artifact{Id: &evidence, AssessmentId: &assessment, Kind: "report",
				Digest: "4c7f0b9e2a1c3d4f", Uri: "artifacts/report/4c7f0b9e2a1c3d4f.json"},
			disposition: `attachment; filename="report-4c7f0b9e2a1c.json"`,
			mediaType:   "application/json",
		},
		{
			name: "a companion core left untyped",
			artifact: traust.Artifact{Id: &evidence, AssessmentId: &assessment, Kind: "findings_layer",
				Digest: "4c7f0b9e2a1c3d4f", Uri: "artifacts/report/4c7f0b9e2a1c3d4f-findings_layer"},
			disposition: `attachment; filename="findings_layer-4c7f0b9e2a1c"`,
			mediaType:   "application/octet-stream",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core := traustest.New(t)
			core.Handle("GET /v1/artifacts/{id}", onlyFor("ada", traustest.JSON(http.StatusOK, tt.artifact)))
			core.Handle("GET /v1/artifacts/{id}/content", onlyFor("ada", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				_, _ = w.Write([]byte(`{"title": "Fixture audit"}`))
			})))
			h := newApp(t, core)
			page := send(h, http.MethodGet, "/artifacts/"+evidence.String(), "ada", nil).Body.String()
			if !strings.Contains(page, `href="/artifacts/`+evidence.String()+`/content"`) || !strings.Contains(page, tt.artifact.Digest) {
				t.Fatalf("artifact page: %s, want its digest and a download link", page)
			}
			download := send(h, http.MethodGet, "/artifacts/"+evidence.String()+"/content", "ada", nil)
			if download.Code != http.StatusOK || download.Body.String() != `{"title": "Fixture audit"}` {
				t.Fatalf("download: %d %s, want core's bytes", download.Code, download.Body)
			}
			if ct := download.Header().Get("Content-Type"); ct != tt.mediaType {
				t.Fatalf("download content type: %q, want %q", ct, tt.mediaType)
			}
			if cd := download.Header().Get("Content-Disposition"); cd != tt.disposition {
				t.Fatalf("download disposition: %q, want %q", cd, tt.disposition)
			}
		})
	}
}

func TestAssessmentRefusalsKeepTheirStatus(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/repositories/{id}/assessments", traustest.Problem(http.StatusForbidden, "assessments are an admin's"))
	core.Handle("GET /v1/assessments/{id}", traustest.Problem(http.StatusForbidden, "assessments are an admin's"))
	core.Handle("GET /v1/artifacts/{id}", traustest.Problem(http.StatusForbidden, "evidence is an admin's"))
	core.Handle("GET /v1/artifacts/{id}/content", traustest.Problem(http.StatusForbidden, "evidence is an admin's"))
	h := newApp(t, core)
	id := uuid.NewString()
	for _, path := range []string{
		"/repositories/" + id + "/assessments",
		"/assessments/" + id,
		"/artifacts/" + id,
		"/artifacts/" + id + "/content",
	} {
		if w := send(h, http.MethodGet, path, "ada", nil); w.Code != http.StatusForbidden || !strings.Contains(w.Body.String(), "admin&#39;s") {
			t.Fatalf("%s: %d %s, want core's refusal shown", path, w.Code, w.Body)
		}
	}
}
