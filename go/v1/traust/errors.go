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

// Named failures of LoadConfig, for errors.Is.
var (
	// ErrNoConfigFile: the config file is needed to fill in the Config but missing.
	ErrNoConfigFile = errors.New("traust: no config file")
	// ErrNoProfile: nothing selected a profile — no Profile option, no TRAUST_PROFILE, no default.
	ErrNoProfile = errors.New("traust: no profile selected")
	// ErrUnknownProfile: the selected profile is not in the config file.
	ErrUnknownProfile = errors.New("traust: unknown profile")
	// ErrProfileBothCredentials: a profile names a token_file and a client identity.
	ErrProfileBothCredentials = errors.New("names both a token_file and a client identity")
	// ErrProfileNoCredential: a profile names neither a token_file nor a client identity.
	ErrProfileNoCredential = errors.New("names no credential")
)

// ProfileKeyError reports a profile in the config file that is missing its url or one of the
// credential keys.
type ProfileKeyError struct {
	Profile string
	Key     string
}

func (e *ProfileKeyError) Error() string {
	return fmt.Sprintf("traust: profile %q has no %s", e.Profile, e.Key)
}

// ProfileCredentialError reports a profile in the config file that does not name exactly one
// credential: token_file, or the client-credentials trio. Err is ErrProfileBothCredentials or
// ErrProfileNoCredential.
type ProfileCredentialError struct {
	Profile string
	Err     error
}

func (e *ProfileCredentialError) Error() string {
	return fmt.Sprintf("traust: profile %q %s", e.Profile, e.Err)
}

func (e *ProfileCredentialError) Unwrap() error { return e.Err }
