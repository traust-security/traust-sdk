package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/enums"
	"github.com/traust-security/traust-sdk/go/v1/types"
)

// Fixtures preserve the sparse/historical inputs from contracts
// 2e521561 tests/test_adapter_observations.py, without filling missing evidence.
func TestAdapterObservationsRoundTripAndStorage(t *testing.T) {
	eachDialect(t, func(t *testing.T, client *Client) {
		ctx := context.Background()
		for _, name := range []string{"sparse", "rich"} {
			t.Run(name, func(t *testing.T) {
				payload := readFixture(t, "testdata/adapter-result-"+name+".test.json")
				artifact := mustParseArtifact(t, payload, types.ParseAdapterResultArtifact)
				value, err := artifact.Value()
				if err != nil {
					t.Fatal(err)
				}
				encoded, err := types.EncodeAdapterResultArtifact(value)
				if err != nil {
					t.Fatal(err)
				}
				assertProjectionJSON(t, encoded.Payload(), payload)
				for mode, candidate := range map[string]types.Artifact[types.AdapterResult]{
					"parsed": artifact, "encoded": encoded,
				} {
					reference := resolverOf(client).Write("s3://results/adapter/"+name+"/"+mode+".json", candidate.Payload())
					input := SaveAdapterResultInput{
						Binding: runBinding(name+"-"+mode, nil), Artifact: candidate,
						References: []string{reference},
					}
					result, err := client.SaveAdapterResult(ctx, input)
					if err != nil {
						t.Fatal(err)
					}
					stored, err := client.GetAdapterResult(ctx, result.BindingID)
					if err != nil {
						t.Fatal(err)
					}
					if !bytes.Equal(stored.Payload(), candidate.Payload()) {
						t.Fatal("stored adapter payload lost exact bytes")
					}
					storedValue, err := stored.Value()
					if err != nil || !reflect.DeepEqual(storedValue, value) {
						t.Fatalf("stored typed observation = %+v, %v; want %+v", storedValue, err, value)
					}
					prefix, placeholder := "", "?"
					if client.store.dialect == dialectPostgres {
						prefix, placeholder = "traust_storage.", "$1"
					}
					var findings, metadata, digest string
					if err := sqlDB(client).QueryRowContext(ctx,
						"SELECT findings, metadata, artifact_digest FROM "+prefix+"adapter_result WHERE binding_id = "+placeholder,
						result.BindingID,
					).Scan(&findings, &metadata, &digest); err != nil {
						t.Fatal(err)
					}
					var original map[string]json.RawMessage
					if err := json.Unmarshal(payload, &original); err != nil {
						t.Fatal(err)
					}
					assertProjectionJSON(t, []byte(findings), original["findings"])
					assertProjectionJSON(t, []byte(metadata), original["metadata"])
					wantDigest, _ := identifyArtifact(candidate.Payload())
					binding, err := client.GetBinding(ctx, result.BindingID)
					if err != nil || binding.Digest != wantDigest || digest != wantDigest ||
						binding.ByteSize != int64(len(candidate.Payload())) ||
						!reflect.DeepEqual(binding.References, []string{reference}) {
						t.Fatalf("adapter evidence = %+v, projection digest %q, %v", binding, digest, err)
					}
					again, err := client.SaveAdapterResult(ctx, input)
					if err != nil || !again.AlreadyBound || again.BindingID != result.BindingID {
						t.Fatalf("adapter retry = %+v, %v", again, err)
					}
				}
			})
		}
	})
}

func assertProjectionJSON(t *testing.T, got, want []byte) {
	t.Helper()
	var actual, expected any
	if err := json.Unmarshal(got, &actual); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(want, &expected); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("adapter JSON changed:\ngot  %s\nwant %s", got, want)
	}
}

func TestAdapterObservationsRejectMalformedPayloads(t *testing.T) {
	payload := readFixture(t, "testdata/adapter-result-sparse.test.json")
	for _, tc := range []struct {
		name, field string
		value       any
	}{
		{"empty-id", "id", ""},
		{"empty-title", "title", ""},
		{"severity", "severity", "urgent"},
		{"cwe-format", "cwes", []any{"79"}},
		{"cwe-type", "cwes", []any{79}},
		{"null-description", "description", nil},
		{"null-remediation", "remediation", nil},
		{"empty-locations", "locations", []any{}},
		{"empty-path", "locations", []any{map[string]any{"path": ""}}},
		{"line-type", "locations", []any{map[string]any{"path": "a", "lines": 12}}},
		{"validation-status", "validation_status", "verified"},
		{"origin", "origin", "scanner"},
		{"effort", "remediation_effort", "unknown"},
		{"blocking-type", "blocked_external", "false"},
		{"cvss", "cvss", map[string]any{"score": 11, "vector": "native"}},
		{"dependency", "dependency", map[string]any{"advisory": "GHSA-example"}},
		{"disposition", "disposition", map[string]any{"validity": "confirmed"}},
		{"empty-peach", "peach_references", []any{}},
		{"empty-llm", "owasp_llm_references", []any{}},
		{"empty-isolation", "isolation_dimensions", []any{}},
		{"unknown-field", "unexpected", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			invalid := replaceJSONField(t, payload, func(document map[string]any) {
				document["findings"].([]any)[0].(map[string]any)[tc.field] = tc.value
			})
			if _, err := types.ParseAdapterResultArtifact(invalid); err == nil {
				t.Fatal("malformed adapter observation passed the Artifact parse gate")
			}
		})
	}
	for _, field := range []string{"id", "title", "severity", "cwes", "locations", "description", "remediation"} {
		t.Run("missing-"+field, func(t *testing.T) {
			invalid := replaceJSONField(t, payload, func(document map[string]any) {
				delete(document["findings"].([]any)[0].(map[string]any), field)
			})
			if _, err := types.ParseAdapterResultArtifact(invalid); err == nil {
				t.Fatal("missing required observation field passed the Artifact parse gate")
			}
		})
	}
}

func TestAdapterObservationsRejectMalformedEncodedValues(t *testing.T) {
	payload := readFixture(t, "testdata/adapter-result-rich.test.json")
	for name, mutate := range map[string]func(*types.AdapterFinding){
		"severity": func(f *types.AdapterFinding) { f.Severity = enums.Severity("urgent") },
		"cwe":      func(f *types.AdapterFinding) { f.Cwes = []string{"79"} },
		"location": func(f *types.AdapterFinding) { f.Locations[0].Path = "" },
		"cvss":     func(f *types.AdapterFinding) { f.Cvss.Score = 11 },
		"validity": func(f *types.AdapterFinding) { f.Disposition.Validity = enums.Validity("unknown") },
	} {
		t.Run(name, func(t *testing.T) {
			artifact := mustParseArtifact(t, payload, types.ParseAdapterResultArtifact)
			value, err := artifact.Value()
			if err != nil {
				t.Fatal(err)
			}
			mutate(&value.Findings[0])
			if _, err := types.EncodeAdapterResultArtifact(value); err == nil {
				t.Fatal("malformed adapter observation passed the Artifact encode gate")
			}
		})
	}
}

func TestAdapterStorageRevalidatesArtifactSchema(t *testing.T) {
	eachDialect(t, func(t *testing.T, client *Client) {
		artifact := mustParseArtifact(t, sampleArtifacts(t)["report"], func(payload []byte) (types.Artifact[types.AdapterResult], error) {
			return types.ParseArtifact[types.AdapterResult]("report", payload)
		})
		_, err := client.SaveAdapterResult(context.Background(), SaveAdapterResultInput{
			Binding: runBinding("wrong-schema", nil), Artifact: artifact,
		})
		var storageErr *Error
		if !errors.As(err, &storageErr) || storageErr.Phase != PhaseValidate {
			t.Fatalf("wrong-schema adapter save = %v, want validation failure", err)
		}
		prefix := ""
		if client.store.dialect == dialectPostgres {
			prefix = "traust_storage."
		}
		for _, table := range []string{"artifact_evidence", "artifact_binding", "adapter_result"} {
			var count int
			if err := sqlDB(client).QueryRow("SELECT count(*) FROM " + prefix + table).Scan(&count); err != nil || count != 0 {
				t.Fatalf("rejected adapter save wrote %s: count = %d, %v", table, count, err)
			}
		}
	})
}
