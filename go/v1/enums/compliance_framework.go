// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

package enums

// ComplianceFramework — Framework or baseline a compliance assessment is scoped to. fedramp-high and fedramp-moderate are 800-53B baselines assessed over the nist-800-53-rev5 control catalog (compliance_control_catalog).
type ComplianceFramework string

const (
	ComplianceFrameworkNist80053Rev5   ComplianceFramework = "nist-800-53-rev5"
	ComplianceFrameworkFedrampHigh     ComplianceFramework = "fedramp-high"
	ComplianceFrameworkFedrampModerate ComplianceFramework = "fedramp-moderate"
	ComplianceFrameworkPciDssV4        ComplianceFramework = "pci-dss-v4"
	ComplianceFrameworkSoc2Tsc         ComplianceFramework = "soc2-tsc"
	ComplianceFrameworkGdprTechnical   ComplianceFramework = "gdpr-technical"
)

// ComplianceFrameworkValues returns all valid ComplianceFramework values.
func ComplianceFrameworkValues() []ComplianceFramework {
	return []ComplianceFramework{
		ComplianceFrameworkNist80053Rev5,
		ComplianceFrameworkFedrampHigh,
		ComplianceFrameworkFedrampModerate,
		ComplianceFrameworkPciDssV4,
		ComplianceFrameworkSoc2Tsc,
		ComplianceFrameworkGdprTechnical,
	}
}
