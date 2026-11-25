package web

import (
	"errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// PathID is the {id} or {token} a route names, or a 400 already answered.
func PathID(w http.ResponseWriter, r *http.Request, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(r.PathValue(name))
	if err != nil {
		Render(w, r, http.StatusBadRequest, Failure("That link names nothing."))
	}
	return id, err == nil
}

// Status passes core's refusal through; anything else is core being unavailable.
func Status(err error) int {
	if p, ok := errors.AsType[*traust.ProblemError](err); ok && p.StatusCode < http.StatusInternalServerError {
		return p.StatusCode
	}
	return http.StatusBadGateway
}

// Fail answers with core's refusal and its reason, which core words for people; anything else is
// core being unavailable.
func Fail(w http.ResponseWriter, r *http.Request, err error) {
	code := Status(err)
	message := http.StatusText(code)
	if p, ok := errors.AsType[*traust.ProblemError](err); ok && p.Problem != nil && code < http.StatusInternalServerError {
		message = p.Problem.Title
		if p.Problem.Detail != "" {
			message += ": " + p.Problem.Detail
		}
	}
	Render(w, r, code, Failure(message))
}

// Render answers with a page or a fragment.
func Render(w http.ResponseWriter, r *http.Request, code int, view templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	_ = view.Render(r.Context(), w)
}
