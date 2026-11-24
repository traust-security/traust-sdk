// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
