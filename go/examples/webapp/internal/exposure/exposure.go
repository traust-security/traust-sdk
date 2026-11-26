package exposure

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// exposure is the Exposure page: what core counts as exposed, by ecosystem; how that moved,
// bucket by bucket in core's order; the impacts it analysed; and how fresh its feeds are.
func (h *Handler) exposure(w http.ResponseWriter, r *http.Request) {
	rollups, err := h.coreExposure.SummarizeExposure(r.Context(), &traust.SummarizeExposureParams{
		GroupBy: traust.ExposureDimensionEcosystem,
	})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	timeline, err := h.coreExposure.ExposureTimeline(r.Context(), &traust.ExposureTimelineParams{})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	impacts := make([]traust.ImpactAnalysis, 0)
	for impact, err := range h.coreExposure.ListImpactsAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		impacts = append(impacts, impact)
	}
	feeds, err := h.coreExposure.FeedStatuses(r.Context())
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, exposurePage(rollups, timeline, impacts, feeds))
}

// impactInputs answers the advisory page's section for one ref: the commits core prepared
// against the advisory, and what their SBOMs hold of its packages.
func (h *Handler) impactInputs(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	ref, err := uuid.Parse(strings.TrimSpace(r.URL.Query().Get("ref")))
	if err != nil {
		web.Render(w, r, http.StatusBadRequest, web.Failure("That names no repository ref."))
		return
	}
	inputs, err := h.coreExposure.PrepareAdvisoryImpacts(r.Context(), id, &traust.PrepareAdvisoryImpactsParams{
		RepositoryRefId: ref,
	})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, impactInputs(ref, inputs))
}

// exportImpacts is the advisory's impact report: core's export, as a file under its name.
func (h *Handler) exportImpacts(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	export, err := h.coreExposure.ExportAdvisoryImpacts(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	document, err := json.Marshal(export)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", `attachment; filename="`+impactsFile(id)+`"`)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(document)
}

// impactsFile is the name the advisory's export downloads under.
func impactsFile(id uuid.UUID) string {
	return "impacts-" + id.String() + ".json"
}

// fetched words when core last fetched a feed; never is not yet.
func fetched(feed traust.FeedStatus) string {
	if feed.FetchedAt == nil {
		return "not fetched yet"
	}
	return feed.FetchedAt.Format(time.DateTime)
}

// errored words a feed's last error, with its separator when there was one.
func errored(feed traust.FeedStatus) string {
	if err := web.Text(feed.LastError); err != "" {
		return " · errored: " + err
	}
	return ""
}
