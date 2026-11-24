package exposure

import (
	"errors"
	"net/http"
	"strings"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// packages is the Packages page: the package core knows under the ecosystem and name asked for,
// as core recorded it. A miss says so; anything else is core's to word.
func (h *Handler) packages(w http.ResponseWriter, r *http.Request) {
	ecosystem := strings.TrimSpace(r.URL.Query().Get("ecosystem"))
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if ecosystem == "" || name == "" {
		web.Render(w, r, http.StatusOK, packagesPage(ecosystem, name, nil))
		return
	}
	found, err := h.coreExposure.FindPackage(r.Context(), &traust.FindPackageParams{
		Ecosystem: traust.Ecosystem(ecosystem), Name: name,
	})
	if err != nil {
		if errors.Is(err, traust.ErrNotFound) {
			web.Render(w, r, http.StatusNotFound, web.Failure("No package matches that ecosystem and name."))
			return
		}
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, packagesPage(ecosystem, name, found))
}

// packagePage is one package's page: the refs whose SBOMs hold it, narrowed by the version asked
// for, each linking to its own page.
func (h *Handler) packagePage(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	version := strings.TrimSpace(r.URL.Query().Get("version"))
	var params *traust.DependentsParams
	if version != "" {
		params = &traust.DependentsParams{Version: &version}
	}
	dependents, err := h.coreExposure.Dependents(r.Context(), id, params)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, packagePage(id, version, dependents))
}

// refDependencies answers the ref page's dependency section: the packages its SBOMs hold, as
// core resolved them.
func (h *Handler) refDependencies(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	dependencies, err := h.coreExposure.Dependencies(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, dependencyList(dependencies))
}
