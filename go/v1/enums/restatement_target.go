// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
