// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// PackRequiredConfig — Config a pack asset can require: traust.yaml top-level keys (traust.schema.json) plus model-registry for model-registry.yaml.
type PackRequiredConfig string

const (
	PackRequiredConfigHome             PackRequiredConfig = "home"
	PackRequiredConfigDatabase         PackRequiredConfig = "database"
	PackRequiredConfigArtifacts        PackRequiredConfig = "artifacts"
	PackRequiredConfigProcess          PackRequiredConfig = "process"
	PackRequiredConfigHttp             PackRequiredConfig = "http"
	PackRequiredConfigLogging          PackRequiredConfig = "logging"
	PackRequiredConfigRating           PackRequiredConfig = "rating"
	PackRequiredConfigReview           PackRequiredConfig = "review"
	PackRequiredConfigLedger           PackRequiredConfig = "ledger"
	PackRequiredConfigIdentity         PackRequiredConfig = "identity"
	PackRequiredConfigRecording        PackRequiredConfig = "recording"
	PackRequiredConfigAllowedRepoHosts PackRequiredConfig = "allowed_repo_hosts"
	PackRequiredConfigForgeHosts       PackRequiredConfig = "forge_hosts"
	PackRequiredConfigForgeTokens      PackRequiredConfig = "forge_tokens"
	PackRequiredConfigBudgets          PackRequiredConfig = "budgets"
	PackRequiredConfigRemediation      PackRequiredConfig = "remediation"
	PackRequiredConfigCampaign         PackRequiredConfig = "campaign"
	PackRequiredConfigCompliance       PackRequiredConfig = "compliance"
	PackRequiredConfigFeeds            PackRequiredConfig = "feeds"
	PackRequiredConfigExecution        PackRequiredConfig = "execution"
	PackRequiredConfigLocations        PackRequiredConfig = "locations"
	PackRequiredConfigOverrides        PackRequiredConfig = "overrides"
	PackRequiredConfigModelRegistry    PackRequiredConfig = "model-registry"
)

// PackRequiredConfigValues returns all valid PackRequiredConfig values.
func PackRequiredConfigValues() []PackRequiredConfig {
	return []PackRequiredConfig{
		PackRequiredConfigHome,
		PackRequiredConfigDatabase,
		PackRequiredConfigArtifacts,
		PackRequiredConfigProcess,
		PackRequiredConfigHttp,
		PackRequiredConfigLogging,
		PackRequiredConfigRating,
		PackRequiredConfigReview,
		PackRequiredConfigLedger,
		PackRequiredConfigIdentity,
		PackRequiredConfigRecording,
		PackRequiredConfigAllowedRepoHosts,
		PackRequiredConfigForgeHosts,
		PackRequiredConfigForgeTokens,
		PackRequiredConfigBudgets,
		PackRequiredConfigRemediation,
		PackRequiredConfigCampaign,
		PackRequiredConfigCompliance,
		PackRequiredConfigFeeds,
		PackRequiredConfigExecution,
		PackRequiredConfigLocations,
		PackRequiredConfigOverrides,
		PackRequiredConfigModelRegistry,
	}
}
