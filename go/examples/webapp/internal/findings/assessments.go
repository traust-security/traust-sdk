package findings

import (
	"errors"
	"maps"
	"mime"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/traust"
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
	listed, err := h.core.ListRepositoryAssessments(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	var latest uuid.UUID
	if recorded, err := h.core.GetLatestRecordedAssessment(r.Context(), id, &traust.GetLatestRecordedAssessmentParams{Kind: scanKind}); err == nil {
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
	found, err := h.core.GetAssessment(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	intake, err := h.core.GetAssessmentIntake(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	artifacts, err := h.core.ListAssessmentArtifacts(r.Context(), id, nil)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, assessment(*found, *intake, artifacts))
}

// allAssessments is every assessment core recorded for the person, across their repositories.
func (h *Handler) allAssessments(w http.ResponseWriter, r *http.Request) {
	var listed []traust.Assessment
	for a, err := range h.core.ListAssessmentsAll(r.Context(), nil) {
		if err != nil {
			web.Fail(w, r, err)
			return
		}
		listed = append(listed, a)
	}
	web.Render(w, r, http.StatusOK, allAssessments(listed))
}

// artifact is one artifact's page: the evidence's metadata, and its bytes as a download.
func (h *Handler) artifact(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.core.GetArtifact(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	web.Render(w, r, http.StatusOK, artifact(*found))
}

// artifactContent streams an artifact's bytes under the name and type its metadata earned it.
func (h *Handler) artifactContent(w http.ResponseWriter, r *http.Request) {
	id, ok := web.PathID(w, r, "id")
	if !ok {
		return
	}
	found, err := h.core.GetArtifact(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	content, err := h.core.GetArtifactContent(r.Context(), id)
	if err != nil {
		web.Fail(w, r, err)
		return
	}
	w.Header().Set("Content-Type", artifactMediaType(*found))
	w.Header().Set("Content-Disposition", `attachment; filename="`+artifactFile(*found)+`"`)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(content)
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
// digest, the extension its uri names — keeping only the runes a filename may carry.
func artifactFile(a traust.Artifact) string {
	name := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-', r == '_':
			return r
		}
		return -1
	}, a.Kind+"-"+web.Short(&a.Digest)+filepath.Ext(a.Uri))
	if name != "" {
		return name
	}
	if a.Id != nil {
		return a.Id.String()
	}
	return "artifact"
}

// artifactMediaType is the artifact's own, from the extension its uri names; anything core left
// untyped downloads as octet stream.
func artifactMediaType(a traust.Artifact) string {
	if known := mime.TypeByExtension(filepath.Ext(a.Uri)); known != "" {
		return known
	}
	return "application/octet-stream"
}
