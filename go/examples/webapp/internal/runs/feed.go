package runs

import (
	"net/http"
	"net/url"
	"slices"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// activity is core's ops feed, a page at a time: the page itself names no cursor, and a request
// naming core's answers only the next page's rows, for the more link to swap in.
func (h *Handler) activity(w http.ResponseWriter, r *http.Request) {
	after := r.URL.Query().Get("after")
	var params *traust.ReadFeedParams
	if after != "" {
		params = &traust.ReadFeedParams{After: &after}
	}
	page, err := h.coreRuns.ReadFeed(r.Context(), params)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	if after == "" {
		web.Render(w, r, http.StatusOK, activityPage(*page))
		return
	}
	web.Render(w, r, http.StatusOK, feedRows(*page))
}

// morePath links the feed's next page by core's cursor.
func morePath(cursor string) string {
	return "/activity?" + url.Values{"after": {cursor}}.Encode()
}

// newestFirst turns one of core's oldest-first pages around, so the latest change leads the list.
func newestFirst(events []traust.FeedEvent) []traust.FeedEvent {
	newest := append([]traust.FeedEvent(nil), events...)
	slices.Reverse(newest)
	return newest
}
