// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
