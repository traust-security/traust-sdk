package traust

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
)

// Refusals by HTTP status, for errors.Is. A deny code core sends matches too:
// errors.Is(err, traust.BudgetExhausted).
var (
	ErrUnauthorized = errors.New("traust: unauthorized")
	ErrForbidden    = errors.New("traust: forbidden")
	ErrNotFound     = errors.New("traust: not found")
	ErrConflict     = errors.New("traust: conflict")
	ErrInvalid      = errors.New("traust: invalid request")
)

var statusErrors = map[int]error{
	http.StatusUnauthorized:        ErrUnauthorized,
	http.StatusForbidden:           ErrForbidden,
	http.StatusNotFound:            ErrNotFound,
	http.StatusConflict:            ErrConflict,
	http.StatusUnprocessableEntity: ErrInvalid,
}

// ProblemError reports a non-retryable failed response. Problem carries the
// RFC 7807 document when the server sent one; otherwise Body holds the raw
// response.
type ProblemError struct {
	StatusCode int
	Problem    *Problem
	Body       string
}

func (e *ProblemError) Error() string {
	if e.Problem != nil {
		detail := e.Problem.Detail
		if detail == "" {
			detail = e.Problem.Title
		}
		request := ""
		if e.Problem.RequestId != "" {
			request = " (request " + e.Problem.RequestId + ")"
		}
		return fmt.Sprintf("traust-core: %d %s: %s%s", e.StatusCode, e.Problem.Title, detail, request)
	}
	return fmt.Sprintf("traust-core: HTTP %d", e.StatusCode)
}

// Is matches the status sentinel, or a DenyCode carried as the problem's code or an issue's.
func (e *ProblemError) Is(target error) bool {
	code, ok := target.(DenyCode)
	if !ok {
		sentinel, known := statusErrors[e.StatusCode]
		return known && sentinel == target
	}
	if e.Problem == nil {
		return false
	}
	if e.Problem.Code != nil && *e.Problem.Code == code {
		return true
	}
	return e.Problem.Issues != nil && slices.ContainsFunc(*e.Problem.Issues, func(i DenyIssue) bool { return i.Code == code })
}

// Error makes a DenyCode an errors.Is target.
func (c DenyCode) Error() string { return "traust: denied: " + string(c) }
