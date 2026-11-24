// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
