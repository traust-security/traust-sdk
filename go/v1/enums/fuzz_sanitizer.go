// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
