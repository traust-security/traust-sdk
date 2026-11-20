// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// FuzzCrashClass — What a fuzz crasher did to the target.
type FuzzCrashClass string

const (
	FuzzCrashClassPanic              FuzzCrashClass = "panic"
	FuzzCrashClassIndexOutOfRange    FuzzCrashClass = "index_out_of_range"
	FuzzCrashClassNilDereference     FuzzCrashClass = "nil_dereference"
	FuzzCrashClassOutOfMemory        FuzzCrashClass = "out_of_memory"
	FuzzCrashClassTimeout            FuzzCrashClass = "timeout"
	FuzzCrashClassIntegerOverflow    FuzzCrashClass = "integer_overflow"
	FuzzCrashClassMemoryCorruption   FuzzCrashClass = "memory_corruption"
	FuzzCrashClassMemoryLeak         FuzzCrashClass = "memory_leak"
	FuzzCrashClassUncaughtException  FuzzCrashClass = "uncaught_exception"
	FuzzCrashClassProcessExit        FuzzCrashClass = "process_exit"
	FuzzCrashClassInvariantViolation FuzzCrashClass = "invariant_violation"
)

// FuzzCrashClassValues returns all valid FuzzCrashClass values.
func FuzzCrashClassValues() []FuzzCrashClass {
	return []FuzzCrashClass{
		FuzzCrashClassPanic,
		FuzzCrashClassIndexOutOfRange,
		FuzzCrashClassNilDereference,
		FuzzCrashClassOutOfMemory,
		FuzzCrashClassTimeout,
		FuzzCrashClassIntegerOverflow,
		FuzzCrashClassMemoryCorruption,
		FuzzCrashClassMemoryLeak,
		FuzzCrashClassUncaughtException,
		FuzzCrashClassProcessExit,
		FuzzCrashClassInvariantViolation,
	}
}
