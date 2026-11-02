package ops

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// RunnerConfig binds a pinned pack and the assets this host runs.
type RunnerConfig struct {
	Client   *traust.Client
	Pack     *Pack
	Executor Executor
	Assets   []string
	// ForgeHosts are the repository hosts {repo_url} may name: the keys of core's forge_hosts.
	// Without them a repository-bound order is always refused.
	ForgeHosts        []string
	LeaseSeconds      int
	HeartbeatInterval time.Duration
	RequestTimeout    time.Duration
	PollInterval      time.Duration
}

// Runner claims queued orders for its assets and runs each one's pinned pack command.
type Runner struct {
	config RunnerConfig
}

// NewRunner refuses assets a runner cannot run from its pinned pack, and an asset naming
// {repo_url} without the forge hosts to check its repository against.
func NewRunner(cfg RunnerConfig) (*Runner, error) {
	if cfg.Client == nil || cfg.Pack == nil || cfg.Executor == nil {
		return nil, errors.New("ops: runner needs a Client, a Pack and an Executor")
	}
	if len(cfg.Assets) == 0 || cfg.LeaseSeconds <= 0 || cfg.LeaseSeconds > maxLeaseSeconds {
		return nil, fmt.Errorf("ops: runner needs assets and a lease of 1..%d seconds", maxLeaseSeconds)
	}
	lease := leaseDuration(cfg.LeaseSeconds)
	cfg.HeartbeatInterval = cmp.Or(cfg.HeartbeatInterval, lease/3)
	cfg.RequestTimeout = cmp.Or(cfg.RequestTimeout, defaultRequestTimeout)
	if cfg.HeartbeatInterval <= 0 || cfg.HeartbeatInterval >= lease || cfg.RequestTimeout <= 0 || cfg.PollInterval <= 0 {
		return nil, errors.New("ops: runner needs positive request and poll intervals and a heartbeat shorter than the lease")
	}
	for i, name := range cfg.Assets {
		if slices.Contains(cfg.Assets[:i], name) {
			return nil, fmt.Errorf("ops: asset %q listed twice", name)
		}
		command, err := cfg.Pack.command(name)
		if err != nil {
			return nil, err
		}
		if command.names(paramRepoURL) && len(cfg.ForgeHosts) == 0 {
			return nil, fmt.Errorf("ops: asset %q names %s: RunnerConfig.ForgeHosts is required", name, paramRepoURL)
		}
	}
	cfg.Assets = slices.Clone(cfg.Assets)
	cfg.ForgeHosts = slices.Clone(cfg.ForgeHosts)
	return &Runner{config: cfg}, nil
}

// Run claims and runs orders until ctx is done, waiting PollInterval after an empty claim.
func (r *Runner) Run(ctx context.Context) error {
	for {
		order, err := r.RunOnce(ctx)
		if err != nil {
			return err
		}
		if order != nil {
			continue
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(r.config.PollInterval):
		}
	}
}

// RunOnce claims and runs one order. It returns nil, nil when the queue has nothing for this runner.
func (r *Runner) RunOnce(ctx context.Context) (*traust.RunOrder, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	order, err := r.claim(ctx)
	if err != nil || order == nil {
		return nil, err
	}
	l, err := hold(r.config.Client, order, r.config.LeaseSeconds, r.config.HeartbeatInterval, r.config.RequestTimeout)
	if err != nil {
		return order, err
	}
	command, err := r.config.Pack.command(order.Asset)
	if err != nil || !slices.Contains(r.config.Assets, order.Asset) {
		return r.finish(ctx, l, refused(traust.RunOrderReasonAssetNotInPack))
	}
	if err := l.renew(ctx); err != nil {
		return l.snapshot(), err
	}
	pinned, err := supervise(ctx, l, func(read context.Context) (bool, error) { return r.pinned(read, l) })
	if err != nil {
		return l.snapshot(), err
	}
	if !pinned {
		return r.finish(ctx, l, refused(traust.RunOrderReasonBadRequest))
	}
	argv, err := supervise(ctx, l, func(read context.Context) ([]string, error) { return r.resolve(read, l, command) })
	var unresolved *unresolvedParam
	if errors.As(err, &unresolved) {
		return r.finish(ctx, l, refused(traust.RunOrderReasonBadRequest))
	}
	if err != nil {
		return l.snapshot(), err
	}
	if err := l.renew(ctx); err != nil {
		return l.snapshot(), err
	}
	produced, err := l.run(ctx, r.config.Executor, Execution{Argv: argv})
	if lost(err) || ctx.Err() != nil {
		return l.snapshot(), err
	}
	if err == nil {
		err = checkOutcome(produced.Result)
	}
	if err != nil {
		finished, finishErr := r.finish(ctx, l, traust.JobResult{Outcome: traust.OutcomeFailed, Reason: reason(traust.RunOrderReasonRunnerCrashed)})
		return finished, errors.Join(err, finishErr)
	}
	return r.finish(ctx, l, produced.Result)
}

func (r *Runner) claim(ctx context.Context) (*traust.RunOrder, error) {
	ctx, cancel := context.WithTimeout(ctx, r.config.RequestTimeout)
	defer cancel()
	assets := slices.Clone(r.config.Assets)
	order, err := r.config.Client.Claim(ctx, traust.ClaimRequest{Assets: &assets, LeaseSeconds: r.config.LeaseSeconds})
	if err != nil {
		return nil, fmt.Errorf("ops: claim: %w", err)
	}
	return order, nil
}

// pinned reports whether core planned the order with exactly this runner's pack. It runs under
// supervise, so a slow read keeps the lease and a lost lease cancels it.
func (r *Runner) pinned(ctx context.Context, l *lease) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.config.RequestTimeout)
	defer cancel()
	manifest, err := r.config.Client.GetOrderManifest(ctx, l.id())
	if err != nil {
		return false, fmt.Errorf("ops: order %s manifest: %w", l.id(), err)
	}
	if manifest.Config == nil {
		return false, nil
	}
	var digests []string
	for _, config := range *manifest.Config {
		if config.Name == r.config.Pack.ref() {
			digests = append(digests, config.Digest)
		}
	}
	return len(digests) == 1 && digests[0] == r.config.Pack.digest, nil
}

func (r *Runner) finish(ctx context.Context, l *lease, result traust.JobResult) (*traust.RunOrder, error) {
	completed, err := l.complete(ctx, result)
	order := l.latest(completed)
	return &order, err
}

// resolve fills the command's parameters from the order, its lease and the pack: the order's id
// and routing rule, the claiming worker, the asset's model role, and — whatever the command
// names — the repository URL of a repository-bound order, so no such order runs for a host off
// the allowlist.
func (r *Runner) resolve(ctx context.Context, l *lease, cmd command) ([]string, error) {
	order := l.order()
	values := map[string]string{paramRunID: l.id().String()}
	if order.LeaseHolder != nil {
		values[paramWorker] = *order.LeaseHolder
	}
	if role := r.config.Pack.role(order.Asset); role != "" {
		values[paramRole] = role
	}
	if order.Rule != nil {
		values[paramRule] = *order.Rule
	}
	if order.RepositoryId != nil {
		repoURL, err := r.repositoryURL(ctx, *order.RepositoryId)
		if err != nil {
			return nil, err
		}
		values[paramRepoURL] = repoURL
	}
	return cmd.fill(values)
}

// repositoryURL is the repository's URL, read through core: a repository the order names but
// core does not know, or whose URL core would not parse, cannot be resolved.
func (r *Runner) repositoryURL(ctx context.Context, id uuid.UUID) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, r.config.RequestTimeout)
	defer cancel()
	repository, err := r.config.Client.GetRepository(ctx, id)
	if err != nil {
		var problem *traust.ProblemError
		if errors.As(err, &problem) && problem.StatusCode == http.StatusNotFound {
			return "", &unresolvedParam{fmt.Errorf("repository %s: %w", id, err)}
		}
		return "", fmt.Errorf("ops: order repository %s: %w", id, err)
	}
	normalized, err := httpsRepoURL(repository.Url, r.config.ForgeHosts)
	if err != nil {
		return "", &unresolvedParam{err}
	}
	return normalized, nil
}

var (
	// A plain hostname, never path-like: a repository's host becomes a directory under a
	// mirror cache root.
	hostRE = regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9.-]*[A-Za-z0-9])?$`)
	// One path segment of a repository URL: an owner, group or repository name.
	segmentRE = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)
)

// httpsRepoURL is core's HttpsRepoUrl: an https repository URL, host and path only, its host
// on the allowlist. The normalized URL is what {repo_url} fills.
func httpsRepoURL(raw string, allow []string) (string, error) {
	if !isASCII(raw) || raw != strings.TrimSpace(raw) || strings.HasPrefix(raw, "-") {
		return "", fmt.Errorf("repo url has unsafe characters: %q", raw)
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("repo url is malformed: %q", raw)
	}
	if parsed.Scheme != "https" {
		return "", fmt.Errorf("repo url must be https: %q", raw)
	}
	if parsed.User != nil || parsed.Port() != "" || parsed.RawQuery != "" || parsed.Fragment != "" || parsed.RawFragment != "" {
		return "", fmt.Errorf("repo url must be host and path only: %q", raw)
	}
	host := strings.ToLower(parsed.Hostname())
	if !hostRE.MatchString(host) || strings.Contains(host, "..") {
		return "", fmt.Errorf("repo url host is invalid: %q", raw)
	}
	path := strings.TrimSuffix(strings.Trim(parsed.Path, "/"), ".git")
	segments := strings.Split(path, "/")
	if len(segments) < 2 {
		return "", fmt.Errorf("repo url path is invalid: %q", raw)
	}
	for _, segment := range segments {
		if segment == "." || segment == ".." || strings.HasPrefix(segment, "-") || !segmentRE.MatchString(segment) {
			return "", fmt.Errorf("repo url path is invalid: %q", raw)
		}
	}
	if !slices.ContainsFunc(allow, func(allowed string) bool { return strings.ToLower(allowed) == host }) {
		return "", fmt.Errorf("repo host not allowed: %q", host)
	}
	return "https://" + host + "/" + path, nil
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x80 {
			return false
		}
	}
	return true
}

func refused(why traust.RunOrderReason) traust.JobResult {
	return traust.JobResult{Outcome: traust.OutcomeRefused, Reason: reason(why)}
}

func reason(why traust.RunOrderReason) *traust.RunOrderReason { return &why }
