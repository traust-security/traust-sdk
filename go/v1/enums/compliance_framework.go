// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

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
