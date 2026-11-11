//go:build unix

package ops

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// ProcessExecutor runs an order's argv as a child process: no shell, its own process group, and
// the whole group killed when the lease is lost. The command prints its JobResult as JSON on
// stdout, which is also the report; Outcome overrides that for commands that print something else.
type ProcessExecutor struct {
	Dir     string
	Env     []string // nil inherits the runner's environment
	Outcome func(stdout []byte) (traust.JobResult, error)
}

func (p ProcessExecutor) Execute(ctx context.Context, execution Execution) (ProducedReport, error) {
	if len(execution.Argv) == 0 {
		return ProducedReport{}, errors.New("ops: execution has no argv")
	}
	// #nosec G204 -- argv is the pinned, schema-validated pack's command for this asset
	cmd := exec.CommandContext(ctx, execution.Argv[0], execution.Argv[1:]...)
	cmd.Dir, cmd.Env = p.Dir, p.Env
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error { return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL) }
	stdout, runErr := cmd.Output()
	var exit *exec.ExitError
	if runErr != nil && !errors.As(runErr, &exit) {
		return ProducedReport{}, fmt.Errorf("ops: run %s: %w", execution.Argv[0], runErr)
	}
	outcome := p.Outcome
	if outcome == nil {
		outcome = jobResult
	}
	result, err := outcome(stdout)
	if err != nil {
		return ProducedReport{Payload: stdout}, errors.Join(fmt.Errorf("ops: %s output: %w", execution.Argv[0], err), runErr)
	}
	return ProducedReport{Payload: stdout, Result: result}, nil
}

func jobResult(stdout []byte) (traust.JobResult, error) {
	var result traust.JobResult
	if err := json.Unmarshal(stdout, &result); err != nil {
		return traust.JobResult{}, err
	}
	return result, checkOutcome(result)
}
