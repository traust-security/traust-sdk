// Package ledger is the app's proof chain: a repository's decision ledger, its checkpoints and
// their verification.
package ledger

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's ledgers.
type Handler struct{ core *traust.Client }

// New builds the feature's handler.
func New(core *traust.Client) *Handler { return &Handler{core: core} }

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /repositories/{id}/ledger", web.SignedIn(h.ledger))
	mux.HandleFunc("POST /repositories/{id}/checkpoint", web.SignedIn(h.checkpoint))
	mux.HandleFunc("POST /layers/{id}/verify", web.SignedIn(h.verify))
}

func (h *Handler) ledger(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	layers, err := h.core.ListLayers(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, ledger(id.String(), layers))
}

// checkpoint restamps and signs the repository's layer; core lets only admins.
func (h *Handler) checkpoint(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	if _, err := h.core.Checkpoint(r.Context(), id); err != nil {
		web.Fail(w, r, err)
		return
	}
	h.ledger(w, r)
}

func (h *Handler) verify(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	result, err := h.core.VerifyLayer(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, verified(*result))
}
