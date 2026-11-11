//go:build !unix

package ops

import (
	"context"
	"errors"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// ProcessExecutor needs unix process groups to stop a command's descendants with it.
type ProcessExecutor struct {
	Dir     string
	Env     []string
	Outcome func(stdout []byte) (traust.JobResult, error)
}

func (ProcessExecutor) Execute(context.Context, Execution) (ProducedReport, error) {
	return ProducedReport{}, errors.New("ops: ProcessExecutor needs a unix process group")
}
