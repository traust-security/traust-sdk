// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
