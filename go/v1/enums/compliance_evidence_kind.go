// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
