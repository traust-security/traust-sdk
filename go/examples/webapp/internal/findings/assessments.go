package findings

import (
	"context"
	"errors"
	"maps"
	"mime"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/findings"
)

// scanKind is the assessment kind this app's scans record, so the section marks the latest one
// core recorded — the same assessment a triage judges.
const scanKind = traust.AssessmentKindSecureCode1

// assessments answers the repository page's assessment section: the repository's assessments,
// the latest scan core recorded marked.
func (h *Handler) assessments(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	listed, err := h.coreFindings.ListRepositoryAssessments(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	var latest uuid.UUID
	if recorded, err := h.coreFindings.GetLatestRecordedAssessment(r.Context(), id, &traust.GetLatestRecordedAssessmentParams{Kind: scanKind}); err == nil {
		if recorded.Id != nil {
			latest = *recorded.Id
		}
	} else if !errors.Is(err, traust.ErrNotFound) { // nothing recorded yet: nothing to mark
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, assessmentList(listed, latest))
}

// assessment is one assessment's page: what core recorded it as, the findings its intake
// created, and the evidence it left.
func (h *Handler) assessment(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.coreFindings.GetAssessment(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	intake, err := h.coreFindings.GetAssessmentIntake(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	artifacts, err := h.coreFindings.ListAssessmentArtifacts(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, assessment(*found, *intake, artifacts))
}

// allAssessments is every assessment core recorded for the person, across their repositories.
func (h *Handler) allAssessments(w http.ResponseWriter, r *http.Request) {
	var ids []uuid.UUID
	for repo, err := range h.coreInventory.ListRepositoriesAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		if repo.Id != nil {
			ids = append(ids, *repo.Id)
		}
	}
	var listed []traust.Assessment
	if len(ids) > 0 { // core refuses a set with no repositories, so an empty scope asks nothing
		for a, err := range h.coreFindings.ListAssessmentsAll(r.Context(), &traust.ListAssessmentsParams{RepositoryIds: &ids}) {
			if err != nil {
				web.Fail(w, r, err)
				return
			}
			listed = append(listed, a)
		}
	}
	web.Render(w, r, http.StatusOK, allAssessments(listed))
}

// artifact is one artifact's page: the evidence's metadata, and — for a report — the
// document core kept, decoded as the contract type it publishes. A report that cannot be read
// says so on the page instead of hiding the evidence already in hand.
func (h *Handler) artifact(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.coreFindings.GetArtifact(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	report, reportErr := h.report(r.Context(), id, found.Kind)
	web.Render(w, r, http.StatusOK, artifact(*found, report, reportErr))
}

// report is the artifact's report document when it is one, so the page reads the assessment's
// own words instead of bytes; every other kind stays a download.
func (h *Handler) report(ctx context.Context, id uuid.UUID, kind traust.ArtifactKind) (*traust.ReportArtifact, error) {
	if kind != traust.ArtifactKindReport {
		return nil, nil
	}
	obj, err := h.coreFindings.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return findings.Decode[traust.ReportArtifact](obj)
}

// artifactContent streams an artifact's bytes under the name and type its content earned it.
func (h *Handler) artifactContent(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	obj, err := h.coreFindings.Get(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	w.Header().Set("Content-Type", obj.MediaType)
	w.Header().Set("Content-Disposition", `attachment; filename="`+artifactFile(obj)+`"`)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(obj.Body)
}

// assessmentKind words an assessment's kind; one core worded unexpectedly says so.
func assessmentKind(a traust.Assessment) string {
	if kind, err := a.Kind.AsAssessmentKind(); err == nil {
		return string(kind)
	}
	return "unknown kind"
}

// intakeLabels is an intake's report labels, ordered so the page reads the same every time.
func intakeLabels(intake traust.AssessmentIntakeBindings) []string {
	return slices.Sorted(maps.Keys(intake.Findings))
}

// artifactFile names an artifact's download after what core recorded of it — its kind, its
// digest, the extension its media type names — keeping only the runes a filename may carry.
func artifactFile(obj *findings.Object) string {
	extension := ".bin"
	if known, err := mime.ExtensionsByType(obj.MediaType); err == nil && len(known) > 0 {
		extension = known[0]
	}
	name := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-', r == '_':
			return r
		}
		return -1
	}, string(obj.Kind)+"-"+web.Short(&obj.Digest)+extension)
	if name != "" {
		return name
	}
	return "artifact"
}
