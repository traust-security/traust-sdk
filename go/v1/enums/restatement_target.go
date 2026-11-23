// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
