// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
