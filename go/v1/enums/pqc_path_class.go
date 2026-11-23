// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
