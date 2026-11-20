package main

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

func (a *app) startRun(w http.ResponseWriter, r *http.Request) {
	order := traust.OrderRequest{Asset: r.PostFormValue("asset"), Priority: 1}
	if repo := r.PostFormValue("repository_id"); repo != "" {
		id, err := uuid.Parse(repo)
		if err != nil {
			render(w, r, http.StatusBadRequest, failure("That is not a repository."))
			return
		}
		order.RepositoryId = &id
	}
	poller, err := a.core.BeginSubmit(r.Context(), traust.SubmitRequest{Order: order, Config: &[]traust.ConfigDigest{a.pack}})
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, runRow(poller.ResumeToken(), traust.RunOrderStateQueued, false))
}

// runStatus answers a run's row; a finished run also tells the page, so lists that depend on it
// (a repository's findings) refresh themselves.
func (a *app) runStatus(w http.ResponseWriter, r *http.Request) {
	poller, err := a.core.ResumeOrder(r.PathValue("token"))
	if err != nil {
		render(w, r, http.StatusBadRequest, failure("That is not a run."))
		return
	}
	order, done, err := poller.Poll(r.Context())
	if err != nil {
		fail(w, r, err)
		return
	}
	if done {
		w.Header().Set("HX-Trigger", "run-finished")
	}
	render(w, r, http.StatusOK, runRow(poller.ResumeToken(), order.State, done))
}

// runHistory is a repository's latest runs, from any requester.
func (a *app) runHistory(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "id")
	if !ok {
		return
	}
	limit := 10
	page, err := a.core.ListOrders(r.Context(), &traust.ListOrdersParams{RepositoryId: &id, Limit: &limit})
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, runHistory(page.Items))
}

func (a *app) startReport(w http.ResponseWriter, r *http.Request) {
	name, format := traust.ViewName(r.PostFormValue("report")), traust.ReportFormatMarkdown
	poller, err := a.core.BeginRequestReport(r.Context(), traust.ReportRequest{Report: name, Format: &format})
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, reportRow(poller.ResumeToken(), traust.Report{Report: name, Status: traust.ReportStatusQueued}, false))
}

func (a *app) reportStatus(w http.ResponseWriter, r *http.Request) {
	poller, err := a.core.ResumeReport(r.PathValue("token"))
	if err != nil {
		render(w, r, http.StatusBadRequest, failure("That is not a report."))
		return
	}
	report, done, err := poller.Poll(r.Context())
	if err != nil {
		fail(w, r, err)
		return
	}
	render(w, r, http.StatusOK, reportRow(poller.ResumeToken(), report, done))
}

func (a *app) reportFile(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r, "token")
	if !ok {
		return
	}
	file, err := a.core.GetReportFile(r.Context(), id)
	if err != nil {
		fail(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="report.md"`)
	_, _ = w.Write(file)
}
