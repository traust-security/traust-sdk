// Code generated from traust-contracts bf4eb27b9eff357587fe1369db3b43e4a9534eba. DO NOT EDIT.

package enums

// LoggingOutput — Built-in destinations for traust-core operational logs.
type LoggingOutput string

const (
	LoggingOutputStderr LoggingOutput = "stderr"
	LoggingOutputFile   LoggingOutput = "file"
)

// LoggingOutputValues returns all valid LoggingOutput values.
func LoggingOutputValues() []LoggingOutput {
	return []LoggingOutput{
		LoggingOutputStderr,
		LoggingOutputFile,
	}
}
