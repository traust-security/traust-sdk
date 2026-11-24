// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

package types

type FindingIntake struct {
	Bindings        []FindingIntakeBindingsItem `json:"bindings"`
	Commit          string                      `json:"commit"`
	Repository      string                      `json:"repository"`
	RepositoryId    string                      `json:"repository_id"`
	RepositoryRefId string                      `json:"repository_ref_id"`
	SchemaVersion   interface{}                 `json:"schema_version"`
}

type FindingIntakeBindingsItem struct {
	Fingerprint     string `json:"fingerprint"`
	FingerprintAlgo string `json:"fingerprint_algo"`
	Label           string `json:"label"`
}
