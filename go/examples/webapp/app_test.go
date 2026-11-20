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
	return (&app{core.Client(), pack.Pin()}).routes()
}

// send makes one request, signed in when token is set.
func send(h http.Handler, method, path, token string, form url.Values) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if token != "" {
		r.AddCookie(&http.Cookie{Name: sessionCookie, Value: token})
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
	var asked traust.DashboardQuery
	if err := json.Unmarshal(core.Requests()[2].Body, &asked); err != nil || asked.RepositoryIds == nil || (*asked.RepositoryIds)[0] != repo {
		t.Fatalf("dashboard query %+v %v, want it scoped to the person's repositories", asked, err)
	}
	for _, want := range []string{"Signed in as Ada", "https://forge.example/a", `hx-post="/runs"`, "<dt>high</dt><dd>3</dd>"} {
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
	done := send(h, http.MethodGet, "/runs/"+id.String(), "ada", nil).Body.String()
	if !strings.Contains(done, "succeeded") || strings.Contains(done, "hx-trigger") {
		t.Fatalf("finished: %s, want succeeded and polling stopped", done)
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

func TestWorkNeedsASignedInPerson(t *testing.T) {
	core := traustest.New(t)
	h := newApp(t, core)
	for _, route := range [][2]string{
		{http.MethodPost, "/runs"}, {http.MethodGet, "/runs/" + uuid.NewString()},
		{http.MethodPost, "/reports"}, {http.MethodGet, "/reports/" + uuid.NewString()}, {http.MethodGet, "/reports/" + uuid.NewString() + "/file"},
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
