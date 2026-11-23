// Package findings is the app's vuln management: a repository's findings, the decisions a person
// records on them, and the review queue that waits on one.
package findings

import (
	"net/http"
	"strings"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's findings and their decisions.
type Handler struct{ core *traust.Client }

// New builds the feature's handler.
func New(core *traust.Client) *Handler { return &Handler{core: core} }

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /repositories/{id}/findings", web.SignedIn(h.findings))
	mux.HandleFunc("GET /repositories/{id}/assessments", web.SignedIn(h.assessments))
	mux.HandleFunc("GET /assessments", web.SignedIn(h.allAssessments))
	mux.HandleFunc("GET /assessments/{id}", web.SignedIn(h.assessment))
	mux.HandleFunc("GET /artifacts/{id}", web.SignedIn(h.artifact))
	mux.HandleFunc("GET /artifacts/{id}/content", web.SignedIn(h.artifactContent))
	mux.HandleFunc("GET /findings/{id}", web.SignedIn(h.finding))
	mux.HandleFunc("POST /findings/{id}/decide", web.SignedIn(h.decide))
	mux.HandleFunc("POST /decisions/{id}/restate", web.SignedIn(h.restate))
	mux.HandleFunc("GET /repositories/{id}/decisions", web.SignedIn(h.decisions))
	mux.HandleFunc("GET /decisions/{id}", web.SignedIn(h.decision))
	mux.HandleFunc("GET /review", web.SignedIn(h.review))
}

func (h *Handler) findings(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	var found []traust.FindingView
	for f, err := range h.core.ListFindingsAll(r.Context(), id, nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		found = append(found, f)
	}
	web.Render(w, r, http.StatusOK, findingList(found))
}

func (h *Handler) finding(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	f, err := h.core.GetFinding(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	decisions, err := h.core.ListDecisions(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, finding(*f, decisions))
}

func (h *Handler) decide(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	validity := traust.Validity(r.PostFormValue("validity"))
	if _, err := h.core.Decide(r.Context(), id, traust.DecisionBody{
		Validity: &validity, Rationale: strings.TrimSpace(r.PostFormValue("rationale")),
	}); err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, decided())
}

func (h *Handler) restate(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	validity := traust.Validity(r.PostFormValue("validity"))
	if _, err := h.core.Restate(r.Context(), id, traust.RestatementBody{
		Validity: &validity, Reason: traust.RestatementReason(r.PostFormValue("reason")),
		Ticket: strings.TrimSpace(r.PostFormValue("ticket")), Rationale: strings.TrimSpace(r.PostFormValue("rationale")),
	}); err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, restated())
}

// review is the countersign queue across the person's repositories.
func (h *Handler) review(w http.ResponseWriter, r *http.Request) {
	var cases []traust.ReviewCase
	for c, err := range h.core.ReviewQueueAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		cases = append(cases, c)
	}
	web.Render(w, r, http.StatusOK, review(cases))
}
