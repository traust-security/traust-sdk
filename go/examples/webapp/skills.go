package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
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

// fixtureSkills answers the scan skill with canned findings about the repository it is given,
// so the app is playable without an agent or a model key; every other asset runs its pack
// command. A live deployment uses ops.ProcessExecutor (or its own) and the pack's real skill.
type fixtureSkills struct {
	commands ops.ProcessExecutor
	now      func() time.Time
}

func (f fixtureSkills) Execute(ctx context.Context, execution ops.Execution) (ops.ProducedReport, error) {
	if execution.Order.Asset != "scan" {
		return f.commands.Execute(ctx, execution)
	}
	var subject traust.SubjectRef
	if err := json.Unmarshal(execution.Input, &subject); err != nil || len(subject.Sha) < 7 {
		return ops.ProducedReport{}, fmt.Errorf("fixture scan needs the prepared subject: %v", err)
	}
	payload, err := fixtureReport(subject, f.now())
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

func jsonString(s string) (string, error) {
	quoted, err := json.Marshal(s)
	return string(quoted), err
}
