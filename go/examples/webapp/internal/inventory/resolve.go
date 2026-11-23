package inventory

import (
	"net/http"
	"strings"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// resolve answers the home page's lookup: each pasted URL's repository, or its honest absence.
func (h *Handler) resolve(w http.ResponseWriter, r *http.Request) {
	urls := pastedURLs(r)
	if len(urls) == 0 {
		web.Render(w, r, http.StatusBadRequest, web.Failure("Paste at least one repository URL."))
		return
	}
	results, err := h.core.ResolveRepositories(r.Context(), traust.RepositoryResolutionRequest{Urls: urls})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, resolutions(results))
}

// pastedURLs reads the lookup's textarea: one repository URL per line, blank lines dropped.
func pastedURLs(r *http.Request) []string {
	var urls []string
	for line := range strings.Lines(r.PostFormValue("urls")) {
		if url := strings.TrimSpace(line); url != "" {
			urls = append(urls, url)
		}
	}
	return urls
}

// correctURL moves a repository to its new spelling; core decides who may correct.
func (h *Handler) correctURL(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	url := strings.TrimSpace(r.PostFormValue("url"))
	if url == "" {
		web.Render(w, r, http.StatusBadRequest, web.Failure("The new URL is required."))
		return
	}
	repo, err := h.core.CorrectRepositoryUrl(r.Context(), id, traust.UrlCorrection{Url: url})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, corrected(*repo))
}

// verify asks core to observe a repository; core decides who may verify.
func (h *Handler) verify(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	verification, err := h.core.VerifyRepository(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, verified(*verification))
}

// standing words a repository's disposition; an undecided one says so.
func standing(d *traust.RepositoryDisposition) string {
	if d == nil {
		return "undecided"
	}
	return string(*d)
}
