// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
