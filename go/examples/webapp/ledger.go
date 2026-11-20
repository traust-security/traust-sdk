package main

import "net/http"

func (a *app) ledger(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	layers, err := a.core.ListLayers(r.Context(), id)
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, ledger(id.String(), layers))
}

// checkpoint restamps and signs the repository's layer; core lets only admins.
func (a *app) checkpoint(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	if _, err := a.core.Checkpoint(r.Context(), id); err != nil {
		fail(w, r, err)
		return
	}
	a.ledger(w, r)
}

func (a *app) verify(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	result, err := a.core.VerifyLayer(r.Context(), id)
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, verified(*result))
}
