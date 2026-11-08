// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

package types

import "encoding/json"

func ParseAdapterResultArtifact(payload []byte) (Artifact[AdapterResult], error) {
	return ParseArtifact[AdapterResult]("adapter-result", payload)
}

func EncodeAdapterResultArtifact(value AdapterResult) (Artifact[AdapterResult], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[AdapterResult]{}, err
	}
	return ParseAdapterResultArtifact(payload)
}

func ParseAdrRegistryArtifact(payload []byte) (Artifact[AdrRegistry], error) {
	return ParseArtifact[AdrRegistry]("adr-registry", payload)
}

func EncodeAdrRegistryArtifact(value AdrRegistry) (Artifact[AdrRegistry], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[AdrRegistry]{}, err
	}
	return ParseAdrRegistryArtifact(payload)
}

func ParseAttackMappingArtifact(payload []byte) (Artifact[AttackMapping], error) {
	return ParseArtifact[AttackMapping]("attack-mapping", payload)
}

func EncodeAttackMappingArtifact(value AttackMapping) (Artifact[AttackMapping], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[AttackMapping]{}, err
	}
	return ParseAttackMappingArtifact(payload)
}

func ParseBelowRangePacketsArtifact(payload []byte) (Artifact[BelowRangePackets], error) {
	return ParseArtifact[BelowRangePackets]("below-range-packets", payload)
}

func EncodeBelowRangePacketsArtifact(value BelowRangePackets) (Artifact[BelowRangePackets], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[BelowRangePackets]{}, err
	}
	return ParseBelowRangePacketsArtifact(payload)
}

func ParseBelowRangeVerdictsArtifact(payload []byte) (Artifact[BelowRangeVerdicts], error) {
	return ParseArtifact[BelowRangeVerdicts]("below-range-verdicts", payload)
}

func EncodeBelowRangeVerdictsArtifact(value BelowRangeVerdicts) (Artifact[BelowRangeVerdicts], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[BelowRangeVerdicts]{}, err
	}
	return ParseBelowRangeVerdictsArtifact(payload)
}

func ParseBenchmarkTargetArtifact(payload []byte) (Artifact[BenchmarkTarget], error) {
	return ParseArtifact[BenchmarkTarget]("benchmark-target", payload)
}

func EncodeBenchmarkTargetArtifact(value BenchmarkTarget) (Artifact[BenchmarkTarget], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[BenchmarkTarget]{}, err
	}
	return ParseBenchmarkTargetArtifact(payload)
}

func ParseCloudConfigAuditArtifact(payload []byte) (Artifact[CloudConfigAudit], error) {
	return ParseArtifact[CloudConfigAudit]("cloud-config-audit", payload)
}

func EncodeCloudConfigAuditArtifact(value CloudConfigAudit) (Artifact[CloudConfigAudit], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[CloudConfigAudit]{}, err
	}
	return ParseCloudConfigAuditArtifact(payload)
}

func ParseCloudConfigFindingsCurrentArtifact(payload []byte) (Artifact[CloudConfigFindingsCurrent], error) {
	return ParseArtifact[CloudConfigFindingsCurrent]("cloud-config-findings-current", payload)
}

func EncodeCloudConfigFindingsCurrentArtifact(value CloudConfigFindingsCurrent) (Artifact[CloudConfigFindingsCurrent], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[CloudConfigFindingsCurrent]{}, err
	}
	return ParseCloudConfigFindingsCurrentArtifact(payload)
}

func ParseCodeFactsArtifact(payload []byte) (Artifact[CodeFacts], error) {
	return ParseArtifact[CodeFacts]("code-facts", payload)
}

func EncodeCodeFactsArtifact(value CodeFacts) (Artifact[CodeFacts], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[CodeFacts]{}, err
	}
	return ParseCodeFactsArtifact(payload)
}

func ParseComplianceAssessmentArtifact(payload []byte) (Artifact[ComplianceAssessment], error) {
	return ParseArtifact[ComplianceAssessment]("compliance-assessment", payload)
}

func EncodeComplianceAssessmentArtifact(value ComplianceAssessment) (Artifact[ComplianceAssessment], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ComplianceAssessment]{}, err
	}
	return ParseComplianceAssessmentArtifact(payload)
}

func ParseComplianceMappingArtifact(payload []byte) (Artifact[ComplianceMapping], error) {
	return ParseArtifact[ComplianceMapping]("compliance-mapping", payload)
}

func EncodeComplianceMappingArtifact(value ComplianceMapping) (Artifact[ComplianceMapping], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ComplianceMapping]{}, err
	}
	return ParseComplianceMappingArtifact(payload)
}

func ParseComplianceScopeArtifact(payload []byte) (Artifact[ComplianceScope], error) {
	return ParseArtifact[ComplianceScope]("compliance-scope", payload)
}

func EncodeComplianceScopeArtifact(value ComplianceScope) (Artifact[ComplianceScope], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ComplianceScope]{}, err
	}
	return ParseComplianceScopeArtifact(payload)
}

func ParseCorpusRegistryArtifact(payload []byte) (Artifact[CorpusRegistry], error) {
	return ParseArtifact[CorpusRegistry]("corpus-registry", payload)
}

func EncodeCorpusRegistryArtifact(value CorpusRegistry) (Artifact[CorpusRegistry], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[CorpusRegistry]{}, err
	}
	return ParseCorpusRegistryArtifact(payload)
}

func ParseDocVarianceArtifact(payload []byte) (Artifact[DocVariance], error) {
	return ParseArtifact[DocVariance]("doc-variance", payload)
}

func EncodeDocVarianceArtifact(value DocVariance) (Artifact[DocVariance], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[DocVariance]{}, err
	}
	return ParseDocVarianceArtifact(payload)
}

func ParseFleetFixArtifact(payload []byte) (Artifact[FleetFix], error) {
	return ParseArtifact[FleetFix]("fleet-fix", payload)
}

func EncodeFleetFixArtifact(value FleetFix) (Artifact[FleetFix], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[FleetFix]{}, err
	}
	return ParseFleetFixArtifact(payload)
}

func ParseFuzzCampaignArtifact(payload []byte) (Artifact[FuzzCampaign], error) {
	return ParseArtifact[FuzzCampaign]("fuzz-campaign", payload)
}

func EncodeFuzzCampaignArtifact(value FuzzCampaign) (Artifact[FuzzCampaign], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[FuzzCampaign]{}, err
	}
	return ParseFuzzCampaignArtifact(payload)
}

func ParseFuzzReportArtifact(payload []byte) (Artifact[FuzzReport], error) {
	return ParseArtifact[FuzzReport]("fuzz-report", payload)
}

func EncodeFuzzReportArtifact(value FuzzReport) (Artifact[FuzzReport], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[FuzzReport]{}, err
	}
	return ParseFuzzReportArtifact(payload)
}

func ParseImpactAnalysisArtifact(payload []byte) (Artifact[ImpactAnalysis], error) {
	return ParseArtifact[ImpactAnalysis]("impact-analysis", payload)
}

func EncodeImpactAnalysisArtifact(value ImpactAnalysis) (Artifact[ImpactAnalysis], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ImpactAnalysis]{}, err
	}
	return ParseImpactAnalysisArtifact(payload)
}

func ParseIsolationReviewArtifact(payload []byte) (Artifact[IsolationReview], error) {
	return ParseArtifact[IsolationReview]("isolation-review", payload)
}

func EncodeIsolationReviewArtifact(value IsolationReview) (Artifact[IsolationReview], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[IsolationReview]{}, err
	}
	return ParseIsolationReviewArtifact(payload)
}

func ParseLayerArtifact(payload []byte) (Artifact[Layer], error) {
	return ParseArtifact[Layer]("layer", payload)
}

func EncodeLayerArtifact(value Layer) (Artifact[Layer], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[Layer]{}, err
	}
	return ParseLayerArtifact(payload)
}

func ParseManifestCoverageArtifact(payload []byte) (Artifact[ManifestCoverage], error) {
	return ParseArtifact[ManifestCoverage]("manifest-coverage", payload)
}

func EncodeManifestCoverageArtifact(value ManifestCoverage) (Artifact[ManifestCoverage], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ManifestCoverage]{}, err
	}
	return ParseManifestCoverageArtifact(payload)
}

func ParseOperatorPrivProfileArtifact(payload []byte) (Artifact[OperatorPrivProfile], error) {
	return ParseArtifact[OperatorPrivProfile]("operator-priv-profile", payload)
}

func EncodeOperatorPrivProfileArtifact(value OperatorPrivProfile) (Artifact[OperatorPrivProfile], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[OperatorPrivProfile]{}, err
	}
	return ParseOperatorPrivProfileArtifact(payload)
}

func ParseOrgParametersArtifact(payload []byte) (Artifact[OrgParameters], error) {
	return ParseArtifact[OrgParameters]("org-parameters", payload)
}

func EncodeOrgParametersArtifact(value OrgParameters) (Artifact[OrgParameters], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[OrgParameters]{}, err
	}
	return ParseOrgParametersArtifact(payload)
}

func ParsePqcBlockersArtifact(payload []byte) (Artifact[PqcBlockers], error) {
	return ParseArtifact[PqcBlockers]("pqc-blockers", payload)
}

func EncodePqcBlockersArtifact(value PqcBlockers) (Artifact[PqcBlockers], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[PqcBlockers]{}, err
	}
	return ParsePqcBlockersArtifact(payload)
}

func ParsePqcDecisionTreeArtifact(payload []byte) (Artifact[PqcDecisionTree], error) {
	return ParseArtifact[PqcDecisionTree]("pqc-decision-tree", payload)
}

func EncodePqcDecisionTreeArtifact(value PqcDecisionTree) (Artifact[PqcDecisionTree], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[PqcDecisionTree]{}, err
	}
	return ParsePqcDecisionTreeArtifact(payload)
}

func ParsePqcFactsArtifact(payload []byte) (Artifact[PqcFacts], error) {
	return ParseArtifact[PqcFacts]("pqc-facts", payload)
}

func EncodePqcFactsArtifact(value PqcFacts) (Artifact[PqcFacts], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[PqcFacts]{}, err
	}
	return ParsePqcFactsArtifact(payload)
}

func ParsePqcReadinessArtifact(payload []byte) (Artifact[PqcReadiness], error) {
	return ParseArtifact[PqcReadiness]("pqc-readiness", payload)
}

func EncodePqcReadinessArtifact(value PqcReadiness) (Artifact[PqcReadiness], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[PqcReadiness]{}, err
	}
	return ParsePqcReadinessArtifact(payload)
}

func ParsePqcVendorClaimsArtifact(payload []byte) (Artifact[PqcVendorClaims], error) {
	return ParseArtifact[PqcVendorClaims]("pqc-vendor-claims", payload)
}

func EncodePqcVendorClaimsArtifact(value PqcVendorClaims) (Artifact[PqcVendorClaims], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[PqcVendorClaims]{}, err
	}
	return ParsePqcVendorClaimsArtifact(payload)
}

func ParseProductDefinitionsMapArtifact(payload []byte) (Artifact[ProductDefinitionsMap], error) {
	return ParseArtifact[ProductDefinitionsMap]("product-definitions-map", payload)
}

func EncodeProductDefinitionsMapArtifact(value ProductDefinitionsMap) (Artifact[ProductDefinitionsMap], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ProductDefinitionsMap]{}, err
	}
	return ParseProductDefinitionsMapArtifact(payload)
}

func ParseRefutedRegisterArtifact(payload []byte) (Artifact[RefutedRegister], error) {
	return ParseArtifact[RefutedRegister]("refuted-register", payload)
}

func EncodeRefutedRegisterArtifact(value RefutedRegister) (Artifact[RefutedRegister], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[RefutedRegister]{}, err
	}
	return ParseRefutedRegisterArtifact(payload)
}

func ParseRemediationArtifact(payload []byte) (Artifact[Remediation], error) {
	return ParseArtifact[Remediation]("remediation", payload)
}

func EncodeRemediationArtifact(value Remediation) (Artifact[Remediation], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[Remediation]{}, err
	}
	return ParseRemediationArtifact(payload)
}

func ParseRepoLivenessArtifact(payload []byte) (Artifact[RepoLiveness], error) {
	return ParseArtifact[RepoLiveness]("repo-liveness", payload)
}

func EncodeRepoLivenessArtifact(value RepoLiveness) (Artifact[RepoLiveness], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[RepoLiveness]{}, err
	}
	return ParseRepoLivenessArtifact(payload)
}

func ParseRepoStateArtifact(payload []byte) (Artifact[RepoState], error) {
	return ParseArtifact[RepoState]("repo-state", payload)
}

func EncodeRepoStateArtifact(value RepoState) (Artifact[RepoState], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[RepoState]{}, err
	}
	return ParseRepoStateArtifact(payload)
}

func ParseReportArtifact(payload []byte) (Artifact[Report], error) {
	return ParseArtifact[Report]("report", payload)
}

func EncodeReportArtifact(value Report) (Artifact[Report], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[Report]{}, err
	}
	return ParseReportArtifact(payload)
}

func ParseRiskRatingMethodologyArtifact(payload []byte) (Artifact[RiskRatingMethodology], error) {
	return ParseArtifact[RiskRatingMethodology]("risk-rating-methodology", payload)
}

func EncodeRiskRatingMethodologyArtifact(value RiskRatingMethodology) (Artifact[RiskRatingMethodology], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[RiskRatingMethodology]{}, err
	}
	return ParseRiskRatingMethodologyArtifact(payload)
}

func ParseRunManifestArtifact(payload []byte) (Artifact[RunManifest], error) {
	return ParseArtifact[RunManifest]("run-manifest", payload)
}

func EncodeRunManifestArtifact(value RunManifest) (Artifact[RunManifest], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[RunManifest]{}, err
	}
	return ParseRunManifestArtifact(payload)
}

func ParseSessionResultArtifact(payload []byte) (Artifact[SessionResult], error) {
	return ParseArtifact[SessionResult]("session-result", payload)
}

func EncodeSessionResultArtifact(value SessionResult) (Artifact[SessionResult], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[SessionResult]{}, err
	}
	return ParseSessionResultArtifact(payload)
}

func ParseSessionTaskArtifact(payload []byte) (Artifact[SessionTask], error) {
	return ParseArtifact[SessionTask]("session-task", payload)
}

func EncodeSessionTaskArtifact(value SessionTask) (Artifact[SessionTask], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[SessionTask]{}, err
	}
	return ParseSessionTaskArtifact(payload)
}

func ParseSlaPolicyArtifact(payload []byte) (Artifact[SlaPolicy], error) {
	return ParseArtifact[SlaPolicy]("sla-policy", payload)
}

func EncodeSlaPolicyArtifact(value SlaPolicy) (Artifact[SlaPolicy], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[SlaPolicy]{}, err
	}
	return ParseSlaPolicyArtifact(payload)
}

func ParseThreatModelArtifact(payload []byte) (Artifact[ThreatModel], error) {
	return ParseArtifact[ThreatModel]("threat-model", payload)
}

func EncodeThreatModelArtifact(value ThreatModel) (Artifact[ThreatModel], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ThreatModel]{}, err
	}
	return ParseThreatModelArtifact(payload)
}

func ParseTriageArtifact(payload []byte) (Artifact[Triage], error) {
	return ParseArtifact[Triage]("triage", payload)
}

func EncodeTriageArtifact(value Triage) (Artifact[Triage], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[Triage]{}, err
	}
	return ParseTriageArtifact(payload)
}

func ParseValidationAuthorizationArtifact(payload []byte) (Artifact[ValidationAuthorization], error) {
	return ParseArtifact[ValidationAuthorization]("validation-authorization", payload)
}

func EncodeValidationAuthorizationArtifact(value ValidationAuthorization) (Artifact[ValidationAuthorization], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ValidationAuthorization]{}, err
	}
	return ParseValidationAuthorizationArtifact(payload)
}

func ParseValidationPlanArtifact(payload []byte) (Artifact[ValidationPlan], error) {
	return ParseArtifact[ValidationPlan]("validation-plan", payload)
}

func EncodeValidationPlanArtifact(value ValidationPlan) (Artifact[ValidationPlan], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[ValidationPlan]{}, err
	}
	return ParseValidationPlanArtifact(payload)
}

func ParseValidationArtifact(payload []byte) (Artifact[Validation], error) {
	return ParseArtifact[Validation]("validation", payload)
}

func EncodeValidationArtifact(value Validation) (Artifact[Validation], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[Validation]{}, err
	}
	return ParseValidationArtifact(payload)
}

func ParseVerificationArtifact(payload []byte) (Artifact[Verification], error) {
	return ParseArtifact[Verification]("verification", payload)
}

func EncodeVerificationArtifact(value Verification) (Artifact[Verification], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[Verification]{}, err
	}
	return ParseVerificationArtifact(payload)
}

func ParseVulnFindingsArtifact(payload []byte) (Artifact[VulnFindings], error) {
	return ParseArtifact[VulnFindings]("vuln-findings", payload)
}

func EncodeVulnFindingsArtifact(value VulnFindings) (Artifact[VulnFindings], error) {
	payload, err := json.Marshal(value)
	if err != nil {
		return Artifact[VulnFindings]{}, err
	}
	return ParseVulnFindingsArtifact(payload)
}
