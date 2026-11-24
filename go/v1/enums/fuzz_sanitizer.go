// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

package enums

// FuzzSanitizer — The sanitizer that reported a fuzz crash.
type FuzzSanitizer string

const (
	FuzzSanitizerAddress   FuzzSanitizer = "address"
	FuzzSanitizerMemory    FuzzSanitizer = "memory"
	FuzzSanitizerUndefined FuzzSanitizer = "undefined"
	FuzzSanitizerThread    FuzzSanitizer = "thread"
	FuzzSanitizerLeak      FuzzSanitizer = "leak"
)

// FuzzSanitizerValues returns all valid FuzzSanitizer values.
func FuzzSanitizerValues() []FuzzSanitizer {
	return []FuzzSanitizer{
		FuzzSanitizerAddress,
		FuzzSanitizerMemory,
		FuzzSanitizerUndefined,
		FuzzSanitizerThread,
		FuzzSanitizerLeak,
	}
}
