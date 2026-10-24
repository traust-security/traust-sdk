// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
