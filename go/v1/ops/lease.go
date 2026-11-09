package ops

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// lease is an order this process holds. Every call after claim or start goes through it, and
// origin keeps the server's id and start time, which no heartbeat may change.
type lease struct {
	client   *traust.Client
	origin   traust.RunOrder
	current  traust.RunOrder
	seconds  int
	interval time.Duration
	timeout  time.Duration
}

func hold(client *traust.Client, order *traust.RunOrder, seconds int, interval, timeout time.Duration) (*lease, error) {
	if err := live(order); err != nil {
		return nil, &LeaseLostError{Err: err}
	}
	return &lease{client: client, origin: *order, current: *order, seconds: seconds, interval: interval, timeout: timeout}, nil
}

func live(order *traust.RunOrder) error {
	if order == nil || order.State != traust.RunOrderStateLeased ||
		order.Id == nil || *order.Id == uuid.Nil ||
		order.StartedAt == nil || order.StartedAt.IsZero() ||
		order.LeaseExpiresAt == nil || !order.LeaseExpiresAt.After(time.Now()) {
		return errors.New("core returned no live order lease")
	}
	return nil
}

func (l *lease) id() uuid.UUID { return *l.origin.Id }

// order is the server's original order with its own copies of the pointers an executor could
// otherwise use to rewrite the lease.
func (l *lease) order() traust.RunOrder {
	order := l.origin
	id, started, expires := *l.origin.Id, *l.origin.StartedAt, *l.origin.LeaseExpiresAt
	order.Id, order.StartedAt, order.LeaseExpiresAt = &id, &started, &expires
	if l.origin.WorkItemId != nil {
		workItem := *l.origin.WorkItemId
		order.WorkItemId = &workItem
	}
	return order
}

// check fails with lease loss once the lease has expired or the caller has given up.
func (l *lease) check(ctx context.Context) error {
	if err := live(&l.current); err != nil {
		return &LeaseLostError{Err: err}
	}
	if err := ctx.Err(); err != nil {
		return &LeaseLostError{Err: err}
	}
	return nil
}

// call bounds one mutation by the request timeout and the lease; a dead lease makes no call.
func (l *lease) call(ctx context.Context) (context.Context, context.CancelFunc, error) {
	if err := l.check(ctx); err != nil {
		return nil, nil, err
	}
	deadline := time.Now().Add(l.timeout)
	if l.current.LeaseExpiresAt.Before(deadline) {
		deadline = *l.current.LeaseExpiresAt
	}
	ctx, cancel := context.WithDeadline(ctx, deadline)
	return ctx, cancel, nil
}

// renew heartbeats the lease. Any failure is lease loss; a heartbeat is never retried.
func (l *lease) renew(ctx context.Context) error {
	ctx, cancel, err := l.call(ctx)
	if err != nil {
		return err
	}
	defer cancel()
	order, err := l.client.Heartbeat(ctx, l.id(), traust.LeaseRequest{LeaseSeconds: l.seconds})
	if err == nil {
		err = live(order)
	}
	if err == nil && (*order.Id != l.id() || !order.StartedAt.Equal(*l.origin.StartedAt)) {
		err = errors.New("heartbeat changed the order's id or start time")
	}
	if err != nil {
		return &LeaseLostError{Err: err}
	}
	l.current = *order
	return nil
}

// run calls the executor under supervise, with its own copy of the order.
func (l *lease) run(ctx context.Context, executor Executor, execution Execution) (ProducedReport, error) {
	if err := l.check(ctx); err != nil {
		return ProducedReport{}, err
	}
	execution.Order = l.order()
	return supervise(ctx, l, func(work context.Context) (ProducedReport, error) {
		return executor.Execute(work, execution)
	})
}

// supervise runs work while renewing the lease. Losing the lease cancels work and waits for it to
// stop before returning, so nothing outlives the lease it ran under.
func supervise[T any](ctx context.Context, l *lease, work func(context.Context) (T, error)) (T, error) {
	var zero T
	if err := l.check(ctx); err != nil {
		return zero, err
	}
	workCtx, stop := context.WithCancel(ctx)
	defer stop()
	type outcome struct {
		value T
		err   error
	}
	done := make(chan outcome, 1)
	go func() {
		value, err := work(workCtx)
		done <- outcome{value, err}
	}()
	abort := func(cause error) (T, error) {
		stop()
		<-done
		if lost(cause) {
			return zero, cause
		}
		return zero, &LeaseLostError{Err: cause}
	}
	heartbeat := time.NewTicker(l.interval)
	defer heartbeat.Stop()
	expiry := time.NewTimer(time.Until(*l.current.LeaseExpiresAt))
	defer expiry.Stop()
	for {
		select {
		case <-ctx.Done():
			return abort(ctx.Err())
		case <-expiry.C:
			return abort(errors.New("order lease expired"))
		case <-heartbeat.C:
			if err := l.renew(ctx); err != nil {
				return abort(err)
			}
			expiry.Reset(time.Until(*l.current.LeaseExpiresAt))
		case finished := <-done:
			if err := l.check(ctx); err != nil {
				return zero, err
			}
			return finished.value, finished.err
		}
	}
}

// superviseRequest renews during one bounded call without racing its deadline against renewals.
func superviseRequest[T any](ctx context.Context, l *lease, work func(context.Context) (T, error)) (T, error) {
	deadline := time.Now().Add(l.timeout)
	if l.current.LeaseExpiresAt.Before(deadline) {
		deadline = *l.current.LeaseExpiresAt
	}
	return supervise(ctx, l, func(workCtx context.Context) (T, error) {
		requestCtx, cancel := context.WithDeadline(workCtx, deadline)
		defer cancel()
		return work(requestCtx)
	})
}

// complete renews the lease, then finishes the order with result.
func (l *lease) complete(ctx context.Context, result traust.JobResult) (*traust.RunOrder, error) {
	if err := l.renew(ctx); err != nil {
		return nil, err
	}
	ctx, cancel, err := l.call(ctx)
	if err != nil {
		return nil, err
	}
	defer cancel()
	return l.client.Complete(ctx, l.id(), traust.CompleteRequest{Result: result})
}

// snapshot is a copy of the lease's current view of the order.
func (l *lease) snapshot() *traust.RunOrder {
	order := l.current
	return &order
}

// latest is the completed order when there is one, else the lease's current view of it.
func (l *lease) latest(completed *traust.RunOrder) traust.RunOrder {
	if completed != nil {
		return *completed
	}
	return l.current
}
