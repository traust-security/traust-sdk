// Code generated from traust-contracts 485578cbe584ad224d57b2c7c778174c0b66ba7c. DO NOT EDIT.

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
