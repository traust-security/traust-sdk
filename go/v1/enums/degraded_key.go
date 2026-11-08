// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// DegradedKey — Common degraded coverage causes and audit counters. Registered feed and insight-view identifiers remain owned by their registries; Core validates those subjects without collapsing their counts into generic causes.
type DegradedKey string

const (
	DegradedKeyQuotaDeferred     DegradedKey = "quota-deferred"
	DegradedKeyUnsupportedHost   DegradedKey = "unsupported-host"
	DegradedKeyNoRepoUrl         DegradedKey = "no-repo-url"
	DegradedKeyNoCredentials     DegradedKey = "no-credentials"
	DegradedKeyUnreachable       DegradedKey = "unreachable"
	DegradedKeyError             DegradedKey = "error"
	DegradedKeyBanSuspected      DegradedKey = "ban-suspected"
	DegradedKeyNetworkSkipped    DegradedKey = "network-skipped"
	DegradedKeyDbMissing         DegradedKey = "db-missing"
	DegradedKeyDbStale           DegradedKey = "db-stale"
	DegradedKeyDbVersionMismatch DegradedKey = "db-version-mismatch"
	DegradedKeyFailed            DegradedKey = "failed"
	DegradedKeyUnverified        DegradedKey = "unverified"
)

// DegradedKeyValues returns all valid DegradedKey values.
func DegradedKeyValues() []DegradedKey {
	return []DegradedKey{
		DegradedKeyQuotaDeferred,
		DegradedKeyUnsupportedHost,
		DegradedKeyNoRepoUrl,
		DegradedKeyNoCredentials,
		DegradedKeyUnreachable,
		DegradedKeyError,
		DegradedKeyBanSuspected,
		DegradedKeyNetworkSkipped,
		DegradedKeyDbMissing,
		DegradedKeyDbStale,
		DegradedKeyDbVersionMismatch,
		DegradedKeyFailed,
		DegradedKeyUnverified,
	}
}
