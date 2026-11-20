package traust

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
)

// fakeOrders serves POST /v1/orders and GET /v1/orders/{id}; the order finishes on the
// finishOn-th read, and every request's Authorization header is recorded.
func fakeOrders(t *testing.T, finishOn int32) (*Client, func() []string) {
	t.Helper()
	id := uuid.New()
	var reads atomic.Int32
	var mu sync.Mutex
	var auth []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		auth = append(auth, r.Header.Get("Authorization"))
		mu.Unlock()
		order := RunOrder{Id: &id, Asset: "scan", State: RunOrderStateQueued}
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v1/orders":
		case r.Method == http.MethodGet && r.URL.Path == "/v1/orders/"+id.String():
			if reads.Add(1) >= finishOn {
				finished := time.Now()
				order.State, order.FinishedAt = RunOrderStateSucceeded, &finished
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(order)
	}))
	t.Cleanup(server.Close)
	client, err := NewClient(Config{BaseURL: server.URL, Tokens: StaticToken("machine")})
	if err != nil {
		t.Fatal(err)
	}
	return client, func() []string { mu.Lock(); defer mu.Unlock(); return append([]string(nil), auth...) }
}

func TestBeginSubmitWaitsForCoreToFinishTheOrder(t *testing.T) {
	client, _ := fakeOrders(t, 3)
	poller, err := client.BeginSubmit(t.Context(), SubmitRequest{Order: OrderRequest{Asset: "scan"}})
	if err != nil {
		t.Fatal(err)
	}
	order, err := poller.Wait(t.Context(), time.Millisecond)
	if err != nil || order.State != RunOrderStateSucceeded || order.Id.String() != poller.ResumeToken() {
		t.Fatalf("order %+v err %v, want the submitted order succeeded", order, err)
	}
}

func TestResumeOrderPicksUpFromTheToken(t *testing.T) {
	client, _ := fakeOrders(t, 1)
	submitted, err := client.BeginSubmit(t.Context(), SubmitRequest{Order: OrderRequest{Asset: "scan"}})
	if err != nil {
		t.Fatal(err)
	}
	resumed, err := client.ResumeOrder(submitted.ResumeToken())
	if err != nil {
		t.Fatal(err)
	}
	order, done, err := resumed.Poll(t.Context())
	if err != nil || !done || order.State != RunOrderStateSucceeded {
		t.Fatalf("order %+v done %v err %v, want finished", order, done, err)
	}
}

func TestPollerRefusesBadInput(t *testing.T) {
	client, _ := fakeOrders(t, 1)
	if _, err := client.ResumeOrder("not-an-order"); err == nil {
		t.Fatal("ResumeOrder accepted a token that names no order")
	}
	poller, err := client.ResumeOrder(uuid.NewString())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := poller.Wait(t.Context(), 0); err == nil {
		t.Fatal("Wait accepted a zero interval")
	}
	if _, err := poller.Wait(t.Context(), time.Millisecond); !errors.Is(err, ErrNotFound) {
		t.Fatalf("an unknown order: %v, want ErrNotFound", err)
	}
}

func TestWaitStopsWhenTheContextEnds(t *testing.T) {
	client, _ := fakeOrders(t, 1<<30)
	poller, err := client.BeginSubmit(t.Context(), SubmitRequest{Order: OrderRequest{Asset: "scan"}})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(t.Context(), 20*time.Millisecond)
	defer cancel()
	if _, err := poller.Wait(ctx, time.Millisecond); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("got %v, want the context's deadline", err)
	}
}

func TestOnBehalfOfSendsThePersonsToken(t *testing.T) {
	client, auth := fakeOrders(t, 1)
	if _, err := client.Submit(OnBehalfOf(t.Context(), "person"), SubmitRequest{}); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Submit(t.Context(), SubmitRequest{}); err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(auth(), ","); got != "Bearer person,Bearer machine" {
		t.Fatalf("Authorization %q, want the person's then the client's", got)
	}
	_, err := client.Submit(OnBehalfOf(t.Context(), " "), SubmitRequest{})
	if err == nil || len(auth()) != 2 {
		t.Fatalf("an empty person token: err %v after %d requests, want an error and nothing sent", err, len(auth()))
	}
}
