package main

import (
	"encoding/json"
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
