// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
