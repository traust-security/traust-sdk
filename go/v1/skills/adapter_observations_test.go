package skills_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/skills"
	"github.com/traust-security/traust-sdk/go/v1/skills/skillstest"
)

func TestScanRejectsNativeObservationsAsAuthoredFindings(t *testing.T) {
	payload, err := json.Marshal(skillstest.FixtureReport())
	if err != nil {
		t.Fatal(err)
	}
	var report map[string]json.RawMessage
	if err := json.Unmarshal(payload, &report); err != nil {
		t.Fatal(err)
	}
	report["findings"] = json.RawMessage(`[{
		"id":"native:rule/observation#1", "title":"x", "severity":"high",
		"cwes":[], "locations":[{"path":"config/example.env","lines":"12"}],
		"description":"", "remediation":""
	}]`)
	payload, err = json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	_, err = skills.Scan.Run(context.Background(), rawProvider{payload: payload}, skills.ScanInput{
		Repo: "https://github.com/org/repo", Ref: "main",
	})
	var skillErr *skills.SkillError
	if !errors.As(err, &skillErr) || skillErr.Phase != skills.PhaseValidateOutput {
		t.Fatalf("native observation in authored report = %v, want output validation failure", err)
	}
}
