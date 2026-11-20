package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// The session cookie holds the signed-in person's traust token; a real app's OIDC login sets it.
const sessionCookie = "traust_token"

type app struct {
	core *traust.Client
	pack traust.ConfigDigest
}

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("POST /sign-in", a.signIn)
	mux.HandleFunc("POST /sign-out", a.signOut)
	mux.HandleFunc("POST /runs", a.signedIn(a.startRun))
	mux.HandleFunc("GET /runs/{token}", a.signedIn(a.runStatus))
	mux.HandleFunc("POST /reports", a.signedIn(a.startReport))
	mux.HandleFunc("GET /reports/{token}", a.signedIn(a.reportStatus))
	mux.HandleFunc("GET /reports/{token}/file", a.signedIn(a.reportFile))
	return mux
}

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sessionCookie)
	if err != nil {
		render(w, r, http.StatusOK, signIn(""))
		return
	}
	ctx := traust.OnBehalfOf(r.Context(), cookie.Value)
	me, err := a.core.Whoami(ctx)
	if errors.Is(err, traust.ErrUnauthorized) {
		http.SetCookie(w, &http.Cookie{Name: sessionCookie, MaxAge: -1})
		render(w, r, http.StatusUnauthorized, signIn("Your session ended. Sign in again."))
		return
	}
	if err != nil {
		fail(w, r, err)
		return
	}
	var repos []traust.Repository
	for repo, err := range a.core.ListRepositoriesAll(ctx, nil) {
		if err != nil {
			fail(w, r, err)
			return
		}
		repos = append(repos, repo)
	}
	severities, err := a.findingsBySeverity(ctx, repos)
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, home(*me, repos, severities))
}

// findingsBySeverity is one dashboard query over the person's repositories.
func (a *app) findingsBySeverity(ctx context.Context, repos []traust.Repository) ([]tally, error) {
	ids := make([]uuid.UUID, 0, len(repos))
	for _, repo := range repos {
		if repo.Id != nil {
			ids = append(ids, *repo.Id)
		}
	}
	if len(ids) == 0 {
		return nil, nil
	}
	result, err := a.core.QueryDashboard(ctx, traust.DashboardQuery{
		Subject:       traust.SubjectFindings,
		Measures:      []traust.Measure{traust.MeasureFindings},
		GroupBy:       &[]traust.Dimension{traust.DimensionSeverity},
		RepositoryIds: &ids,
	})
	if err != nil {
		return nil, err
	}
	tallies := make([]tally, 0, len(result.Rows))
	for _, row := range result.Rows {
		label, count := "unrated", "0"
		if s := row.Group(traust.DimensionSeverity); s != nil {
			label = *s
		}
		if n := row.Measure(traust.MeasureFindings); n != nil {
			count = fmt.Sprint(*n)
		}
		tallies = append(tallies, tally{label, count})
	}
	return tallies, nil
}

func (a *app) signIn(w http.ResponseWriter, r *http.Request) {
	token := r.PostFormValue("token")
	if _, err := a.core.Whoami(traust.OnBehalfOf(r.Context(), token)); err != nil {
		render(w, r, status(err), signIn("traust refused that token."))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name: sessionCookie, Value: token, Path: "/",
		HttpOnly: true, Secure: r.TLS != nil, SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) signOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: sessionCookie, Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) startRun(w http.ResponseWriter, r *http.Request) {
	poller, err := a.core.BeginSubmit(r.Context(), traust.SubmitRequest{
		Order:  traust.OrderRequest{Asset: r.PostFormValue("asset"), Priority: 1},
		Config: &[]traust.ConfigDigest{a.pack},
	})
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, runRow(poller.ResumeToken(), traust.RunOrderStateQueued, false))
}

func (a *app) runStatus(w http.ResponseWriter, r *http.Request) {
	poller, err := a.core.ResumeOrder(r.PathValue("token"))
	if err != nil {
		render(w, r, http.StatusBadRequest, failure("That is not a run."))
		return
	}
	order, done, err := poller.Poll(r.Context())
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, runRow(poller.ResumeToken(), order.State, done))
}

func (a *app) startReport(w http.ResponseWriter, r *http.Request) {
	name, format := traust.ViewName(r.PostFormValue("report")), traust.ReportFormatMarkdown
	poller, err := a.core.BeginRequestReport(r.Context(), traust.ReportRequest{Report: name, Format: &format})
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, reportRow(poller.ResumeToken(), traust.Report{Report: name, Status: traust.ReportStatusQueued}, false))
}

func (a *app) reportStatus(w http.ResponseWriter, r *http.Request) {
	poller, err := a.core.ResumeReport(r.PathValue("token"))
	if err != nil {
		render(w, r, http.StatusBadRequest, failure("That is not a report."))
		return
	}
	report, done, err := poller.Poll(r.Context())
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, reportRow(poller.ResumeToken(), report, done))
}

func (a *app) reportFile(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("token"))
	if err != nil {
		http.Error(w, "That is not a report.", http.StatusBadRequest)
		return
	}
	file, err := a.core.GetReportFile(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(status(err)), status(err))
		return
	}
	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="report.md"`)
	_, _ = w.Write(file)
}

// signedIn sends the handler's traust calls as the signed-in person.
func (a *app) signedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(sessionCookie)
		if err != nil {
			render(w, r, http.StatusUnauthorized, failure("Sign in first."))
			return
		}
		next(w, r.WithContext(traust.OnBehalfOf(r.Context(), cookie.Value)))
	}
}

// status passes core's refusal through; anything else is core being unavailable.
func status(err error) int {
	if p, ok := errors.AsType[*traust.ProblemError](err); ok && p.StatusCode < http.StatusInternalServerError {
		return p.StatusCode
	}
	return http.StatusBadGateway
}

func fail(w http.ResponseWriter, r *http.Request, err error) {
	code := status(err)
	render(w, r, code, failure(http.StatusText(code)))
}

func render(w http.ResponseWriter, r *http.Request, code int, view templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	_ = view.Render(r.Context(), w)
}
