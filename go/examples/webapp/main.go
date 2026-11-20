// Command webapp is a small web app on traust: sign in, see your repositories, start a run and
// watch it finish. One binary serves the pages and runs the work its pack declares.
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
	"syscall"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

//go:embed pack.yaml
var packYAML []byte

func main() {
	addr := flag.String("addr", "localhost:3000", "listen address")
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := run(ctx, *addr); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string) error {
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
	runner, err := ops.NewRunner(ops.RunnerConfig{
		Client: core, Pack: pack, Executor: ops.ProcessExecutor{}, Assets: pack.Assets(),
		Logger: slog.Default(), LeaseSeconds: 60, PollInterval: time.Second,
	})
	if err != nil {
		return err
	}
	server := &http.Server{Addr: addr, Handler: (&app{core, pack.Pin()}).routes(), ReadHeaderTimeout: 10 * time.Second}
	return serve(ctx, server, runner)
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
