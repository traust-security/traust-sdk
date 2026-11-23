// Package exposure is the app's advisories: finding one by its identifier, reading it, and the
// packages and repositories it reaches.
package exposure

import (
	"errors"
	"net/http"
	"strings"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	coreexposure "github.com/traust-security/traust-sdk/go/v1/traust/exposure"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's advisories.
type Handler struct{ coreExposure *coreexposure.Client }

// New builds the feature's handler.
func New(core *traust.Client) *Handler {
	return &Handler{coreExposure: coreexposure.New(core)}
}

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /advisories", web.SignedIn(h.search))
	mux.HandleFunc("GET /advisories/{id}", web.SignedIn(h.advisory))
	mux.HandleFunc("GET /findings/{id}/advisories", web.SignedIn(h.findingAdvisories))
}

// search is the Advisories page: the advisory core knows under the identifier asked for, as its
// source states it. A miss says so; anything else is core's to word.
func (h *Handler) search(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		web.Render(w, r, http.StatusOK, searchPage("", nil))
		return
	}
	record, err := h.coreExposure.FindAdvisory(r.Context(), &traust.FindAdvisoryParams{Identifier: query})
	if err != nil {
		if errors.Is(err, traust.ErrNotFound) {
			web.Render(w, r, http.StatusNotFound, web.Failure("No advisory matches that identifier."))
			return
		}
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, searchPage(query, record))
}

// advisory is one advisory's page: the record's words for it, the packages it names, and the
// repositories and images whose SBOMs reach them.
func (h *Handler) advisory(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.coreExposure.GetAdvisory(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	record, err := h.coreExposure.FindAdvisory(r.Context(), &traust.FindAdvisoryParams{Identifier: found.Identifier})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	packages, err := h.coreExposure.PackagesAffected(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	affected, err := h.coreExposure.Affected(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, advisoryPage(*found, record, packages, affected))
}

// findingAdvisories answers the finding page's advisory section: the advisories core has tied to
// the finding.
func (h *Handler) findingAdvisories(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	advisories, err := h.coreExposure.AdvisoriesFor(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, tiedAdvisories(advisories))
}
