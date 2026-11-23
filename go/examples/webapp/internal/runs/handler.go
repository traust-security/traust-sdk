// Package runs is the app's work: starting a run, watching it finish, and what ran before.
package runs

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's runs.
type Handler struct {
	core *traust.Client
	pack traust.ConfigDigest
}

// New builds the feature's handler; a run it starts is pinned to the pack.
func New(core *traust.Client, pack traust.ConfigDigest) *Handler {
	return &Handler{core: core, pack: pack}
}

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /runs", web.SignedIn(h.startRun))
	mux.HandleFunc("GET /runs/{token}", web.SignedIn(h.runStatus))
	mux.HandleFunc("GET /repositories/{id}/runs", web.SignedIn(h.runHistory))
}

func (h *Handler) startRun(w http.ResponseWriter, r *http.Request) {
	order := traust.OrderRequest{Asset: r.PostFormValue("asset"), Priority: 1}
	if repo := r.PostFormValue("repository_id"); repo != "" {
		id, err := uuid.Parse(repo)
		if err != nil {
			web.Render(w, r, http.StatusBadRequest, web.Failure("That is not a repository."))
			return
		}
		order.RepositoryId = &id
	}
	poller, err := h.core.BeginSubmit(r.Context(), traust.SubmitRequest{Order: order, Config: &[]traust.ConfigDigest{h.pack}})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, runRow(poller.ResumeToken(), traust.RunOrderStateQueued, false))
}

// runStatus answers a run's row; a finished run also tells the page, so lists that depend on it
// (a repository's findings) refresh themselves.
func (h *Handler) runStatus(w http.ResponseWriter, r *http.Request) {
	poller, err := h.core.ResumeOrder(r.PathValue("token"))
	if err != nil {
		web.Render(w, r, http.StatusBadRequest, web.Failure("That is not a run."))
		return
	}
	order, done, err := poller.Poll(r.Context())
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	if done {
		w.Header().Set("HX-Trigger", "run-finished")
	}
	web.Render(w, r, http.StatusOK, runRow(poller.ResumeToken(), order.State, done))
}

// runHistory is a repository's latest runs, from any requester.
func (h *Handler) runHistory(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	limit := 10
	page, err := h.core.ListOrders(r.Context(), &traust.ListOrdersParams{RepositoryId: &id, Limit: &limit})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, runHistory(page.Items))
}
