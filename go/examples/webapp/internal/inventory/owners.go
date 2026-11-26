package inventory

import (
	"net/http"
	"strings"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// owners answers a repository page's lazy owners section: the teams owning it now.
func (h *Handler) owners(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	teams, err := h.coreInventory.ListOwners(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, ownerList(id.String(), teams))
}

// setOwners mirrors a repository's owners from the form's teams; core decides who may set them.
// The section re-lists, so the teams appear among their peers.
func (h *Handler) setOwners(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	teams := teamNames(r)
	if len(teams) == 0 {
		web.Render(w, r, http.StatusBadRequest, web.Failure("Set owners needs at least one team."))
		return
	}
	if _, err := h.coreInventory.SetOwners(r.Context(), id, traust.Owners{Teams: teams}); err != nil {
		web.Fail(w, r, err)
		return
	}
	h.owners(w, r)
}

// teamNames reads the set-owners form: one team per line, blank lines dropped.
func teamNames(r *http.Request) []string {
	var names []string
	for line := range strings.Lines(r.PostFormValue("teams")) {
		if name := strings.TrimSpace(line); name != "" {
			names = append(names, name)
		}
	}
	return names
}

// team answers a team's page: the repositories it owns now, each read as core words it.
func (h *Handler) team(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	ids, err := h.coreInventory.ListOwnedRepositories(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	repos := make([]traust.Repository, 0, len(ids))
	for _, repository := range ids {
		found, err := h.coreInventory.GetRepository(r.Context(), repository)
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		repos = append(repos, *found)
	}
	web.Render(w, r, http.StatusOK, teamDetail(repos))
}
