package types_test

import (
	"encoding/json"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/types"
)

func TestSessionTaskClaimBinding(t *testing.T) {
	const order = "00000000-0000-4000-8000-000000000001"
	for _, tc := range []struct {
		name       string
		order      any
		generation any
		valid      bool
	}{
		{"legacy ordered task must be restaged", order, nil, false},
		{"zero is not a claim", order, 0, false},
		{"ordered task with claim", order, 2, true},
		{"manual task without claim", nil, nil, true},
		{"manual task with claim", nil, 2, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			task := types.SessionTask{
				SchemaVersion: "1", RunId: tc.order, LeaseGeneration: tc.generation,
				Skill: "dependency-watch", Mode: "below-range", Args: []string{},
				Models: map[string]string{"judge": "review-model"}, Input: []string{},
				Output: types.SessionTaskOutput{Path: "output/verdicts.json", Schema: "below-range-verdicts"},
			}
			payload, err := json.Marshal(task)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := types.ParseSessionTaskArtifact(payload); (err == nil) != tc.valid {
				t.Errorf("parse task: want valid=%t, got err=%v", tc.valid, err)
			}
			if _, err := types.EncodeSessionTaskArtifact(task); (err == nil) != tc.valid {
				t.Errorf("encode task: want valid=%t, got err=%v", tc.valid, err)
			}
		})
	}
}

func TestHistoricalSessionManifestRetainsTaskDigest(t *testing.T) {
	const digest = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	manifest, err := types.ParseRunManifestArtifact([]byte(`{"artifact":"run-manifest","sessions":[{
		"skill":"dependency-watch","mode":"below-range","executor_session":null,
		"models":{"judge":"review-model"},"task_sha256":"` + digest + `","inputs":[],
		"output":{"path":"output/verdicts.json","sha256":null,"artifact_id":null}
	}]}`))
	if err != nil {
		t.Fatal(err)
	}
	value, err := manifest.Value()
	if err != nil {
		t.Fatal(err)
	}
	if len(value.Sessions) != 1 || value.Sessions[0].TaskSha256 != digest || value.Sessions[0].LeaseGeneration != nil {
		t.Fatalf("expected original digest and no claim: %+v", value.Sessions)
	}
}
