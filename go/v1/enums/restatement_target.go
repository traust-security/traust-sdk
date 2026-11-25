// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
