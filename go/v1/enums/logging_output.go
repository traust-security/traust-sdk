// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
