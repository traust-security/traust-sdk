package ops_test

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

const pack = `pack: example
assets:
  ledger_audit:
    run: [traust, ledger, audit, --run-id, "{run_id}"]
`

func ExampleLoadPack() {
	p, err := ops.LoadPack([]byte(pack))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p.Assets(), p.Pin().Name)
	// Output: [ledger_audit] pack:example
}

// A runner claims the orders its pack can run, runs each asset's command and completes it. The
// person who asks for a run pins the order to the same pack.
func ExampleRunner() {
	ctx := context.Background()
	p, err := ops.LoadPack([]byte(pack))
	if err != nil {
		log.Fatal(err)
	}
	runnerClient, err := traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_CORE_URL"),
		Tokens:  traust.FileTokenSource{Path: "/var/run/secrets/traust-core/token"},
	})
	if err != nil {
		log.Fatal(err)
	}
	runner, err := ops.NewRunner(ops.RunnerConfig{
		Client: runnerClient, Pack: p, Executor: ops.ProcessExecutor{},
		Logger: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
		Assets: p.Assets(), LeaseSeconds: 300, PollInterval: 10 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	person, err := traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_CORE_URL"), Tokens: traust.StaticToken(os.Getenv("TRAUST_TOKEN")),
	})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := person.Submit(ctx, traust.SubmitRequest{
		Order:  traust.OrderRequest{Asset: "ledger_audit", Priority: 1},
		Config: &[]traust.ConfigDigest{p.Pin()},
	}); err != nil {
		log.Fatal(err)
	}
	log.Fatal(runner.Run(ctx))
}
