// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package enums

// AssessmentKind — Assessment.kind: the producing assessment, also the classifier source kind. Rebaseline and restatement are operations, not kinds of evidence.
type AssessmentKind string

const (
	AssessmentKindSecureCodeCanonical      AssessmentKind = "secure_code"
	AssessmentKindContainer                AssessmentKind = "container"
	AssessmentKindRpm                      AssessmentKind = "rpm"
	AssessmentKindCloudConfigCanonical     AssessmentKind = "cloud_config"
	AssessmentKindVulnScanCanonical        AssessmentKind = "vuln_scan"
	AssessmentKindScanner                  AssessmentKind = "scanner"
	AssessmentKindToolScan                 AssessmentKind = "tool_scan"
	AssessmentKindManifestSbom             AssessmentKind = "manifest_sbom"
	AssessmentKindThreatModelCanonical     AssessmentKind = "threat_model"
	AssessmentKindIsolationReviewCanonical AssessmentKind = "isolation_review"
	AssessmentKindPrivProfileCanonical     AssessmentKind = "priv_profile"
	AssessmentKindPqcReadinessCanonical    AssessmentKind = "pqc_readiness"
	AssessmentKindFuzz                     AssessmentKind = "fuzz"
	AssessmentKindRepoStateCanonical       AssessmentKind = "repo_state"
	AssessmentKindAudit                    AssessmentKind = "audit"
	AssessmentKindTriage                   AssessmentKind = "triage"
	AssessmentKindValidation               AssessmentKind = "validation"
	AssessmentKindVerification             AssessmentKind = "verification"
	AssessmentKindImpact                   AssessmentKind = "impact"
	AssessmentKindCompliance               AssessmentKind = "compliance"
	AssessmentKindBenchmarkAudit           AssessmentKind = "benchmark_audit"
	AssessmentKindRemediation              AssessmentKind = "remediation"
	AssessmentKindFleetFix                 AssessmentKind = "fleet_fix"
	AssessmentKindSecureCode               AssessmentKind = "secure-code"
	AssessmentKindCloudConfig              AssessmentKind = "cloud-config"
	AssessmentKindVulnScan                 AssessmentKind = "vuln-scan"
	AssessmentKindThreatModel              AssessmentKind = "threat-model"
	AssessmentKindIsolationReview          AssessmentKind = "isolation-review"
	AssessmentKindPrivProfile              AssessmentKind = "priv-profile"
	AssessmentKindPqcReadiness             AssessmentKind = "pqc-readiness"
	AssessmentKindRepoState                AssessmentKind = "repo-state"
)

// AssessmentKindValues returns all valid AssessmentKind values.
func AssessmentKindValues() []AssessmentKind {
	return []AssessmentKind{
		AssessmentKindSecureCodeCanonical,
		AssessmentKindContainer,
		AssessmentKindRpm,
		AssessmentKindCloudConfigCanonical,
		AssessmentKindVulnScanCanonical,
		AssessmentKindScanner,
		AssessmentKindToolScan,
		AssessmentKindManifestSbom,
		AssessmentKindThreatModelCanonical,
		AssessmentKindIsolationReviewCanonical,
		AssessmentKindPrivProfileCanonical,
		AssessmentKindPqcReadinessCanonical,
		AssessmentKindFuzz,
		AssessmentKindRepoStateCanonical,
		AssessmentKindAudit,
		AssessmentKindTriage,
		AssessmentKindValidation,
		AssessmentKindVerification,
		AssessmentKindImpact,
		AssessmentKindCompliance,
		AssessmentKindBenchmarkAudit,
		AssessmentKindRemediation,
		AssessmentKindFleetFix,
		AssessmentKindSecureCode,
		AssessmentKindCloudConfig,
		AssessmentKindVulnScan,
		AssessmentKindThreatModel,
		AssessmentKindIsolationReview,
		AssessmentKindPrivProfile,
		AssessmentKindPqcReadiness,
		AssessmentKindRepoState,
	}
}
