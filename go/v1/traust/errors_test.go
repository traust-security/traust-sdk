package traust

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestProblemErrorMatchesStatusSentinelsAndDenyCodes(t *testing.T) {
	budget := BudgetExhausted
	tests := []struct {
		name    string
		problem *ProblemError
		target  error
		want    bool
	}{
		{"401", &ProblemError{StatusCode: http.StatusUnauthorized}, ErrUnauthorized, true},
		{"403", &ProblemError{StatusCode: http.StatusForbidden}, ErrForbidden, true},
		{"404", &ProblemError{StatusCode: http.StatusNotFound}, ErrNotFound, true},
		{"409", &ProblemError{StatusCode: http.StatusConflict}, ErrConflict, true},
		{"422", &ProblemError{StatusCode: http.StatusUnprocessableEntity}, ErrInvalid, true},
		{"another status's sentinel", &ProblemError{StatusCode: http.StatusNotFound}, ErrConflict, false},
		{"an unmapped status is no sentinel", &ProblemError{StatusCode: http.StatusInternalServerError}, ErrNotFound, false},
		{"an unmapped status is not nil", &ProblemError{StatusCode: http.StatusInternalServerError}, nil, false},
		{"the problem's code", &ProblemError{StatusCode: 403, Problem: &Problem{Code: &budget}}, BudgetExhausted, true},
		{"an issue's code", &ProblemError{StatusCode: 403, Problem: &Problem{Issues: &[]DenyIssue{{Code: RoleDenied}}}}, RoleDenied, true},
		{"a code core did not send", &ProblemError{StatusCode: 403, Problem: &Problem{Code: &budget}}, RoleDenied, false},
		{"no problem document", &ProblemError{StatusCode: 403, Body: "denied"}, RoleDenied, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fmt.Errorf("Submit: %w", tt.problem)
			if got := errors.Is(err, tt.target); got != tt.want {
				t.Fatalf("errors.Is(%v, %v) = %v, want %v", err, tt.target, got, tt.want)
			}
		})
	}
}
