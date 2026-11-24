// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
