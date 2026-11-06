// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package types

type ValidationAuthorization struct {
	ApprovedAt        string      `json:"approved_at"`
	ApprovedBy        string      `json:"approved_by"`
	Artifact          interface{} `json:"artifact"`
	Mode              interface{} `json:"mode"`
	PermitDestructive bool        `json:"permit_destructive"`
	PlanDigest        string      `json:"plan_digest"`
	PlanId            string      `json:"plan_id"`
	Scope             Scope       `json:"scope"`
	SkipStepIds       []string    `json:"skip_step_ids"`
}

type Browser struct {
	AllowedVerbs    []string `json:"allowed_verbs"`
	AttackerOrigins []string `json:"attacker_origins"`
	VictimOrigin    string   `json:"victim_origin"`
}

type Cluster struct {
	Context            string   `json:"context"`
	ExplicitNamespaces []string `json:"explicit_namespaces"`
	Namespaces         []string `json:"namespaces"`
	VerbsDenied        []string `json:"verbs_denied"`
}

type OffLimit struct {
	Context   interface{} `json:"context,omitempty"`
	Image     interface{} `json:"image,omitempty"`
	Name      interface{} `json:"name,omitempty"`
	Namespace interface{} `json:"namespace,omitempty"`
	Resource  interface{} `json:"resource,omitempty"`
	Verb      interface{} `json:"verb,omitempty"`
}

type Scope struct {
	AuthorizedBy             interface{}   `json:"authorized_by"`
	Browser                  []Browser     `json:"browser"`
	Clusters                 []Cluster     `json:"clusters"`
	Containers               []string      `json:"containers"`
	CredentialProbeClasses   []string      `json:"credential_probe_classes"`
	Engagement               interface{}   `json:"engagement"`
	Environment              interface{}   `json:"environment"`
	Expires                  interface{}   `json:"expires"`
	Images                   []string      `json:"images"`
	InventoryAwsProfiles     []string      `json:"inventory_aws_profiles"`
	InventoryClusterContexts []string      `json:"inventory_cluster_contexts"`
	Modes                    []interface{} `json:"modes"`
	OffLimits                []OffLimit    `json:"off_limits"`
	WasmArtifactIds          []string      `json:"wasm_artifact_ids"`
}
