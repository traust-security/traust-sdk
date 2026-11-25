// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

package enums

// ImpactEvidenceBinaryLinkedLibrary — DT_NEEDED linked-library scan of shipped image binaries (C/C++).
type ImpactEvidenceBinaryLinkedLibrary string

const (
	ImpactEvidenceBinaryLinkedLibraryLinked    ImpactEvidenceBinaryLinkedLibrary = "linked"
	ImpactEvidenceBinaryLinkedLibraryNotLinked ImpactEvidenceBinaryLinkedLibrary = "not_linked"
)

// ImpactEvidenceBinaryLinkedLibraryValues returns all valid ImpactEvidenceBinaryLinkedLibrary values.
func ImpactEvidenceBinaryLinkedLibraryValues() []ImpactEvidenceBinaryLinkedLibrary {
	return []ImpactEvidenceBinaryLinkedLibrary{
		ImpactEvidenceBinaryLinkedLibraryLinked,
		ImpactEvidenceBinaryLinkedLibraryNotLinked,
	}
}
