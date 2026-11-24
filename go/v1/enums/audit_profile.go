// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// AuditProfile — Report profile discriminator: 'code' = secure-code-audit (source/manifest audit), 'rpm' = secure-rpm-audit (dist-git packaging audit), 'container' = secure-container-audit (registry image audit via skopeo/syft/grype).
type AuditProfile string

const (
	AuditProfileCode      AuditProfile = "code"
	AuditProfileRpm       AuditProfile = "rpm"
	AuditProfileContainer AuditProfile = "container"
)

// AuditProfileValues returns all valid AuditProfile values.
func AuditProfileValues() []AuditProfile {
	return []AuditProfile{
		AuditProfileCode,
		AuditProfileRpm,
		AuditProfileContainer,
	}
}
