// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
