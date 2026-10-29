//go:build e2e && (darwin || linux)

package ops

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/api"
)

func TestE2ERunnerLedgerAudit(t *testing.T) {
	packBytes, err := os.ReadFile(os.Getenv("TRAUST_E2E_PACK"))
	if err != nil {
		t.Fatal(err)
	}
	pack, err := LoadPack(packBytes)
	if err != nil {
		t.Fatal(err)
	}
	client, err := api.NewBoundClient(api.ClientConfig{
		BaseURL: os.Getenv("TRAUST_E2E_URL"),
		Tokens:  api.FileTokenSource{Path: os.Getenv("TRAUST_E2E_TOKEN_FILE")},
	})
	if err != nil {
		t.Fatal(err)
	}
	python, configHome := os.Getenv("TRAUST_E2E_PYTHON"), os.Getenv("TRAUST_E2E_CONFIG_HOME")
	if python == "" || configHome == "" {
		t.Fatal("e2e requires a Traust Python executable and temporary config home")
	}
	executor := executorFunc(func(ctx context.Context, execution Execution) (ProducedReport, error) {
		if len(execution.Argv) == 0 || execution.Argv[0] != "python3" {
			return ProducedReport{}, errors.New("unexpected pack command")
		}
		command := exec.CommandContext(ctx, python, execution.Argv[1:]...)
		// A lost lease must stop descendants as well as the command process.
		command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		command.Cancel = func() error { return syscall.Kill(-command.Process.Pid, syscall.SIGKILL) }
		command.Env = []string{
			"PATH=" + os.Getenv("PATH"),
			"PYTHONPATH=" + os.Getenv("TRAUST_E2E_PYTHONPATH"),
			"TRAUST_CONFIG_HOME=" + configHome,
			"HOME=" + os.Getenv("HOME"),
		}
		output, err := command.Output()
		var exit *exec.ExitError
		if err != nil && !errors.As(err, &exit) {
			return ProducedReport{}, err
		}
		var result api.JobResult
		if err := json.Unmarshal(output, &result); err != nil {
			return ProducedReport{}, err
		}
		return ProducedReport{Payload: output, Result: result}, nil
	})
	runner, err := NewRunner(RunnerConfig{
		Client: client, Pack: pack, Executor: executor, Assets: []string{"ledger_audit"},
		LeaseSeconds: 30, HeartbeatInterval: 5 * time.Second, RequestTimeout: 10 * time.Second, PollInterval: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	order, err := runner.RunOnce(ctx)
	if err != nil || order == nil || order.State != api.RunOrderStateSucceeded {
		t.Fatalf("completed order=%+v err=%v", order, err)
	}
	if idle, err := runner.RunOnce(ctx); err != nil || idle != nil {
		t.Fatalf("empty queue order=%+v err=%v", idle, err)
	}
	if !strings.HasPrefix(order.RequestedBy, "human:") {
		t.Fatalf("order lost its planning principal: %q", order.RequestedBy)
	}
}

func TestE2ERunnerLocalCommand(t *testing.T) {
	packBytes, err := os.ReadFile("testdata/local-command.yaml")
	if err != nil {
		t.Fatal(err)
	}
	pack, err := LoadPack(packBytes)
	if err != nil {
		t.Fatal(err)
	}
	client, err := api.NewBoundClient(api.ClientConfig{
		BaseURL: os.Getenv("TRAUST_E2E_URL"),
		Tokens:  api.FileTokenSource{Path: os.Getenv("TRAUST_E2E_TOKEN_FILE")},
	})
	if err != nil {
		t.Fatal(err)
	}
	executor := executorFunc(func(ctx context.Context, execution Execution) (ProducedReport, error) {
		if len(execution.Argv) != 4 || execution.Argv[0] != "python3" || execution.Argv[1] != "-c" || execution.Argv[3] != execution.Order.Id.String() {
			return ProducedReport{}, errors.New("unexpected resolved command")
		}
		command := exec.CommandContext(ctx, execution.Argv[0], execution.Argv[1:]...)
		command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		command.Cancel = func() error { return syscall.Kill(-command.Process.Pid, syscall.SIGKILL) }
		output, err := command.Output()
		if err != nil {
			return ProducedReport{}, err
		}
		if string(output) != "completed\n" {
			return ProducedReport{}, errors.New("unexpected command output")
		}
		return ProducedReport{Payload: output, Result: api.JobResult{Outcome: api.OutcomeSucceeded}}, nil
	})
	runner, err := NewRunner(RunnerConfig{
		Client: client, Pack: pack, Executor: executor, Assets: []string{"command_result"},
		LeaseSeconds: 30, HeartbeatInterval: 5 * time.Second, RequestTimeout: 10 * time.Second, PollInterval: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	order, err := runner.RunOnce(ctx)
	if err != nil || order == nil || order.State != api.RunOrderStateSucceeded {
		t.Fatalf("completed order=%+v err=%v", order, err)
	}
	if idle, err := runner.RunOnce(ctx); err != nil || idle != nil {
		t.Fatalf("empty queue order=%+v err=%v", idle, err)
	}
}
