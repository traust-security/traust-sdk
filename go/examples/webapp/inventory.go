package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sessionCookie)
	if err != nil {
		render(w, r, http.StatusOK, signIn(""))
		return
	}
	ctx := traust.OnBehalfOf(r.Context(), cookie.Value)
	me, err := a.core.Whoami(ctx)
	if errors.Is(err, traust.ErrUnauthorized) {
		http.SetCookie(w, &http.Cookie{Name: sessionCookie, MaxAge: -1})
		render(w, r, http.StatusUnauthorized, signIn("Your session ended. Sign in again."))
		return
	}
	if err != nil {
		fail(w, r, err)
		return
	}
	var repos []traust.Repository
	for repo, err := range a.core.ListRepositoriesAll(ctx, nil) {
		if err != nil {
			fail(w, r, err)
			return
		}
		repos = append(repos, repo)
	}
	severities, err := a.findingsBySeverity(ctx, repos)
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, home(*me, repos, severities))
}

// findingsBySeverity is one dashboard query over the person's repositories.
func (a *app) findingsBySeverity(ctx context.Context, repos []traust.Repository) ([]tally, error) {
	ids := make([]uuid.UUID, 0, len(repos))
	for _, repo := range repos {
		if repo.Id != nil {
			ids = append(ids, *repo.Id)
		}
	}
	if len(ids) == 0 {
		return nil, nil
	}
	result, err := a.core.QueryDashboard(ctx, traust.DashboardQuery{
		Subject:       traust.SubjectFindings,
		Measures:      []traust.Measure{traust.MeasureFindings},
		GroupBy:       &[]traust.Dimension{traust.DimensionSeverity},
		RepositoryIds: &ids,
	})
	if err != nil {
		return nil, err
	}
	tallies := make([]tally, 0, len(result.Rows))
	for _, row := range result.Rows {
		label, count := "unrated", "0"
		if s := row.Group(traust.DimensionSeverity); s != nil {
			label = *s
		}
		if n := row.Measure(traust.MeasureFindings); n != nil {
			count = fmt.Sprint(*n)
		}
		tallies = append(tallies, tally{label, count})
	}
	return tallies, nil
}

func (a *app) onboardForm(w http.ResponseWriter, r *http.Request) {
	render(w, r, http.StatusOK, onboardForm())
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

func (a *app) planOnboarding(w http.ResponseWriter, r *http.Request) {
	plan, err := a.core.PlanOnboarding(r.Context(), onboarding(r))
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, onboardPlan(*plan))
}

func (a *app) onboard(w http.ResponseWriter, r *http.Request) {
	if _, err := a.core.Onboard(r.Context(), onboarding(r)); err != nil {
		fail(w, r, err)
		return
	}
	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusNoContent)
}

func (a *app) repository(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	repo, err := a.core.GetRepository(r.Context(), id)
	if err != nil {
		fail(w, r, err)
		return
	}
	ref, err := a.core.GetCurrentDefaultRef(r.Context(), id)
	if err != nil && !errors.Is(err, traust.ErrNotFound) {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, repository(*repo, ref))
}
