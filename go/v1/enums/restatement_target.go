// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// RestatementTarget — The metadata field being restated.
type RestatementTarget string

const (
	RestatementTargetClaimHashes       RestatementTarget = "claim_hashes"
	RestatementTargetAuditReportSha256 RestatementTarget = "audit_report_sha256"
	RestatementTargetArtifactDigests   RestatementTarget = "artifact_digests"
)

// RestatementTargetValues returns all valid RestatementTarget values.
func RestatementTargetValues() []RestatementTarget {
	return []RestatementTarget{
		RestatementTargetClaimHashes,
		RestatementTargetAuditReportSha256,
		RestatementTargetArtifactDigests,
	}
}
