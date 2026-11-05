// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// EntryPointKind — What kind of interface a threat model's entry point is. A coverage check maps each kind to the deterministic pre-scan that enumerates it; kinds no pre-scan enumerates are reported as uncovered, and an entry point with no kind as unexplained.
type EntryPointKind string

const (
	EntryPointKindHttpApi        EntryPointKind = "http-api"
	EntryPointKindGrpcApi        EntryPointKind = "grpc-api"
	EntryPointKindK8sController  EntryPointKind = "k8s-controller"
	EntryPointKindLlmIntegration EntryPointKind = "llm-integration"
	EntryPointKindEventConsumer  EntryPointKind = "event-consumer"
	EntryPointKindNetworkPeer    EntryPointKind = "network-peer"
	EntryPointKindCli            EntryPointKind = "cli"
	EntryPointKindFileInput      EntryPointKind = "file-input"
	EntryPointKindDeployConfig   EntryPointKind = "deploy-config"
	EntryPointKindCiSupplyChain  EntryPointKind = "ci-supply-chain"
	EntryPointKindMetricsDebug   EntryPointKind = "metrics-debug"
	EntryPointKindOther          EntryPointKind = "other"
)

// EntryPointKindValues returns all valid EntryPointKind values.
func EntryPointKindValues() []EntryPointKind {
	return []EntryPointKind{
		EntryPointKindHttpApi,
		EntryPointKindGrpcApi,
		EntryPointKindK8sController,
		EntryPointKindLlmIntegration,
		EntryPointKindEventConsumer,
		EntryPointKindNetworkPeer,
		EntryPointKindCli,
		EntryPointKindFileInput,
		EntryPointKindDeployConfig,
		EntryPointKindCiSupplyChain,
		EntryPointKindMetricsDebug,
		EntryPointKindOther,
	}
}
