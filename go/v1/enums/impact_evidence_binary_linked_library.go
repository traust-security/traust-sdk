// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
