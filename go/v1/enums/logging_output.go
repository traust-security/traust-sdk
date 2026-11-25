// Code generated from traust-contracts 847e012cb3bf7af5dc16ade911ee0c813962b54e. DO NOT EDIT.

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
