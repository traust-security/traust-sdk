package findings

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// decisions answers the repository page's decisions section: who decided what, when, and whether
// it restates an earlier decision.
func (h *Handler) decisions(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	decisions, err := h.coreFindings.ListRepositoryDecisions(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, decisionList(decisions))
}

// decision answers a decision's page: the decision itself, its finding, and its restatement.
func (h *Handler) decision(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	d, err := h.coreFindings.GetDecision(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	view, err := h.coreFindings.GetFinding(r.Context(), d.FindingId)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, decisionDetail(*d, *view))
}

// decisionReason words a recorded decision's reason; a reasonless one says so.
func decisionReason(d traust.Decision) string {
	if d.Reason == nil {
		return "no reason"
	}
	if reason, err := d.Reason.AsDecisionReason(); err == nil {
		return string(reason)
	}
	return "unreadable reason"
}
