// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

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
