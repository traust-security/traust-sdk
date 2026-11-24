// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
