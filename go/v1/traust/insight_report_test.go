package traust_test

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

// A report is the rendered file, asked for by format, scoped like the view (panels included).
func TestGetInsightReportAsksForTheFormatAndThePanels(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/insights/executive-summary/report", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte("<!doctype html><title>Executive summary</title>"))
	}))
	format := traust.ReportFormatHtml
	panels := []traust.ExecutivePanel{traust.ExecutivePanelIsolation, traust.ExecutivePanelIac}
	body, err := core.Client().GetInsightReport(context.Background(), traust.ExecutiveSummary,
		&traust.GetInsightReportParams{Format: &format, ExecutivePanel: &panels})
	if err != nil {
		t.Fatalf("GetInsightReport: %v", err)
	}
	if string(body) != "<!doctype html><title>Executive summary</title>" {
		t.Fatalf("body = %q", body)
	}
	query := core.Requests()[0].Query
	if query.Get("format") != "html" || !reflect.DeepEqual(query["executive_panel"], []string{"isolation", "iac"}) {
		t.Fatalf("query = %v", query)
	}
}

// Core's refusal (csv for a view without an export table) is a *ProblemError.
func TestGetInsightReportReturnsCoresRefusalAsAProblem(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/insights/executive-summary/report",
		traustest.Problem(http.StatusUnprocessableEntity, "executive-summary has no export table"))
	format := traust.ReportFormatCsv
	_, err := core.Client().GetInsightReport(context.Background(), traust.ExecutiveSummary,
		&traust.GetInsightReportParams{Format: &format})
	var problem *traust.ProblemError
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("err = %v", err)
	}
}

func askReport(t *testing.T, core *traustest.Server) ([]byte, error) {
	t.Helper()
	format := traust.ReportFormatMarkdown
	return core.Client().GetInsightReport(context.Background(), traust.ExecutiveSummary,
		&traust.GetInsightReportParams{Format: &format})
}

// Every refusal is a *ProblemError naming the call; a body that is no problem still is one.
func TestGetInsightReportRefusalsAreProblemsNamingTheCall(t *testing.T) {
	for _, tc := range []struct {
		name    string
		handler http.Handler
		status  int
	}{
		{"server error", traustest.Problem(http.StatusInternalServerError, "boom"), http.StatusInternalServerError},
		{"unauthenticated", traustest.Problem(http.StatusUnauthorized, "no token"), http.StatusUnauthorized},
		{"unknown view", traustest.Problem(http.StatusNotFound, "no such view"), http.StatusNotFound},
		{"malformed problem", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/problem+json")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("{not json"))
		}), http.StatusInternalServerError},
		{"plain text", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "gateway trouble", http.StatusInternalServerError)
		}), http.StatusInternalServerError},
	} {
		t.Run(tc.name, func(t *testing.T) {
			core := traustest.New(t)
			core.Handle("GET /v1/insights/executive-summary/report", tc.handler)
			body, err := askReport(t, core)
			var problem *traust.ProblemError
			if body != nil || !errors.As(err, &problem) || problem.StatusCode != tc.status {
				t.Fatalf("body %q, err %v", body, err)
			}
			if !strings.Contains(err.Error(), "GetInsightReport") {
				t.Fatalf("err %q does not name the call", err)
			}
		})
	}
}

// A success other than 200 carries no file: an error, never empty report bytes.
func TestGetInsightReportRefusesAnEmptySuccess(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/insights/executive-summary/report", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	body, err := askReport(t, core)
	if body != nil || err == nil || !strings.Contains(err.Error(), "HTTP 204, want 200") {
		t.Fatalf("body %q, err %v", body, err)
	}
}

// A transport failure is an error naming the call, not a *ProblemError.
func TestGetInsightReportNamesATransportFailure(t *testing.T) {
	core := traustest.New(t)
	client := core.Client()
	core.Close()
	format := traust.ReportFormatMarkdown
	body, err := client.GetInsightReport(context.Background(), traust.ExecutiveSummary,
		&traust.GetInsightReportParams{Format: &format})
	var problem *traust.ProblemError
	if body != nil || err == nil || errors.As(err, &problem) || !strings.Contains(err.Error(), "GetInsightReport") {
		t.Fatalf("body %q, err %v", body, err)
	}
}
