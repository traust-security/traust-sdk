// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// RemediationStatus — Values of `rem_status` in remediation.schema.json.
type RemediationStatus string

const (
	RemediationStatusCandidate                  RemediationStatus = "candidate"
	RemediationStatusForked                     RemediationStatus = "forked"
	RemediationStatusPatched                    RemediationStatus = "patched"
	RemediationStatusChecksPassed               RemediationStatus = "checks_passed"
	RemediationStatusChecksFailed               RemediationStatus = "checks_failed"
	RemediationStatusRevalidatedFixed           RemediationStatus = "revalidated_fixed"
	RemediationStatusRevalidatedStillVulnerable RemediationStatus = "revalidated_still_vulnerable"
	RemediationStatusPrOpened                   RemediationStatus = "pr_opened"
	RemediationStatusMerged                     RemediationStatus = "merged"
	RemediationStatusAbandoned                  RemediationStatus = "abandoned"
)

// RemediationStatusValues returns all valid RemediationStatus values.
func RemediationStatusValues() []RemediationStatus {
	return []RemediationStatus{
		RemediationStatusCandidate,
		RemediationStatusForked,
		RemediationStatusPatched,
		RemediationStatusChecksPassed,
		RemediationStatusChecksFailed,
		RemediationStatusRevalidatedFixed,
		RemediationStatusRevalidatedStillVulnerable,
		RemediationStatusPrOpened,
		RemediationStatusMerged,
		RemediationStatusAbandoned,
	}
}
