// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
