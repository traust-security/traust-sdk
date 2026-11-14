package traust_test

import (
	"context"
	"errors"
	"net/http"
	"reflect"
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
	format := traust.Html
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
	format := traust.Csv
	_, err := core.Client().GetInsightReport(context.Background(), traust.ExecutiveSummary,
		&traust.GetInsightReportParams{Format: &format})
	var problem *traust.ProblemError
	if !errors.As(err, &problem) || problem.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("err = %v", err)
	}
}
