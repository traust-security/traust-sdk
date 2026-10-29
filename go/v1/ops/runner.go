package ops

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/api"
)

const defaultRunnerRequestTimeout = time.Minute

var commandParameter = regexp.MustCompile(`^\{([a-z][a-z0-9_]*)\}$`)

// RunnerConfig binds a pinned pack and the assets this host can execute.
type RunnerConfig struct {
	Client            *api.BoundClient
	Pack              *Pack
	Executor          Executor
	Assets            []string
	LeaseSeconds      int
	HeartbeatInterval time.Duration
	RequestTimeout    time.Duration
	PollInterval      time.Duration
}

// Runner claims and executes deterministic fleet-wide pack commands.
type Runner struct {
	config RunnerConfig
}

// NewRunner refuses assets this first command runner cannot resolve safely.
func NewRunner(cfg RunnerConfig) (*Runner, error) {
	if cfg.Client == nil || cfg.Pack == nil || cfg.Executor == nil {
		return nil, errors.New("ops: runner requires Client, Pack, and Executor")
	}
	if cfg.LeaseSeconds <= 0 || cfg.LeaseSeconds > 24*60*60 || len(cfg.Assets) == 0 {
		return nil, errors.New("ops: runner requires assets and a lease of 1..86400 seconds")
	}
	if cfg.HeartbeatInterval == 0 {
		cfg.HeartbeatInterval = time.Duration(cfg.LeaseSeconds) * time.Second / 3
	}
	if cfg.RequestTimeout == 0 {
		cfg.RequestTimeout = defaultRunnerRequestTimeout
	}
	if cfg.HeartbeatInterval <= 0 || cfg.HeartbeatInterval >= time.Duration(cfg.LeaseSeconds)*time.Second || cfg.RequestTimeout <= 0 || cfg.PollInterval <= 0 {
		return nil, errors.New("ops: runner needs positive request and poll durations and a heartbeat shorter than the lease")
	}
	seen := make(map[string]bool, len(cfg.Assets))
	for _, name := range cfg.Assets {
		if seen[name] {
			return nil, fmt.Errorf("ops: duplicate asset %q", name)
		}
		seen[name] = true
		asset, ok := cfg.Pack.assets[name]
		if !ok || len(asset.Run) == 0 {
			return nil, fmt.Errorf("ops: asset %q has no runnable command in the pinned pack", name)
		}
		if asset.Partitions != "" || asset.Model != nil || len(asset.Params) != 0 {
			return nil, fmt.Errorf("ops: asset %q needs input resolution beyond this fleet-wide command runner", name)
		}
		for _, token := range asset.Run {
			if match := commandParameter.FindStringSubmatch(token); match != nil {
				if match[1] != "run_id" {
					return nil, fmt.Errorf("ops: asset %q requires unresolved parameter %q", name, match[1])
				}
			} else if strings.ContainsAny(token, "{}") {
				return nil, fmt.Errorf("ops: asset %q has an unresolved command token %q", name, token)
			}
		}
	}
	cfg.Assets = append([]string(nil), cfg.Assets...)
	return &Runner{config: cfg}, nil
}

// Run drains eligible orders and waits after an empty claim until cancellation.
func (r *Runner) Run(ctx context.Context) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		order, err := r.RunOnce(ctx)
		if err != nil {
			return err
		}
		if order != nil {
			continue
		}
		timer := time.NewTimer(r.config.PollInterval)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
	}
}

// RunOnce claims one eligible order, or returns nil when the queue is empty.
func (r *Runner) RunOnce(ctx context.Context) (*api.RunOrder, error) {
	requestCtx, cancel := context.WithTimeout(ctx, r.config.RequestTimeout)
	assets := append([]string(nil), r.config.Assets...)
	response, err := r.config.Client.ClaimWithResponse(requestCtx, api.ClaimRequest{
		Assets: &assets, LeaseSeconds: r.config.LeaseSeconds,
	})
	cancel()
	if err != nil {
		return nil, fmt.Errorf("ops: claim order: %w", err)
	}
	if response.StatusCode() == http.StatusNoContent {
		return nil, nil
	}
	order, err := nativeResponse(response.StatusCode(), response.Body, response.JSON200)
	if err != nil {
		return nil, fmt.Errorf("ops: claim order: %w", err)
	}
	if err := usableOrder(order); err != nil {
		return nil, fmt.Errorf("ops: claimed order: %w", err)
	}
	asset, allowed := r.config.Pack.assets[order.Asset]
	if !allowed || !slices.Contains(r.config.Assets, order.Asset) {
		return r.finish(ctx, order, api.JobResult{Outcome: api.OutcomeRefused, Reason: stringPtr("asset-not-in-pack")})
	}
	if order.ArtifactId == nil || order.RepositoryId != nil {
		return r.finish(ctx, order, api.JobResult{Outcome: api.OutcomeRefused, Reason: stringPtr("bad_request")})
	}
	match, err := r.matchesPack(ctx, *order.ArtifactId)
	if err != nil {
		return order, fmt.Errorf("ops: verify pack for order %s: %w", *order.Id, err)
	}
	if !match {
		return r.finish(ctx, order, api.JobResult{Outcome: api.OutcomeRefused, Reason: stringPtr("pack version mismatch")})
	}
	argv := make([]string, 0, len(asset.Run))
	for _, token := range asset.Run {
		if match := commandParameter.FindStringSubmatch(token); match != nil {
			argv = append(argv, order.Id.String())
		} else {
			argv = append(argv, token)
		}
	}
	run := RecordingRun{
		Client: r.config.Client, Executor: r.config.Executor,
		Start:             api.StartRecordingRequest{LeaseSeconds: r.config.LeaseSeconds},
		HeartbeatInterval: r.config.HeartbeatInterval, RequestTimeout: r.config.RequestTimeout,
	}
	owned, err := run.heartbeat(ctx, *order, order)
	if err != nil {
		return order, err
	}
	produced, err := run.execute(ctx, *order, owned, Execution{Order: *order, Argv: argv})
	if err != nil {
		var lost *LeaseLostError
		if errors.As(err, &lost) || ctx.Err() != nil {
			return order, err
		}
		finished, finishErr := r.finish(ctx, order, api.JobResult{Outcome: api.OutcomeFailed, Reason: stringPtr("runner-crashed")})
		return finished, errors.Join(err, finishErr)
	}
	if err := validateJobResult(produced.Result); err != nil {
		finished, finishErr := r.finish(ctx, order, api.JobResult{Outcome: api.OutcomeFailed, Reason: stringPtr("runner-crashed")})
		return finished, errors.Join(err, finishErr)
	}
	return r.finish(ctx, order, produced.Result)
}

func (r *Runner) matchesPack(ctx context.Context, artifactID uuid.UUID) (bool, error) {
	requestCtx, cancel := context.WithTimeout(ctx, r.config.RequestTimeout)
	response, err := r.config.Client.GetArtifactContentWithResponse(requestCtx, artifactID)
	cancel()
	if err != nil {
		return false, err
	}
	if response.StatusCode() != http.StatusOK {
		_, err := nativeResponse[api.Artifact](response.StatusCode(), response.Body, nil)
		return false, err
	}
	var manifest struct {
		Artifact string `json:"artifact"`
		Config   []struct {
			Name   string `json:"name"`
			Digest string `json:"digest"`
		} `json:"config"`
	}
	if err := json.Unmarshal(response.Body, &manifest); err != nil || manifest.Artifact != "run-manifest" {
		return false, errors.New("Core returned an invalid run manifest")
	}
	var found int
	for _, item := range manifest.Config {
		if item.Name == "pack:"+r.config.Pack.name {
			found++
			if item.Digest != r.config.Pack.digest {
				return false, nil
			}
		}
	}
	return found == 1, nil
}

func (r *Runner) finish(ctx context.Context, order *api.RunOrder, result api.JobResult) (*api.RunOrder, error) {
	run := RecordingRun{
		Client:         r.config.Client,
		Start:          api.StartRecordingRequest{LeaseSeconds: r.config.LeaseSeconds},
		RequestTimeout: r.config.RequestTimeout,
	}
	owned, err := run.heartbeat(ctx, *order, order)
	if err != nil {
		return order, err
	}
	completed, err := run.complete(ctx, owned, result)
	if err != nil {
		return order, err
	}
	return completed, nil
}

func validateJobResult(result api.JobResult) error {
	switch result.Outcome {
	case api.OutcomeSucceeded, api.OutcomeDegraded, api.OutcomeRefused, api.OutcomeFailed:
		return nil
	default:
		return fmt.Errorf("ops: invalid job outcome %q", result.Outcome)
	}
}

func stringPtr(value string) *string { return &value }
