// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
