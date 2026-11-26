// Package identity is the app's sign-in and sign-out: a person's traust token in a cookie, pasted
// or brought by the configured issuer's SSO.
package identity

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's sign-in and sign-out.
type Handler struct {
	core *traust.Client
	sso  *SSO
}

// New builds the feature's handler; sso is nil when no issuer is configured.
func New(core *traust.Client, sso *SSO) *Handler { return &Handler{core: core, sso: sso} }

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /sign-in", h.signIn)
	mux.HandleFunc("POST /sign-out", h.signOut)
	mux.HandleFunc("GET /health", h.healthz)
	mux.HandleFunc("GET /whoami", web.SignedIn(h.whoami))
	if h.sso != nil {
		mux.HandleFunc("GET /sso", h.ssoStart)
		mux.HandleFunc("GET /sso/callback", h.ssoCallback)
	}
}

// whoami answers the layout's ask: who core says is signed in.
func (h *Handler) whoami(w http.ResponseWriter, r *http.Request) {
	me, err := h.core.Whoami(r.Context())
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, person(*me))
}

// healthz answers the layout's ask: core's health, no sign-in needed. A core that won't answer
// is said to be unreachable, never mistaken for healthy.
func (h *Handler) healthz(w http.ResponseWriter, r *http.Request) {
	if core, err := h.core.Healthz(r.Context()); err == nil {
		web.Render(w, r, http.StatusOK, health(core.Status))
		return
	}
	web.Render(w, r, http.StatusOK, unreachable())
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	token := r.PostFormValue("token")
	if _, err := h.core.Whoami(traust.OnBehalfOf(r.Context(), token)); err != nil {
		web.Render(w, r, web.Status(err), web.SignIn("traust refused that token.", h.ssoLink()))
		return
	}
	startSession(w, r, token)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ssoStart sends the person to the configured issuer, one request behind them in a cookie.
func (h *Handler) ssoStart(w http.ResponseWriter, r *http.Request) {
	if err := h.sso.start(w, r); err != nil {
		web.Render(w, r, http.StatusBadGateway, web.SignIn("The sign-in provider is unreachable.", h.ssoLink()))
	}
}

// ssoCallback trades the issuer's code for the person's token, keeps it only when core accepts
// it, and never on a request the app can't vouch for.
func (h *Handler) ssoCallback(w http.ResponseWriter, r *http.Request) {
	if refused := r.URL.Query().Get("error"); refused != "" {
		spent(w)
		web.Render(w, r, http.StatusBadRequest, web.SignIn("The sign-in provider refused that sign-in.", h.ssoLink()))
		return
	}
	verifier, code, ok := answered(r)
	if !ok {
		spent(w)
		web.Render(w, r, http.StatusBadRequest,
			web.SignIn("That sign-in link answers no request. Start again.", h.ssoLink()))
		return
	}
	token, err := h.sso.exchange(r.Context(), code, verifier)
	spent(w)
	if err != nil {
		web.Render(w, r, http.StatusBadGateway, web.SignIn("The sign-in provider refused that sign-in.", h.ssoLink()))
		return
	}
	if _, err := h.core.Whoami(traust.OnBehalfOf(r.Context(), token)); err != nil {
		web.Render(w, r, web.Status(err), web.SignIn("traust refused that token.", h.ssoLink()))
		return
	}
	startSession(w, r, token)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) signOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: web.SessionCookie, Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ssoLink is the sign-in page's SSO offer; nothing when no issuer is configured.
func (h *Handler) ssoLink() string { return h.sso.Link() }

// startSession keeps the token core accepted as the person's session, as both sign-in paths do.
// Secure follows the request's own TLS: behind a TLS-terminating proxy, set Secure always.
func startSession(w http.ResponseWriter, r *http.Request, token string) {
	http.SetCookie(w, &http.Cookie{
		Name: web.SessionCookie, Value: token, Path: "/",
		HttpOnly: true, Secure: r.TLS != nil, SameSite: http.SameSiteStrictMode,
	})
}
