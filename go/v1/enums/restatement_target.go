// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
