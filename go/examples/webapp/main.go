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

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/findings"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/identity"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/insight"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/inventory"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/ledger"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/runs"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
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
	cfg, err := traust.LoadConfig()
	if err != nil {
		return err
	}
	core, err := traust.NewClient(cfg)
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
	server := &http.Server{Addr: addr, Handler: routes(core, pack.Pin()), ReadHeaderTimeout: 10 * time.Second}
	return serve(ctx, server, runner)
}

// routes wires the web app: the shared server, then each feature's pages.
func routes(core *traust.Client, pack traust.ConfigDigest) http.Handler {
	return web.Handler(
		identity.New(core).Routes,
		inventory.New(core).Routes,
		runs.New(core, pack).Routes,
		findings.New(core).Routes,
		ledger.New(core).Routes,
		insight.New(core).Routes,
	)
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
