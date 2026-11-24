// Package remediation is the app's fix-this flow: scheduling a finding's fix and the work
// item core opens for it.
package remediation

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	corefindings "github.com/traust-security/traust-sdk/go/v1/traust/findings"
	coreremediation "github.com/traust-security/traust-sdk/go/v1/traust/remediation"
	coreruns "github.com/traust-security/traust-sdk/go/v1/traust/runs"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's remediation work items.
type Handler struct {
	coreRemediation *coreremediation.Client
	coreRuns        *coreruns.Client
	coreFindings    *corefindings.Client
}

// New builds the feature's handler; the order that fixes a work item is core's scheduling
// door's, queued by the same call that opens the item.
func New(core *traust.Client) *Handler {
	return &Handler{
		coreRemediation: coreremediation.New(core),
		coreRuns:        coreruns.New(core),
		coreFindings:    corefindings.New(core),
	}
}

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /findings/{id}/fix", web.SignedIn(h.fix))
	mux.HandleFunc("GET /work-items", web.SignedIn(h.workItems))
	mux.HandleFunc("GET /work-items/{id}", web.SignedIn(h.workItem))
	mux.HandleFunc("GET /work-items/{id}/verifications", web.SignedIn(h.verifications))
	mux.HandleFunc("GET /work-items/{id}/validation-attempts", web.SignedIn(h.validationAttempts))
	mux.HandleFunc("GET /repositories/{id}/work-items", web.SignedIn(h.repositoryWorkItems))
}

// fix schedules the finding's fix: core's door opens the work item and queues the order that
// runs it, judging the fork the person named against its ForkPolicy.
func (h *Handler) fix(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	repo, err := uuid.Parse(r.PostFormValue("repository_id"))
	if err != nil {
		web.Render(w, r, http.StatusBadRequest, web.Failure("That finding is not on a repository."))
		return
	}
	effort, err := h.coreRemediation.CreateRemediation(r.Context(), traust.RemediationRequest{
		ForkUrl:    strings.TrimSpace(r.PostFormValue("fork_url")),
		ForkBranch: strings.TrimSpace(r.PostFormValue("fork_branch")),
		Findings:   []traust.RemediationFindingRequest{{FindingId: id, RepositoryId: repo}},
	})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, fixed(*effort))
}

// workItems is every scheduled effort, from any requester.
func (h *Handler) workItems(w http.ResponseWriter, r *http.Request) {
	var efforts []traust.ScheduledEffort
	for effort, err := range h.coreRemediation.ListRemediationsAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		efforts = append(efforts, effort)
	}
	web.Render(w, r, http.StatusOK, workItemList(efforts))
}

// workItem is one effort: its status, the findings it fixes, its pull request and the run
// core's door queued for it.
func (h *Handler) workItem(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	effort, err := h.coreRemediation.GetRemediation(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	var order *traust.RunOrder
	if effort.OrderId != nil {
		if order, err = h.coreRuns.GetOrder(r.Context(), *effort.OrderId); err != nil {
			web.Fail(w, r, err)
			return
		}
	}
	web.Render(w, r, http.StatusOK, workItem(*effort, order))
}

// repositoryWorkItems is a repository's work items, from any requester.
func (h *Handler) repositoryWorkItems(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	page, err := h.coreRemediation.ListWorkItems(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, repositoryWorkItems(page.Items))
}
