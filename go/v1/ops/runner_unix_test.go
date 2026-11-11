//go:build unix

package ops

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRunnerStopsACommandsDescendantsOnLeaseLoss(t *testing.T) {
	dir := t.TempDir()
	started, marker := filepath.Join(dir, "started"), filepath.Join(dir, "descendant-survived")
	runner, server := queuedRunner(t, ProcessExecutor{})
	runner.config.HeartbeatInterval = 50 * time.Millisecond
	server.heartbeatStatus = 409
	server.heartbeatFailAt = 6
	asset := server.pack.assets["ledger_audit"]
	asset.Run = []string{"sh", "-c", `touch "$2"; (sleep 0.5; touch "$1") & sleep 10`, "sh", marker, started}
	server.pack.assets["ledger_audit"] = asset
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
	if _, err := os.Stat(started); err != nil {
		t.Fatalf("the command never ran, so the lease was lost before execution: %v", err)
	}
	time.Sleep(700 * time.Millisecond)
	if _, err := os.Stat(marker); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("descendant survived lease loss: %v", err)
	}
}
