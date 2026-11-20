package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

//go:embed fixture-report.json.tmpl
var fixtureTemplate string

var (
	report = template.Must(template.New("report").Funcs(template.FuncMap{"json": jsonString}).Parse(fixtureTemplate))
	unslug = regexp.MustCompile(`[^A-Z0-9]`)
)

// fixtureSkills answers the scan skill with canned findings about the repository it is given and
// the triage skill with a verdict on each finding core prepared, so the app is playable without an
// agent or a model key; every other asset runs its pack command. A live deployment uses
// ops.ProcessExecutor (or its own) and the pack's real skills.
type fixtureSkills struct {
	commands ops.ProcessExecutor
	now      func() time.Time
}

func (f fixtureSkills) Execute(ctx context.Context, execution ops.Execution) (ops.ProducedReport, error) {
	var payload []byte
	var err error
	switch execution.Order.Asset {
	case "scan":
		var subject traust.SubjectRef
		if err := json.Unmarshal(execution.Input, &subject); err != nil || len(subject.Sha) < 7 {
			return ops.ProducedReport{}, fmt.Errorf("fixture scan needs the prepared subject: %v", err)
		}
		payload, err = fixtureReport(subject, f.now())
	case "triage":
		if execution.TriageInput == nil {
			return ops.ProducedReport{}, errors.New("fixture triage needs the prepared findings")
		}
		payload, err = fixtureTriage(*execution.TriageInput, f.now())
	default:
		return f.commands.Execute(ctx, execution)
	}
	if err != nil {
		return ops.ProducedReport{}, err
	}
	return ops.ProducedReport{Payload: payload, Result: traust.JobResult{Outcome: traust.OutcomeSucceeded}}, nil
}

// fixtureReport is a contracts-valid report whose finding ids follow the canonical
// {REPO_SLUG}-{SHORTSHA}-{NNN}, so a rescan of the same commit restates the same findings.
func fixtureReport(subject traust.SubjectRef, now time.Time) ([]byte, error) {
	name := path.Base(strings.TrimSuffix(subject.RepositoryUrl, ".git"))
	slug := unslug.ReplaceAllString(strings.ToUpper(name), "_")
	var out bytes.Buffer
	err := report.Execute(&out, map[string]string{
		"Name": name, "Repo": subject.RepositoryUrl, "Sha": subject.Sha, "Short": strings.ToLower(subject.Sha[:7]),
		"Slug": slug[:min(len(slug), 24)], "Date": now.UTC().Format(time.DateOnly),
	})
	return out.Bytes(), err
}

// verdicts cycle so a fixture triage shows each answer a person then countersigns or overrides
// (core takes hardening as a judgement on a true positive, never as a new verdict).
var verdicts = []string{"true_positive", "false_positive", "undetermined"}

type triageFinding struct {
	ID              string  `json:"id"`
	Title           string  `json:"title"`
	Verdict         string  `json:"verdict"`
	Fingerprint     string  `json:"fingerprint"`
	FingerprintAlgo string  `json:"fingerprint_algo"`
	Confidence      float64 `json:"confidence"`
	Rationale       string  `json:"rationale"`
}

// fixtureTriage is a contracts-valid triage judging exactly the findings core prepared; core
// matches each verdict to its finding by fingerprint.
func fixtureTriage(input traust.TriageInput, now time.Time) ([]byte, error) {
	counts := map[string]int{}
	findings := make([]triageFinding, 0, len(input.Items))
	for i, item := range input.Items {
		verdict := verdicts[i%len(verdicts)]
		counts[verdict]++
		findings = append(findings, triageFinding{
			ID: fmt.Sprintf("f%03d", i+1), Title: item.Finding.Title, Verdict: verdict,
			Fingerprint: item.Finding.Fingerprint, FingerprintAlgo: item.Finding.FingerprintAlgo,
			Confidence: 7, Rationale: "A fixture verdict, for a person to countersign or override.",
		})
	}
	return json.Marshal(map[string]any{
		"triage_completed": now.UTC().Format(time.DateOnly),
		"triage_context": map[string]any{
			"environment": "fixture triage", "votes_per_finding": 1, "repo": input.Subject.RepositoryUrl, "harness_version": "0.0.0",
		},
		"summary": map[string]any{
			"input_count": len(findings), "true_positives": counts["true_positive"], "hardening": 0,
			"false_positives": counts["false_positive"], "undetermined": counts["undetermined"], "duplicates": 0,
			"by_severity": map[string]int{"critical": 0, "high": 0, "medium": 0, "low": 0},
		},
		"findings": findings,
	})
}

func jsonString(s string) (string, error) {
	quoted, err := json.Marshal(s)
	return string(quoted), err
}
