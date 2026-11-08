// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// ComparisonStatus — Outcome of one forge observation of a repository (continuous-ops redesign step 1).
type ComparisonStatus string

const (
	ComparisonStatusOk              ComparisonStatus = "ok"
	ComparisonStatusQuotaDeferred   ComparisonStatus = "quota-deferred"
	ComparisonStatusNoPinnedSha     ComparisonStatus = "no-pinned-sha"
	ComparisonStatusUnsupportedHost ComparisonStatus = "unsupported-host"
	ComparisonStatusNoRepoUrl       ComparisonStatus = "no-repo-url"
	ComparisonStatusNoCredentials   ComparisonStatus = "no-credentials"
	ComparisonStatusUnreachable     ComparisonStatus = "unreachable"
	ComparisonStatusError           ComparisonStatus = "error"
	ComparisonStatusBanSuspected    ComparisonStatus = "ban-suspected"
	ComparisonStatusNetworkSkipped  ComparisonStatus = "network-skipped"
	ComparisonStatusNeverAudited    ComparisonStatus = "never-audited"
)

// ComparisonStatusValues returns all valid ComparisonStatus values.
func ComparisonStatusValues() []ComparisonStatus {
	return []ComparisonStatus{
		ComparisonStatusOk,
		ComparisonStatusQuotaDeferred,
		ComparisonStatusNoPinnedSha,
		ComparisonStatusUnsupportedHost,
		ComparisonStatusNoRepoUrl,
		ComparisonStatusNoCredentials,
		ComparisonStatusUnreachable,
		ComparisonStatusError,
		ComparisonStatusBanSuspected,
		ComparisonStatusNetworkSkipped,
		ComparisonStatusNeverAudited,
	}
}
