// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
