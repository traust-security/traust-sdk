// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
