// Package identity is the app's sign-in and sign-out: a person's traust token in a cookie.
package identity

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's sign-in and sign-out.
type Handler struct{ core *traust.Client }

// New builds the feature's handler.
func New(core *traust.Client) *Handler { return &Handler{core: core} }

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /sign-in", h.signIn)
	mux.HandleFunc("POST /sign-out", h.signOut)
	mux.HandleFunc("GET /health", h.healthz)
	mux.HandleFunc("GET /whoami", web.SignedIn(h.whoami))
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
		web.Render(w, r, web.Status(err), web.SignIn("traust refused that token."))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name: web.SessionCookie, Value: token, Path: "/",
		HttpOnly: true, Secure: r.TLS != nil, SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) signOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: web.SessionCookie, Path: "/", MaxAge: -1})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
