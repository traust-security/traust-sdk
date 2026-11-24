package runs

import (
	"context"
	"errors"
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// awaiting is one order core holds for a person, with the cost estimate its manifest carries,
// when it carries one.
type awaiting struct {
	item     traust.OrderQueryItem
	estimate string
}

// approvals is the Approvals page: every order core holds for a person to release — what it
// runs, for which repository, who asked, when, and what core estimates it costs.
func (h *Handler) approvals(w http.ResponseWriter, r *http.Request) {
	states := []traust.RunOrderState{traust.RunOrderStateAwaitingApproval}
	var rows []awaiting
	for item, err := range h.coreRuns.QueryOrdersAll(r.Context(), traust.OrderQueryRequest{States: &states}) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		estimate, err := h.orderEstimate(r.Context(), item.Order)
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		rows = append(rows, awaiting{item: item, estimate: estimate})
	}
	web.Render(w, r, http.StatusOK, approvalsPage(rows))
}

// approve releases one order into the queue; core decides who may, and its answer re-renders
// the row — queued, or its refusal with its status.
func (h *Handler) approve(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	order, err := h.coreRuns.Approve(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, approvalLine(traust.OrderQueryItem{Order: *order}, ""))
}

// check asks core's preflight for one order's asset and repository; an order naming no
// repository cannot fill preflight's question, and the answer says so.
func (h *Handler) check(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	order, err := h.coreRuns.GetOrder(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	if order.RepositoryId == nil {
		web.Render(w, r, http.StatusOK, noPreflight())
		return
	}
	asked := traust.PreflightRequest{Asset: order.Asset, RepositoryId: *order.RepositoryId, Lane: order.Lane}
	result, err := h.coreRuns.Preflight(r.Context(), asked)
	if err != nil {
		title, issues := denial(err)
		web.Render(w, r, web.Status(err), preflightDenied(title, issues))
		return
	}
	web.Render(w, r, http.StatusOK, preflightAnswer(*result))
}

// orderEstimate is the cost estimate core recorded for an order, from its manifest — an
// estimate_usd entry, or one a budget verdict carries; a manifest without one has none.
func (h *Handler) orderEstimate(ctx context.Context, order traust.RunOrder) (string, error) {
	if order.Id == nil {
		return "", nil
	}
	manifest, err := h.coreRuns.GetOrderManifest(ctx, *order.Id)
	if err != nil {
		return "", err
	}
	if worded, ok := (*manifest)["estimate_usd"].(string); ok {
		return worded, nil
	}
	if verdict, ok := (*manifest)["budget"].(map[string]any); ok {
		if worded, ok := verdict["estimate_usd"].(string); ok {
			return worded, nil
		}
	}
	return "", nil
}

// denial is core's refusal of a preflight: its words, and every typed issue it carries.
func denial(err error) (string, []traust.DenyIssue) {
	title := http.StatusText(web.Status(err))
	var issues []traust.DenyIssue
	if p, ok := errors.AsType[*traust.ProblemError](err); ok && p.Problem != nil {
		title = p.Problem.Title + ": " + p.Problem.Detail
		if p.Problem.Issues != nil {
			issues = *p.Problem.Issues
		}
	}
	return title, issues
}
