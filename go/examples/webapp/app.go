package main

import (
	"errors"
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
	mux.HandleFunc("GET /onboard", a.signedIn(a.onboardForm))
	mux.HandleFunc("POST /onboard/plan", a.signedIn(a.planOnboarding))
	mux.HandleFunc("POST /onboard", a.signedIn(a.onboard))
	mux.HandleFunc("GET /repositories/{id}", a.signedIn(a.repository))
	mux.HandleFunc("GET /repositories/{id}/findings", a.signedIn(a.findings))
	mux.HandleFunc("GET /repositories/{id}/runs", a.signedIn(a.runHistory))
	mux.HandleFunc("GET /findings/{id}", a.signedIn(a.finding))
	mux.HandleFunc("POST /findings/{id}/decide", a.signedIn(a.decide))
	mux.HandleFunc("POST /decisions/{id}/restate", a.signedIn(a.restate))
	mux.HandleFunc("GET /review", a.signedIn(a.review))
	mux.HandleFunc("POST /runs", a.signedIn(a.startRun))
	mux.HandleFunc("GET /runs/{token}", a.signedIn(a.runStatus))
	mux.HandleFunc("POST /reports", a.signedIn(a.startReport))
	mux.HandleFunc("GET /reports/{token}", a.signedIn(a.reportStatus))
	mux.HandleFunc("GET /reports/{token}/file", a.signedIn(a.reportFile))
	return mux
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

// pathID is the {id} or {token} a route names, or a 400 already answered.
func pathID(w http.ResponseWriter, r *http.Request, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(r.PathValue(name))
	if err != nil {
		render(w, r, http.StatusBadRequest, failure("That link names nothing."))
	}
	return id, err == nil
}

// status passes core's refusal through; anything else is core being unavailable.
func status(err error) int {
	if p, ok := errors.AsType[*traust.ProblemError](err); ok && p.StatusCode < http.StatusInternalServerError {
		return p.StatusCode
	}
	return http.StatusBadGateway
}

// fail answers with core's refusal and its reason, which core words for people; anything else is
// core being unavailable.
func fail(w http.ResponseWriter, r *http.Request, err error) {
	code, message := status(err), http.StatusText(status(err))
	if p, ok := errors.AsType[*traust.ProblemError](err); ok && p.Problem != nil && code < http.StatusInternalServerError {
		message = p.Problem.Title + ": " + p.Problem.Detail
	}
	render(w, r, code, failure(message))
}

func render(w http.ResponseWriter, r *http.Request, code int, view templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	_ = view.Render(r.Context(), w)
}

// items reads a generated optional list; nil is an empty one.
func items[T any](list *[]T) []T {
	if list == nil {
		return nil
	}
	return *list
}

func text(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func severity(s *traust.Severity) string {
	if s == nil {
		return "unrated"
	}
	return string(*s)
}

func validity(d *traust.Disposition) string {
	if d == nil || d.Validity == nil {
		return "undecided"
	}
	return string(*d.Validity)
}

func decisionValidity(d traust.Decision) string {
	if d.Validity == nil {
		return "no validity"
	}
	if v, err := d.Validity.AsValidity(); err == nil {
		return string(v)
	}
	return "restated validity"
}

func caseTitle(c traust.ReviewCase) string {
	if c.View != nil {
		return c.View.Finding.Title
	}
	return c.FindingId.String()
}
