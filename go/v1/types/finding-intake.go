// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
