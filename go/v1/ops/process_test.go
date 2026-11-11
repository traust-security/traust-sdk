//go:build unix

package ops

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

func TestProcessExecutorTurnsOutputIntoTheResult(t *testing.T) {
	for _, tc := range []struct {
		name    string
		argv    []string
		outcome func([]byte) (traust.JobResult, error)
		want    traust.Outcome
		fails   bool
	}{
		{"job result on stdout", []string{"sh", "-c", `echo '{"outcome":"degraded"}'`}, nil, traust.OutcomeDegraded, false},
		{"failure the command reports", []string{"sh", "-c", `echo '{"outcome":"failed"}'; exit 3`}, nil, traust.OutcomeFailed, false},
		{"output that is not a job result", []string{"sh", "-c", "echo done"}, nil, "", true},
		{"unknown outcome", []string{"sh", "-c", `echo '{"outcome":"exploded"}'`}, nil, "", true},
		{"custom outcome", []string{"sh", "-c", "echo done"}, func([]byte) (traust.JobResult, error) {
			return traust.JobResult{Outcome: traust.OutcomeSucceeded}, nil
		}, traust.OutcomeSucceeded, false},
		{"missing binary", []string{"traust-sdk-no-such-binary"}, nil, "", true},
		{"no argv", nil, nil, "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			produced, err := ProcessExecutor{Outcome: tc.outcome}.Execute(context.Background(), Execution{Argv: tc.argv})
			if (err != nil) != tc.fails || produced.Result.Outcome != tc.want {
				t.Fatalf("result=%+v err=%v", produced.Result, err)
			}
		})
	}
}

func TestProcessExecutorStopsDescendantsOnCancel(t *testing.T) {
	marker := filepath.Join(t.TempDir(), "survived")
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	script := `(sleep 0.5; touch "$1") & sleep 5`
	_, err := ProcessExecutor{}.Execute(ctx, Execution{Argv: []string{"sh", "-c", script, "sh", marker}})
	if err == nil {
		t.Fatal("a cancelled command reported success")
	}
	time.Sleep(700 * time.Millisecond)
	if _, err := os.Stat(marker); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("descendant outlived the cancel: %v", err)
	}
}
