package findings

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// exposures are the states a finding's folded disposition can be narrowed to, as core counts them.
var exposures = []traust.FindingExposure{traust.FindingExposureOpen, traust.FindingExposureClosed, traust.FindingExposureFalsePositive, traust.FindingExposureHardening}

// severities are the effective severities a person may narrow the list to.
var severities = []traust.Severity{traust.SeverityCritical, traust.SeverityHigh, traust.SeverityMedium, traust.SeverityLow, traust.SeverityInformational}

// section is one repository's slice of the Findings page: the repository and the findings core
// returned for it.
type section struct {
	Repository traust.Repository
	Found      []traust.FindingView
}

// allFindings is the Findings page: every finding across the repositories the person can see,
// narrowed by the severity and state asked for, each repository grouped with a form to look one
// up by its fingerprint.
func (h *Handler) allFindings(w http.ResponseWriter, r *http.Request) {
	var repos []traust.Repository
	for repo, err := range h.core.ListRepositoriesAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		repos = append(repos, repo)
	}
	sections := make([]section, len(repos))
	ids := make([]uuid.UUID, 0, len(repos))
	at := make(map[uuid.UUID]int, len(repos))
	for i, repo := range repos {
		sections[i].Repository = repo
		if repo.Id != nil {
			ids = append(ids, *repo.Id)
			at[*repo.Id] = i
		}
	}
	if len(ids) > 0 { // core refuses a set with no repositories, so an empty scope asks nothing
		params := traust.ListFindingsInParams{RepositoryIds: &ids}
		if asked := trimmed(r.URL.Query()["severity"]); len(asked) > 0 {
			params.Severity = &asked
		}
		if state := strings.TrimSpace(r.URL.Query().Get("exposure")); state != "" {
			exposure := traust.FindingExposure(state)
			params.Exposure = &exposure
		}
		for f, err := range h.core.ListFindingsInAll(r.Context(), &params) {
			if err != nil {
				web.Fail(w, r, err)
				return
			}
			if where, ok := at[f.Finding.RepositoryId]; ok {
				sections[where].Found = append(sections[where].Found, f)
			}
		}
	}
	web.Render(w, r, http.StatusOK, findingsPage(sections))
}

// summary answers the repository page's summary section: the repository's findings counted by
// severity, largest group first, each severity's states.
func (h *Handler) summary(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	rollups, err := h.core.SummarizeFindings(r.Context(), id, &traust.SummarizeFindingsParams{GroupBy: traust.FindingDimensionSeverity})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, summaryList(rollups))
}

// timeline answers the repository page's timeline section: the repository's findings replayed
// bucket by bucket, in core's order.
func (h *Handler) timeline(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	series, err := h.core.FindingTimeline(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, timelineList(series))
}

// byFingerprint answers the Findings page's lookup: the finding core knows by that fingerprint,
// as its own page; a miss says so, and anything else is core's to word.
func (h *Handler) byFingerprint(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	fingerprint := strings.TrimSpace(r.URL.Query().Get("fingerprint"))
	if fingerprint == "" {
		web.Render(w, r, http.StatusBadRequest, web.Failure("A fingerprint is required."))
		return
	}
	found, err := h.core.GetFindingByFingerprint(r.Context(), id, &traust.GetFindingByFingerprintParams{Fingerprint: fingerprint})
	if err != nil {
		if errors.Is(err, traust.ErrNotFound) {
			web.Render(w, r, http.StatusNotFound, web.Failure("No finding matches that fingerprint."))
			return
		}
		web.Fail(w, r, err)
		return
	}
	http.Redirect(w, r, "/findings/"+found.Finding.Id.String(), http.StatusSeeOther)
}

// trimmed is the query's values, trimmed and without empties, so a filter only reaches core when
// the person asked for one.
func trimmed(values []string) []string {
	var kept []string
	for _, value := range values {
		if v := strings.TrimSpace(value); v != "" {
			kept = append(kept, v)
		}
	}
	return kept
}

// rollupKey words a summary's group; one core left keyless says none.
func rollupKey(r traust.FindingRollup) string {
	if key := web.Text(r.Key); key != "" {
		return key
	}
	return "none"
}

// openCount is a timeline bucket's open findings, summed across severities.
func openCount(p traust.FindingSeriesPoint) int {
	open := 0
	for _, n := range p.OpenBySeverity {
		open += n
	}
	return open
}
