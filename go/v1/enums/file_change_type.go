// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// FileChangeType — Values of `change_type` in remediation.schema.json.
type FileChangeType string

const (
	FileChangeTypeModified FileChangeType = "modified"
	FileChangeTypeAdded    FileChangeType = "added"
	FileChangeTypeDeleted  FileChangeType = "deleted"
	FileChangeTypeRenamed  FileChangeType = "renamed"
)

// FileChangeTypeValues returns all valid FileChangeType values.
func FileChangeTypeValues() []FileChangeType {
	return []FileChangeType{
		FileChangeTypeModified,
		FileChangeTypeAdded,
		FileChangeTypeDeleted,
		FileChangeTypeRenamed,
	}
}
