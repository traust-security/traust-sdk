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
