// Code generated from traust-contracts 105cca4f16e2bbf52349156f1971f1dba6f53c22. DO NOT EDIT.

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
