// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
