// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

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
