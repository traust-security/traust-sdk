// Package insight is the app's reporting: building a report, watching it finish, and its file.
package insight

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	coreinsight "github.com/traust-security/traust-sdk/go/v1/traust/insight"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// Handler serves the app's reports.
type Handler struct {
	coreInsight *coreinsight.Client
}

// New builds the feature's handler.
func New(core *traust.Client) *Handler { return &Handler{coreInsight: coreinsight.New(core)} }

// Routes registers the feature's pages.
func (h *Handler) Routes(mux *http.ServeMux) {
	mux.HandleFunc("POST /reports", web.SignedIn(h.startReport))
	mux.HandleFunc("GET /reports/{token}", web.SignedIn(h.reportStatus))
	mux.HandleFunc("GET /reports/{token}/file", web.SignedIn(h.reportFile))
}

func (h *Handler) startReport(w http.ResponseWriter, r *http.Request) {
	name, format := traust.ViewName(r.PostFormValue("report")), traust.ReportFormatMarkdown
	poller, err := h.coreInsight.BeginRequestReport(r.Context(), traust.ReportRequest{Report: name, Format: &format})
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, reportRow(poller.ResumeToken(), traust.Report{Report: name, Status: traust.ReportStatusQueued}, false))
}

func (h *Handler) reportStatus(w http.ResponseWriter, r *http.Request) {
	poller, err := h.coreInsight.ResumeReport(r.PathValue("token"))
	if err != nil {
		web.Render(w, r, http.StatusBadRequest, web.Failure("That is not a report."))
		return
	}
	report, done, err := poller.Poll(r.Context())
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, reportRow(poller.ResumeToken(), report, done))
}

func (h *Handler) reportFile(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "token")
	if !ok {
		return
	}
	file, err := h.coreInsight.GetReportFile(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="report.md"`)
	_, _ = w.Write(file)
}
