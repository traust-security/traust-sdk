package ledger

import (
	"net/http"
	"net/url"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// layer answers a layer's page: what core sealed it with, and the events it holds, a page at a
// time.
func (h *Handler) layer(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.coreLedger.GetLayer(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	page, err := h.coreLedger.ListEvents(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, layerPage(id, *found, *page))
}

// events answers the layer's next page of events; the page's more link asks for it by core's
// cursor.
func (h *Handler) events(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	var params *traust.ListEventsParams
	if cursor := r.URL.Query().Get("cursor"); cursor != "" {
		params = &traust.ListEventsParams{Cursor: &cursor}
	}
	page, err := h.coreLedger.ListEvents(r.Context(), id, params)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, layerEvents(id, *page))
}

// eventsPath links the layer's next page of events by core's cursor.
func eventsPath(id string, cursor string) string {
	return "/layers/" + id + "/events?" + url.Values{"cursor": {cursor}}.Encode()
}
