package main

// The SDK's reference consumer, version one: written against the SDK as it exists today, so
// every gap this app has to build around is felt before the appliance designs it away. Each
// flow marks its gap with GAP-n; GAPLOG.md explains them. The second version of this app,
// after the appliance lands, is the receipt: every line it keeps should be a domain decision.
//
// Environment: TRAUST_CORE_URL, TRAUST_TOKEN_FILE (the person), TRAUST_RUNNER_TOKEN_FILE
// (the worker). `make -C go example` runs the whole thing against a local traust-server.

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/traust-security/traust-sdk/go/v1/ops"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

const (
	assetName    = "smoke"
	fixtureURL   = "https://e2e-forge.test/e2e/fixture"
	leaseSeconds = 30
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute) // GAP-5: no drain
	defer cancel()                                                          // protocol — a long-running consumer must invent its own "finish in-flight,
	// then exit" sequencing.

	person, err := client("TRAUST_TOKEN_FILE")
	fatal("person client", err)
	worker, err := client("TRAUST_RUNNER_TOKEN_FILE")
	fatal("runner client", err)

	repositoryID := onboard(ctx, person)
	runOneAsset(ctx, person, worker, repositoryID)
	readOneScreen(ctx, person, repositoryID)
	downloadOneReport(ctx, person, repositoryID)
}

func client(tokenEnv string) (*traust.Client, error) {
	return traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_CORE_URL"),
		Tokens:  traust.FileTokenSource{Path: os.Getenv(tokenEnv)},
	})
}

// Flow 1 — run one asset: onboard the fixture repository, submit an order pinned to our pack,
// and run it to completion on the worker's runner.
func runOneAsset(ctx context.Context, person, worker *traust.Client, repositoryID openapi_types.UUID) {
	pack, err := ops.LoadPack(mustRead("pack.yaml"))
	fatal("load pack", err)

	order, err := person.Submit(ctx, traust.SubmitRequest{
		Order:  traust.OrderRequest{Asset: assetName, Priority: 1},
		Config: &[]traust.ConfigDigest{pack.Pin()}, // GAP-2: the submitter must know the pack's
		// digest to pin it — no start door that resolves the subject for you.
	})
	fatal("submit order", err)

	runner, err := ops.NewRunner(ops.RunnerConfig{
		Client: worker, Pack: pack, Executor: ops.ProcessExecutor{},
		Logger:       slog.New(slog.NewTextHandler(os.Stderr, nil)),
		Assets:       pack.Assets(),
		LeaseSeconds: leaseSeconds, PollInterval: time.Second,
	})
	fatal("new runner", err)

	// GAP-3: one claim, once. A real worker loops — and the loop is the consumer's: claim
	// cooldown, bounded lanes, fair rotation, panic containment and drain are all ours to
	// build (the reference consumer in production built them; its integration layer is
	// thousands of lines).
	done, err := runner.RunOnce(ctx)
	fatal("run order", err)
	if done == nil || order.Id == nil || done.Id == nil || *done.Id != *order.Id ||
		done.State != traust.RunOrderStateSucceeded {
		state := "missing"
		if done != nil && done.Id != nil {
			state = string(done.State)
		}
		fatal("run order", fmt.Errorf("order %s: state %s", order.Id, state))
	}
	fmt.Printf("flow 1: order %s succeeded\n", order.Id)
}

// onboard registers the fixture repository and returns its core id.
func onboard(ctx context.Context, person *traust.Client) openapi_types.UUID {
	// GAP-1: onboarding is not idempotent — a product name cannot be re-onboarded, so this
	// consumer invents a unique one per run instead of asking "what do I already have?".
	product := fmt.Sprintf("example-%d", time.Now().Unix())
	label := "Example"
	result, err := person.Onboard(ctx, traust.OnboardingRequest{
		Product: product,
		Label:   &label, // GAP-1 includes this: registering needs the label dance, and
		// still cannot re-run against an existing product.
		Repositories: &[]traust.RepositorySpec{{Url: fixtureURL}},
	})
	fatal("onboard", err)
	repositoryID := (*result.RepositoryIds)[0]
	fmt.Printf("onboarded %s as repository %s\n", product, repositoryID)
	return repositoryID
}

// Flow 2 — read one dashboard screen: live counts for our repository.
func readOneScreen(ctx context.Context, person *traust.Client, repositoryID openapi_types.UUID) {
	ids := []openapi_types.UUID{repositoryID}
	// GAP-4: this is a report-shaped fold, not a read model — the shape is fixed (five int
	// columns, severity a stringly map), it cannot be grouped or filtered by the caller, and
	// a screen of N repositories means N calls (SummarizeFindings takes exactly one) unless
	// the whole-estate variant is used.
	rollups, err := person.SummarizeFindingsIn(ctx, &traust.SummarizeFindingsInParams{
		RepositoryIds: &ids,
	})
	fatal("read screen", err)
	if len(rollups) == 0 { // an empty fold — no findings recorded against these repositories yet
		fmt.Println("flow 2: screen is empty — no findings recorded for this repository")
	}
	for _, rollup := range rollups {
		fmt.Printf("flow 2: %s: %d open (%v), %d accepted\n",
			*rollup.Key, rollup.Open, rollup.OpenBySeverity, rollup.Accepted)
	}
}

// Flow 3 — download one report.
func downloadOneReport(ctx context.Context, person *traust.Client, repositoryID openapi_types.UUID) {
	format := traust.ReportFormatMarkdown
	// GAP-6: the report is computed on the request path — the call blocks until the whole
	// view is folded and rendered, and returns bytes, not a handle to a job's artifact.
	report, err := person.GetInsightReport(ctx, traust.FindingsRollup, &traust.GetInsightReportParams{
		RepositoryId: &repositoryID,
		Format:       &format,
	})
	fatal("download report", err)
	name := fmt.Sprintf("findings-rollup-%s.md", repositoryID)
	fatal("write report", os.WriteFile(name, report, 0o600))
	fmt.Printf("flow 3: wrote %s (%d bytes)\n", name, len(report))
}

func mustRead(name string) []byte {
	raw, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return raw
}

func fatal(step string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "example: %s: %v\n", step, err)
		os.Exit(1)
	}
}
