// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
