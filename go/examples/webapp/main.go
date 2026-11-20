// Command webapp is a small web app on traust: onboard repositories, scan them, work their
// findings, watch the dashboard and build reports. One binary serves the pages and runs the work
// its pack declares.
//
//	TRAUST_URL=https://core.example.com TRAUST_TOKEN_FILE=machine-token go run .
package main

import (
	"context"
	_ "embed"
	"errors"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

//go:embed pack.yaml
var packYAML []byte

func main() {
	addr := flag.String("addr", "localhost:3000", "listen address")
	forgeHosts := flag.String("forge-hosts", "github.com,gitlab.com", "the repository hosts a run may fetch from, comma-separated")
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := run(ctx, *addr, strings.Split(*forgeHosts, ",")); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string, forgeHosts []string) error {
	core, err := traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_URL"),
		Tokens:  traust.FileTokenSource{Path: os.Getenv("TRAUST_TOKEN_FILE")},
	})
	if err != nil {
		return err
	}
	pack, err := ops.LoadPack(packYAML)
	if err != nil {
		return err
	}
	runner, err := worker(core, pack, forgeHosts)
	if err != nil {
		return err
	}
	server := &http.Server{Addr: addr, Handler: (&app{core, pack.Pin()}).routes(), ReadHeaderTimeout: 10 * time.Second}
	return serve(ctx, server, runner)
}

// worker claims the pack's orders and runs the fixture skills: a scan records as an assessment of
// the repository's default branch, a triage judges the repository's latest scan.
func worker(core *traust.Client, pack *ops.Pack, forgeHosts []string) (*ops.Runner, error) {
	scans, err := ops.AssessDefaultRef(core, traust.SourceTypeVulnScanReport, "webapp-fixture")
	if err != nil {
		return nil, err
	}
	triages, err := ops.TriageLatest(core, traust.AssessmentKindSecureCode1)
	if err != nil {
		return nil, err
	}
	return ops.NewRunner(ops.RunnerConfig{
		Client: core, Pack: pack, Executor: fixtureSkills{now: time.Now}, Assets: pack.Assets(),
		Recordings: map[string]ops.RecordingPreparer{"scan": scans, "triage": triages}, ForgeHosts: forgeHosts,
		Logger: slog.Default(), LeaseSeconds: 60, PollInterval: time.Second,
	})
}

// serve runs the pages and the runner until ctx ends or either fails, then stops both.
func serve(ctx context.Context, server *http.Server, runner *ops.Runner) error {
	ctx, cancel := context.WithCancelCause(ctx)
	worked := make(chan struct{})
	go func() { defer close(worked); cancel(runner.Run(ctx)) }()
	go func() { <-ctx.Done(); _ = server.Shutdown(context.WithoutCancel(ctx)) }()
	err := server.ListenAndServe()
	cancel(nil)
	<-worked
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	if cause := context.Cause(ctx); !errors.Is(cause, context.Canceled) {
		return cause
	}
	return nil
}
