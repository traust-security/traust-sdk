package traust

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
)

// fakeReports serves POST /v1/reports (202, the build queued) and GET /v1/reports/{id}; the
// build ends as status on the second read.
func fakeReports(t *testing.T, status ReportStatus) *Client {
	t.Helper()
	id := uuid.New()
	var reads atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		report := Report{OrderId: id, Report: FindingsRollup, Format: ReportFormatMarkdown, Status: ReportStatusQueued}
		code := http.StatusOK
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v1/reports":
			code = http.StatusAccepted
		case r.Method == http.MethodGet && r.URL.Path == "/v1/reports/"+id.String():
			if reads.Add(1) >= 2 {
				report.Status = status
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(report)
	}))
	t.Cleanup(server.Close)
	client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("t")})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestBeginRequestReportFollowsTheBuildToItsEnd(t *testing.T) {
	for _, status := range []ReportStatus{ReportStatusReady, ReportStatusFailed} {
		t.Run(string(status), func(t *testing.T) {
			poller, err := fakeReports(t, status).BeginRequestReport(t.Context(), ReportRequest{Report: FindingsRollup})
			if err != nil {
				t.Fatal(err)
			}
			report, err := poller.Wait(t.Context(), time.Millisecond)
			if err != nil || report.Status != status || report.OrderId.String() != poller.ResumeToken() {
				t.Fatalf("report %+v err %v, want the requested build %s", report, err, status)
			}
		})
	}
}

func TestResumeReportRefusesATokenThatNamesNoBuild(t *testing.T) {
	if _, err := fakeReports(t, ReportStatusReady).ResumeReport("not-a-build"); err == nil {
		t.Fatal("ResumeReport accepted a token that names no build")
	}
}
