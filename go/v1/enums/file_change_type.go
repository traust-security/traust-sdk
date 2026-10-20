// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
