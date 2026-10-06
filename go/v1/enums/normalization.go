package enums

import "fmt"

// EnumValue is one vocabulary/value pair. Values are returned by value so
// callers cannot alter the registry or another caller's read view.
type EnumValue struct {
	Enum  string
	Value string
}

// Normalization is an ordered read view, not a replacement event payload.
// Its zero value is empty; successful Normalize calls always retain at least
// one pair, including the original pair for a one-way drop.
type Normalization struct {
	single       EnumValue
	replacements []EnumValue
	count        int
	Dropped      bool
}

// Len returns the number of pairs in the read view.
func (n Normalization) Len() int {
	return n.count
}

// At returns a pair by value in declared replacement order. It panics if index
// is outside [0, Len()), like indexing a slice.
func (n Normalization) At(index int) EnumValue {
	if index < 0 || index >= n.count {
		panic("enum normalization index out of range")
	}
	if n.replacements != nil {
		return n.replacements[index]
	}
	return n.single
}

// Normalize interprets only declared replacements from the generated registry.
// Retired keys are independent of generated constants. Unknown values in a known
// enum retain their exact spelling; unknown enums return an error, not a drop.
// No event, identifier, serialized payload, or registry data is mutated. Results
// expose no shared mutable slice and known-enum calls require no allocations.
func Normalize(enumName, value string) (Normalization, error) {
	deprecated, ok := enumNormalizationRegistry[enumName]
	if !ok {
		return Normalization{}, fmt.Errorf("unknown enum %q", enumName)
	}
	if result, ok := deprecated[value]; ok {
		return result, nil
	}
	return Normalization{single: EnumValue{Enum: enumName, Value: value}, count: 1}, nil
}
