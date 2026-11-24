// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

package enums

// FindingOrigin — Values of `origin` in report.schema.json.
type FindingOrigin string

const (
	FindingOriginVerifyRemediation   FindingOrigin = "verify-remediation"
	FindingOriginVulnScan            FindingOrigin = "vuln-scan"
	FindingOriginCreateFuzzing       FindingOrigin = "create-fuzzing"
	FindingOriginValidateFindings    FindingOrigin = "validate-findings"
	FindingOriginValidationDiscovery FindingOrigin = "validation-discovery"
	FindingOriginImpactAnalysis      FindingOrigin = "impact-analysis"
)

// FindingOriginValues returns all valid FindingOrigin values.
func FindingOriginValues() []FindingOrigin {
	return []FindingOrigin{
		FindingOriginVerifyRemediation,
		FindingOriginVulnScan,
		FindingOriginCreateFuzzing,
		FindingOriginValidateFindings,
		FindingOriginValidationDiscovery,
		FindingOriginImpactAnalysis,
	}
}
