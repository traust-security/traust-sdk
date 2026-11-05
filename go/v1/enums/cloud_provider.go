// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// CloudProvider — Values of `provider` in cloud-config-audit.schema.json.
type CloudProvider string

const (
	CloudProviderAws        CloudProvider = "aws"
	CloudProviderAzure      CloudProvider = "azure"
	CloudProviderGcp        CloudProvider = "gcp"
	CloudProviderKubernetes CloudProvider = "kubernetes"
	CloudProviderDocker     CloudProvider = "docker"
	CloudProviderOther      CloudProvider = "other"
)

// CloudProviderValues returns all valid CloudProvider values.
func CloudProviderValues() []CloudProvider {
	return []CloudProvider{
		CloudProviderAws,
		CloudProviderAzure,
		CloudProviderGcp,
		CloudProviderKubernetes,
		CloudProviderDocker,
		CloudProviderOther,
	}
}
