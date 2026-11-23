package inventory

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// refName words a ref: its kind and its name; a nameless one says so.
func refName(ref traust.RepositoryRef) string {
	if name := web.Text(ref.Name); name != "" {
		return web.Text(ref.Kind) + " " + name
	}
	return web.Text(ref.Kind) + " (unnamed)"
}

// refs answers the repository page's ref section: its branches and tags, the default marked.
func (h *Handler) refs(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	refs, err := h.core.ListRefs(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, refList(id.String(), refs))
}

// ref answers a ref's page: the branch or tag itself, as core knows it.
func (h *Handler) ref(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.core.GetRef(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, refDetail(*found))
}

// recentRefs answers the home page's aside: the latest refs across the repositories it lists.
func (h *Handler) recentRefs(w http.ResponseWriter, r *http.Request) {
	ids := make([]uuid.UUID, 0)
	for repo, err := range h.core.ListRepositoriesAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		if repo.Id != nil {
			ids = append(ids, *repo.Id)
		}
	}
	if len(ids) == 0 {
		web.Render(w, r, http.StatusOK, recentRefs(nil))
		return
	}
	refs, err := h.core.ListRefsForRepositories(r.Context(), &traust.ListRefsForRepositoriesParams{RepositoryIds: &ids})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, recentRefs(refs))
}

// registerRef adds a ref to the repository; core decides who may register one. The section
// re-lists, so the new ref appears among its peers.
func (h *Handler) registerRef(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	kind := traust.RefKind(strings.TrimSpace(r.PostFormValue("kind")))
	name := strings.TrimSpace(r.PostFormValue("name"))
	sha := strings.TrimSpace(r.PostFormValue("sha"))
	if (kind != traust.Branch && kind != traust.Tag) || name == "" || sha == "" {
		web.Render(w, r, http.StatusBadRequest, web.Failure("A ref needs a kind (branch or tag), a name and a sha."))
		return
	}
	if _, err := h.core.RegisterRef(r.Context(), id, traust.RefRegistration{Kind: &kind, Name: &name, Sha: sha}); err != nil {
		web.Fail(w, r, err)
		return
	}
	h.refs(w, r)
}

// makeDefault names one of the repository's branches its default; core decides who may name one.
func (h *Handler) makeDefault(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	branch := strings.TrimSpace(r.PostFormValue("branch"))
	if branch == "" {
		web.Render(w, r, http.StatusBadRequest, web.Failure("The branch is required."))
		return
	}
	if _, err := h.core.MarkDefaultBranch(r.Context(), id, traust.DefaultRefNaming{Branch: branch, Sha: r.PostFormValue("sha")}); err != nil {
		web.Fail(w, r, err)
		return
	}
	h.refs(w, r)
}
