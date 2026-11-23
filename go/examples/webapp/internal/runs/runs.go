package runs

import (
	"encoding/json"
	"maps"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// orderStates are the states a run may be narrowed to, as core counts them.
var orderStates = []traust.RunOrderState{
	traust.RunOrderStateQueued, traust.RunOrderStatePlanned, traust.RunOrderStateAwaitingApproval,
	traust.RunOrderStateLeased, traust.RunOrderStateSucceeded, traust.RunOrderStateFailed,
	traust.RunOrderStateDegraded, traust.RunOrderStateRefused, traust.RunOrderStateWithheld,
}

// allRuns is the Runs page: every order core records, narrowed by the state, asset and
// repository asked for, each linking to its own page.
func (h *Handler) allRuns(w http.ResponseWriter, r *http.Request) {
	asked := traust.OrderQueryRequest{}
	if asset := strings.TrimSpace(r.URL.Query().Get("asset")); asset != "" {
		asked.Asset = &asset
	}
	if state := strings.TrimSpace(r.URL.Query().Get("state")); state != "" {
		states := []traust.RunOrderState{traust.RunOrderState(state)} // core refuses a state it does not count
		asked.States = &states
	}
	if repository := strings.TrimSpace(r.URL.Query().Get("repository")); repository != "" {
		id, err := uuid.Parse(repository)
		if err != nil {
			web.Render(w, r, http.StatusBadRequest, web.Failure("That is not a repository."))
			return
		}
		asked.RepositoryId = &id
	}
	var orders []traust.OrderQueryItem
	for item, err := range h.coreRuns.QueryOrdersAll(r.Context(), asked) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		orders = append(orders, item)
	}
	web.Render(w, r, http.StatusOK, runsPage(orders))
}

// runDetail is one order's page: the order itself, the manifest it was planned with, the
// assessments it produced, and the model its asset's role resolves to.
func (h *Handler) runDetail(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "token")
	if !ok {
		return
	}
	order, err := h.coreRuns.GetOrder(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	manifest, err := h.coreRuns.GetOrderManifest(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	assessments, err := h.coreFindings.ListOrderAssessments(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	role := manifestRole(*manifest)
	var model *traust.RoleRouting
	if role != "" {
		if model, err = h.coreInventory.ResolveRole(r.Context(), &traust.ResolveRoleParams{Role: role}); err != nil {
			web.Fail(w, r, err)
			return
		}
	}
	web.Render(w, r, http.StatusOK, runDetail(*order, *manifest, assessments, role, model))
}

// manifestRole is the role a manifest names for its asset's model routing; an order carries no
// role of its own, and a manifest that names none leaves it unset.
func manifestRole(manifest traust.OrderManifest) string {
	role, _ := manifest["role"].(string)
	return role
}

// manifestKeys is a manifest's keys, ordered so the page reads the same every time.
func manifestKeys(manifest traust.OrderManifest) []string {
	return slices.Sorted(maps.Keys(manifest))
}

// manifestValue words one manifest entry: a string as core worded it, anything else as its JSON.
func manifestValue(manifest traust.OrderManifest, key string) string {
	if worded, ok := manifest[key].(string); ok {
		return worded
	}
	raw, err := json.Marshal(manifest[key])
	if err != nil {
		return "unreadable"
	}
	return string(raw)
}

// assessmentKind words an assessment's kind; one core worded unexpectedly says so.
func assessmentKind(a traust.Assessment) string {
	if kind, err := a.Kind.AsAssessmentKind(); err == nil {
		return string(kind)
	}
	return "unknown kind"
}
