// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

package enums

// ComplianceEvidenceKind — Values of `kind` in compliance-assessment.schema.json.
type ComplianceEvidenceKind string

const (
	ComplianceEvidenceKindInventorySnapshotExcerpt ComplianceEvidenceKind = "inventory_snapshot_excerpt"
	ComplianceEvidenceKindIacFile                  ComplianceEvidenceKind = "iac_file"
	ComplianceEvidenceKindScannerFact              ComplianceEvidenceKind = "scanner_fact"
	ComplianceEvidenceKindPolicyObject             ComplianceEvidenceKind = "policy_object"
	ComplianceEvidenceKindHumanArtifact            ComplianceEvidenceKind = "human_artifact"
)

// ComplianceEvidenceKindValues returns all valid ComplianceEvidenceKind values.
func ComplianceEvidenceKindValues() []ComplianceEvidenceKind {
	return []ComplianceEvidenceKind{
		ComplianceEvidenceKindInventorySnapshotExcerpt,
		ComplianceEvidenceKindIacFile,
		ComplianceEvidenceKindScannerFact,
		ComplianceEvidenceKindPolicyObject,
		ComplianceEvidenceKindHumanArtifact,
	}
}
