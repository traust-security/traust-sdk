// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// PqcPathClass — Values of `path_class` in pqc-facts.schema.json.
type PqcPathClass string

const (
	PqcPathClassFirstParty PqcPathClass = "first_party"
	PqcPathClassVendor     PqcPathClass = "vendor"
	PqcPathClassTestDocs   PqcPathClass = "test_docs"
)

// PqcPathClassValues returns all valid PqcPathClass values.
func PqcPathClassValues() []PqcPathClass {
	return []PqcPathClass{
		PqcPathClassFirstParty,
		PqcPathClassVendor,
		PqcPathClassTestDocs,
	}
}
