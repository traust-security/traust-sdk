package traust_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

func TestQueryDashboardPostsQueryAndDecodesRows(t *testing.T) {
	core := traustest.New(t)
	count := float32(7)
	sev := "high"
	core.Handle("POST /v1/dashboards:query", traustest.JSON(http.StatusOK, traust.DashboardResults{
		Subject: traust.SubjectFindings,
		Rows: []traust.DashboardRow{{
			Groups:   map[string]*string{"severity": &sev},
			Measures: map[string]*float32{"findings": &count},
		}},
	}))

	groupBy := []traust.Dimension{traust.DimensionSeverity}
	results, err := core.Client().QueryDashboard(context.Background(), traust.DashboardQuery{
		Subject:  traust.SubjectFindings,
		Measures: []traust.Measure{traust.MeasureFindings},
		GroupBy:  &groupBy,
	})
	if err != nil {
		t.Fatalf("QueryDashboard: %v", err)
	}
	if results.Subject != traust.SubjectFindings || len(results.Rows) != 1 {
		t.Fatalf("unexpected results: %+v", results)
	}
	if got := results.Rows[0].Measures["findings"]; got == nil || *got != 7 {
		t.Fatalf("findings measure = %v, want 7", got)
	}
}

func TestRequestReportHandlesReady200AndQueued202(t *testing.T) {
	for _, tc := range []struct {
		name   string
		status int
		state  traust.ReportStatus
	}{
		{"ready 200", http.StatusOK, traust.ReportStatusReady},
		{"queued 202", http.StatusAccepted, traust.ReportStatusQueued},
	} {
		t.Run(tc.name, func(t *testing.T) {
			core := traustest.New(t)
			orderID := uuid.New()
			now := time.Now().UTC().Truncate(time.Second)
			core.Handle("POST /v1/reports", traustest.JSON(tc.status, traust.Report{
				OrderId:     orderID,
				Report:      traust.FindingsRollup,
				Format:      traust.ReportFormatJson,
				Status:      tc.state,
				RequestedAt: now,
			}))

			format := traust.ReportFormatJson
			report, err := core.Client().RequestReport(context.Background(), traust.ReportRequest{
				Report: traust.ViewNameOf[traust.FindingsRollupInsight](),
				Format: &format,
			})
			if err != nil {
				t.Fatalf("RequestReport: %v", err)
			}
			if report.OrderId != orderID || report.Status != tc.state {
				t.Fatalf("got order=%s status=%s, want %s/%s", report.OrderId, report.Status, orderID, tc.state)
			}
		})
	}
}

func TestGetReportFileReturnsRawBytesAcrossFormats(t *testing.T) {
	for _, tc := range []struct {
		name        string
		contentType string
		body        string
	}{
		{"markdown", "text/markdown; charset=utf-8", "# Executive Summary\n"},
		{"html", "text/html; charset=utf-8", "<h1>Executive Summary</h1>"},
		{"json", "application/json", `{"repositories":[],"total":{"open":{},"closed":{}}}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			core := traustest.New(t)
			orderID := uuid.New()
			core.Handle("GET /v1/reports/{order_id}/file", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(tc.body))
			}))
			got, err := core.Client().GetReportFile(context.Background(), orderID)
			if err != nil {
				t.Fatalf("GetReportFile: %v", err)
			}
			if string(got) != tc.body {
				t.Fatalf("got %q, want %q", string(got), tc.body)
			}
		})
	}
}

func TestGetReportFileAsDecodesStrictlyAndRejectsUnknownFields(t *testing.T) {
	core := traustest.New(t)
	goodID := uuid.New()
	badID := uuid.New()
	notReadyID := uuid.New()
	core.Handle("GET /v1/reports/{order_id}/file", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.PathValue("order_id") {
		case goodID.String():
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"repositories":[],"total":{"open":{"critical":2}}}`))
		case badID.String():
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"unexpected_field":true}`))
		case notReadyID.String():
			traustest.Problem(http.StatusConflict, "report not ready").ServeHTTP(w, r)
		}
	}))

	rollup, err := traust.GetReportFileAs[traust.FindingsRollupInsight](context.Background(), core.Client(), goodID)
	if err != nil {
		t.Fatalf("GetReportFileAs: %v", err)
	}
	if rollup.Total.Open.Critical != 2 {
		t.Fatalf("critical = %v, want 2", rollup.Total.Open.Critical)
	}

	if _, err := traust.GetReportFileAs[traust.FindingsRollupInsight](context.Background(), core.Client(), badID); err == nil {
		t.Fatal("expected unknown field error, got nil")
	}

	var problem *traust.ProblemError
	if _, err := traust.GetReportFileAs[traust.FindingsRollupInsight](context.Background(), core.Client(), notReadyID); !errors.As(err, &problem) || problem.StatusCode != http.StatusConflict {
		t.Fatalf("want 409 ProblemError, got %v", err)
	}

	if _, err := traust.DecodeReportView[traust.FindingsRollupInsight]([]byte(`{"repositories":[]} {"extra":1}`)); err == nil {
		t.Fatal("expected trailing JSON error, got nil")
	}
}
