package remediation

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// verifications answers the work item page's lazy verifications section: the verifications core
// has run on the work item's repository, narrowed by core's own work item filter.
func (h *Handler) verifications(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	effort, err := h.coreRemediation.GetRemediation(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	var verifications []traust.Verification
	for verification, err := range h.coreRemediation.ListVerificationsAll(r.Context(), effort.RepositoryId, &traust.ListVerificationsParams{WorkItemId: &id}) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		verifications = append(verifications, verification)
	}
	web.Render(w, r, http.StatusOK, verificationList(verifications))
}

// validationAttempts answers the work item page's lazy validation attempts section: the
// attempts core has recorded on each finding the work item fixes, by the run that made each.
func (h *Handler) validationAttempts(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	effort, err := h.coreRemediation.GetRemediation(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	var attempts []traust.ValidationAttempt
	for _, finding := range effort.WorkItem.FindingIds {
		for attempt, err := range h.coreFindings.ListValidationAttemptsAll(r.Context(), &traust.ListValidationAttemptsParams{FindingId: &finding}) {
			if err != nil {
				web.Fail(w, r, err)
				return
			}
			attempts = append(attempts, attempt)
		}
	}
	web.Render(w, r, http.StatusOK, validationAttemptList(attempts))
}
