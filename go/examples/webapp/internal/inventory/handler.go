// Package inventory is the app's home: the person's dashboard, onboarding repositories, and a
// repository's page.
package inventory

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the person's repositories and their pages.
type Handler struct{ core *traust.Client }

// tally is one row of a count: a label and its number, as core answered them.
type tally struct {
	Label string
	Count string
}

// card is one dashboard query's answer.
type card struct {
	Title string
	Rows  []tally
}

// New builds the feature's handler.
func New(core *traust.Client) *Handler { return &Handler{core: core} }

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", h.home)
	mux.HandleFunc("GET /onboard", web.SignedIn(h.onboardForm))
	mux.HandleFunc("POST /onboard/plan", web.SignedIn(h.planOnboarding))
	mux.HandleFunc("POST /onboard", web.SignedIn(h.onboard))
	mux.HandleFunc("GET /repositories/{id}", web.SignedIn(h.repository))
	mux.HandleFunc("POST /repositories/resolve", web.SignedIn(h.resolve))
	mux.HandleFunc("POST /repositories/{id}/url", web.SignedIn(h.correctURL))
	mux.HandleFunc("POST /repositories/{id}/verify", web.SignedIn(h.verify))
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(web.SessionCookie)
	if err != nil {
		web.Render(w, r, http.StatusOK, web.SignIn(""))
		return
	}
	ctx := traust.OnBehalfOf(r.Context(), cookie.Value)
	me, err := h.core.Whoami(ctx)
	if errors.Is(err, traust.ErrUnauthorized) {
		http.SetCookie(w, &http.Cookie{Name: web.SessionCookie, MaxAge: -1})
		web.Render(w, r, http.StatusUnauthorized, web.SignIn("Your session ended. Sign in again."))
		return
	}
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	var repos []traust.Repository
	for repo, err := range h.core.ListRepositoriesAll(ctx, nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		repos = append(repos, repo)
	}
	cards, err := h.dashboard(ctx, repos)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, home(*me, repos, cards))
}

// dashboard is the home page's cards: one dashboard query each, over the person's repositories.
func (h *Handler) dashboard(ctx context.Context, repos []traust.Repository) ([]card, error) {
	ids := make([]uuid.UUID, 0, len(repos))
	for _, repo := range repos {
		if repo.Id != nil {
			ids = append(ids, *repo.Id)
		}
	}
	if len(ids) == 0 {
		return nil, nil
	}
	queries := []struct {
		title   string
		subject traust.Subject
		measure traust.Measure
		by      traust.Dimension
	}{
		{"Findings by severity", traust.SubjectFindings, traust.MeasureFindings, traust.DimensionSeverity},
		{"Findings by state", traust.SubjectFindings, traust.MeasureFindings, traust.DimensionState},
		{"Review queue by status", traust.SubjectReviewItems, traust.MeasureReviewItems, traust.DimensionStatus},
	}
	cards := make([]card, 0, len(queries))
	for _, q := range queries {
		result, err := h.core.QueryDashboard(ctx, traust.DashboardQuery{
			Subject: q.subject, Measures: []traust.Measure{q.measure}, GroupBy: &[]traust.Dimension{q.by}, RepositoryIds: &ids,
		})
		if err != nil {
			return nil, err
		}
		c := card{Title: q.title}
		for _, row := range result.Rows {
			label, count := "none", "0"
			if g := row.Group(q.by); g != nil {
				label = *g
			}
			if n := row.Measure(q.measure); n != nil {
				count = fmt.Sprint(*n)
			}
			c.Rows = append(c.Rows, tally{label, count})
		}
		cards = append(cards, c)
	}
	return cards, nil
}

func (h *Handler) onboardForm(w http.ResponseWriter, r *http.Request) {
	web.Render(w, r, http.StatusOK, onboardForm())
}

// onboarding reads the form: a product, its label when new, and one repository URL per line.
// It plans no baseline: an order onboarding plans carries no pack, so no pinned runner may run
// it; the engine plans each new repository's scan under its pack instead.
func onboarding(r *http.Request) traust.OnboardingRequest {
	request := traust.OnboardingRequest{Product: strings.TrimSpace(r.PostFormValue("product"))}
	if label := strings.TrimSpace(r.PostFormValue("label")); label != "" {
		request.Label = &label
	}
	var specs []traust.RepositorySpec
	for line := range strings.Lines(r.PostFormValue("repositories")) {
		if url := strings.TrimSpace(line); url != "" {
			specs = append(specs, traust.RepositorySpec{Url: url})
		}
	}
	request.Repositories = &specs
	return request
}

func (h *Handler) planOnboarding(w http.ResponseWriter, r *http.Request) {
	plan, err := h.core.PlanOnboarding(r.Context(), onboarding(r))
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, onboardPlan(*plan))
}

func (h *Handler) onboard(w http.ResponseWriter, r *http.Request) {
	if _, err := h.core.Onboard(r.Context(), onboarding(r)); err != nil {
		web.Fail(w, r, err)
		return
	}
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) repository(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	repo, err := h.core.GetRepository(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	ref, err := h.core.GetCurrentDefaultRef(r.Context(), id)
	if err != nil && !errors.Is(err, traust.ErrNotFound) {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, repository(*repo, ref))
}
