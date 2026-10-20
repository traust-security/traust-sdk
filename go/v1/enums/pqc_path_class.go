// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
