//go:build e2e && (darwin || linux)

package ops

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/api"
)

func processGroupCommand(ctx context.Context, name string, args ...string) *exec.Cmd {
	command := exec.CommandContext(ctx, name, args...)
	// A lost lease must stop descendants as well as the command process.
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	command.Cancel = func() error { return syscall.Kill(-command.Process.Pid, syscall.SIGKILL) }
	return command
}

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
		command := processGroupCommand(ctx, python, execution.Argv[1:]...)
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
		command := processGroupCommand(ctx, execution.Argv[0], execution.Argv[1:]...)
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

func TestE2ERunnerLocalCommandStopsDescendantsOnLeaseLoss(t *testing.T) {
	python, err := exec.LookPath("python3")
	if err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(t.TempDir(), "descendant-survived")
	const script = `import subprocess, sys, time
subprocess.Popen([sys.executable, "-c", "import pathlib,sys,time; print('child-ready', flush=True); time.sleep(0.5); pathlib.Path(sys.argv[1]).write_text('survived')", sys.argv[1]])
time.sleep(10)`
	var server *queuedServer
	runner, mock := queuedRunner(t, executorFunc(func(ctx context.Context, _ Execution) (ProducedReport, error) {
		command := processGroupCommand(ctx, python, "-c", script, marker)
		stdout, err := command.StdoutPipe()
		if err != nil {
			return ProducedReport{}, err
		}
		if err := command.Start(); err != nil {
			return ProducedReport{}, err
		}
		defer func() {
			if command.ProcessState == nil {
				_ = command.Cancel()
				_ = command.Wait()
			}
		}()
		ready, err := bufio.NewReader(stdout).ReadString('\n')
		if err != nil || ready != "child-ready\n" {
			return ProducedReport{}, errors.New("child process did not start")
		}
		server.mu.Lock()
		server.heartbeatStatus = 409
		server.mu.Unlock()
		return ProducedReport{}, command.Wait()
	}))
	server = mock
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	order, err := runner.RunOnce(ctx)
	var lost *LeaseLostError
	if order == nil || !errors.As(err, &lost) {
		t.Fatalf("order=%+v err=%v", order, err)
	}
	events, _ := server.snapshot()
	if strings.Contains(strings.Join(events, ","), "complete") {
		t.Fatalf("completed after lease loss: %v", events)
	}
	time.Sleep(700 * time.Millisecond)
	if _, err := os.Stat(marker); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("descendant survived lease loss: %v", err)
	}
}
