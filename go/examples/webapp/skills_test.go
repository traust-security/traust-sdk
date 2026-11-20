package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/validate"
)

func TestTheFixtureScanReportsAboutTheRepositoryItWasGiven(t *testing.T) {
	subject := traust.SubjectRef{RepositoryId: uuid.New(), RepositoryRefId: uuid.New(), RepositoryUrl: "https://forge.example/org/my-app.git", Sha: "ABCDEF0123456789"}
	input, _ := json.Marshal(subject)
	skills := fixtureSkills{now: func() time.Time { return time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC) }}

	produced, err := skills.Execute(t.Context(), ops.Execution{Order: traust.RunOrder{Asset: "scan"}, Input: input})
	if err != nil || produced.Result.Outcome != traust.OutcomeSucceeded {
		t.Fatalf("produced %+v err %v", produced.Result, err)
	}
	if err := validate.Validate("report", produced.Payload); err != nil {
		t.Fatalf("the fixture report breaks the contract: %v", err)
	}
	payload := string(produced.Payload)
	for _, want := range []string{`"MY_APP-abcdef0-001"`, `"repository": "https://forge.example/org/my-app.git"`, `"date": "2026-01-02"`} {
		if !strings.Contains(payload, want) {
			t.Fatalf("report lacks %s", want)
		}
	}
	if _, err := skills.Execute(t.Context(), ops.Execution{Order: traust.RunOrder{Asset: "scan"}}); err == nil {
		t.Fatal("a scan without its prepared subject ran")
	}
}

func TestTheFixtureTriageJudgesEveryPreparedFinding(t *testing.T) {
	input := traust.TriageInput{Subject: traust.SubjectRef{RepositoryUrl: "https://forge.example/org/app"}}
	for i := range 4 {
		input.Items = append(input.Items, traust.TriageItem{Finding: traust.Finding{
			Title: fmt.Sprint("finding ", i), Fingerprint: strings.Repeat(fmt.Sprint(i), 64), FingerprintAlgo: "v1",
		}})
	}
	produced, err := fixtureSkills{now: time.Now}.Execute(t.Context(), ops.Execution{Order: traust.RunOrder{Asset: "triage"}, TriageInput: &input})
	if err != nil {
		t.Fatal(err)
	}
	if err := validate.Validate("triage", produced.Payload); err != nil {
		t.Fatalf("the fixture triage breaks the contract: %v", err)
	}
	var doc struct {
		Findings []triageFinding `json:"findings"`
	}
	if err := json.Unmarshal(produced.Payload, &doc); err != nil || len(doc.Findings) != 4 ||
		doc.Findings[0].Verdict != "true_positive" || doc.Findings[1].Verdict != "false_positive" ||
		doc.Findings[3].Fingerprint != input.Items[3].Finding.Fingerprint {
		t.Fatalf("verdicts %+v %v, want one per finding, cycling, by fingerprint", doc.Findings, err)
	}
	if _, err := (fixtureSkills{now: time.Now}).Execute(t.Context(), ops.Execution{Order: traust.RunOrder{Asset: "triage"}}); err == nil {
		t.Fatal("a triage without its prepared findings ran")
	}
}
