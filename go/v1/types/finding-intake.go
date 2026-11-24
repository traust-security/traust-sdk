// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
