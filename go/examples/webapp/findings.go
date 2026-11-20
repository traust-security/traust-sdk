package main

import (
	"net/http"
	"strings"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

func (a *app) findings(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	var found []traust.FindingView
	for f, err := range a.core.ListFindingsAll(r.Context(), id, nil) {
		if err != nil {
			fail(w, r, err)
			return
		}
		found = append(found, f)
	}
	render(w, r, http.StatusOK, findingList(found))
}

func (a *app) finding(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	f, err := a.core.GetFinding(r.Context(), id)
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, finding(*f))
}

func (a *app) decide(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	validity := traust.Validity(r.PostFormValue("validity"))
	if _, err := a.core.Decide(r.Context(), id, traust.DecisionBody{
		Validity: &validity, Rationale: strings.TrimSpace(r.PostFormValue("rationale")),
	}); err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, decided())
}
