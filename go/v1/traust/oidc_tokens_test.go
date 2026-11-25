package traust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const testClientSecret = "s3cr3t-test-value"

type fakeClock struct{ now time.Time }

func (c *fakeClock) Now() time.Time { return c.now }

func (c *fakeClock) advance(d time.Duration) { c.now = c.now.Add(d) }

// issuerStub counts grants, checks the client's Basic auth and form, and
// answers with the given body and status. The secret stays test-only.
type issuerStub struct {
	grants   atomic.Int64
	server   *httptest.Server
	authSeen atomic.Value // string
	formSeen atomic.Value // url.Values
}

func newIssuerStub(t *testing.T, status int, body string) *issuerStub {
	t.Helper()
	stub := &issuerStub{}
	stub.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stub.grants.Add(1)
		id, secret, _ := r.BasicAuth()
		stub.authSeen.Store(id + ":" + secret)
		_ = r.ParseForm()
		stub.formSeen.Store(r.PostForm)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(stub.server.Close)
	return stub
}

// tokenBody is a well-formed client-credentials response.
func tokenBody(token string, expiresIn int) string {
	return fmt.Sprintf(`{"access_token":%q,"token_type":"Bearer","expires_in":%d}`, token, expiresIn)
}

func newSource(t *testing.T, stub *issuerStub, clock *fakeClock) *ClientCredentialsTokenSource {
	t.Helper()
	src, err := NewClientCredentialsTokenSource(ClientCredentialsTokenSource{
		TokenURL:     stub.server.URL,
		ClientID:     "svc-test",
		ClientSecret: testClientSecret,
		Now:          clock.Now,
	})
	if err != nil {
		t.Fatal(err)
	}
	return src
}

func TestNewClientCredentialsTokenSourceValidates(t *testing.T) {
	tests := []struct {
		name    string
		src     ClientCredentialsTokenSource
		wantErr string
	}{
		{name: "missing TokenURL", src: ClientCredentialsTokenSource{ClientID: "id", ClientSecret: "secret"}, wantErr: "TokenURL is required"},
		{name: "missing ClientID", src: ClientCredentialsTokenSource{TokenURL: "https://issuer.example/token", ClientSecret: "secret"}, wantErr: "ClientID is required"},
		{name: "missing ClientSecret", src: ClientCredentialsTokenSource{TokenURL: "https://issuer.example/token", ClientID: "id"}, wantErr: "ClientSecret is required"},
		{name: "plain http off loopback", src: ClientCredentialsTokenSource{TokenURL: "http://issuer.example/token", ClientID: "id", ClientSecret: "secret"}, wantErr: "must be https"},
		{name: "not a URL", src: ClientCredentialsTokenSource{TokenURL: "issuer.example", ClientID: "id", ClientSecret: "secret"}, wantErr: "not a valid URL"},
		{name: "userinfo in the URL", src: ClientCredentialsTokenSource{TokenURL: "https://u:p@issuer.example/token", ClientID: "id", ClientSecret: "secret"}, wantErr: "no userinfo or query"},
		{name: "query in the URL", src: ClientCredentialsTokenSource{TokenURL: "https://issuer.example/token?key=s3cret", ClientID: "id", ClientSecret: "secret"}, wantErr: "no userinfo or query"},
		{name: "https accepted", src: ClientCredentialsTokenSource{TokenURL: "https://issuer.example/token", ClientID: "id", ClientSecret: "secret"}},
		{name: "http on localhost", src: ClientCredentialsTokenSource{TokenURL: "http://localhost:8080/token", ClientID: "id", ClientSecret: "secret"}},
		{name: "http on loopback IP", src: ClientCredentialsTokenSource{TokenURL: "http://127.0.0.1:8080/token", ClientID: "id", ClientSecret: "secret"}},
		{name: "http on loopback IPv6", src: ClientCredentialsTokenSource{TokenURL: "http://[::1]:8080/token", ClientID: "id", ClientSecret: "secret"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, err := NewClientCredentialsTokenSource(tt.src)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if src == nil {
					t.Fatal("expected a token source")
				}
				return
			}
			if err == nil {
				t.Fatal("expected an error")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error %q does not name %q", err, tt.wantErr)
			}
			if strings.Contains(err.Error(), "s3cret") || strings.Contains(err.Error(), "u:p") {
				t.Fatalf("error %q repeats a credential from the URL", err)
			}
			if tt.src.ClientSecret != "" && strings.Contains(err.Error(), tt.src.ClientSecret) {
				t.Fatalf("error %q leaks the client secret", err)
			}
		})
	}
}

func TestClientCredentialsFetchesOnceAndCaches(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	stub := newIssuerStub(t, 200, tokenBody("tok-1", 3600))
	src := newSource(t, stub, clock)
	src.Scopes = []string{"reads:all", "writes:none"}

	for i := 0; i < 3; i++ {
		token, err := src.BearerToken(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if token != "tok-1" {
			t.Fatalf("got token %q, want %q", token, "tok-1")
		}
	}
	if got := stub.grants.Load(); got != 1 {
		t.Fatalf("%d grants, want 1: the token must be cached until just before expiry", got)
	}
	if got := stub.authSeen.Load().(string); got != url.QueryEscape("svc-test")+":"+url.QueryEscape(testClientSecret) {
		t.Fatalf("Basic auth %q, want the form-urlencoded client credentials (RFC 6749 §2.3.1)", got)
	}
	form := stub.formSeen.Load().(url.Values)
	if got := form.Get("grant_type"); got != "client_credentials" {
		t.Fatalf("grant_type %q, want client_credentials", got)
	}
	if got := form.Get("scope"); got != "reads:all writes:none" {
		t.Fatalf("scope %q, want the joined scopes", got)
	}
}

func TestClientCredentialsRefreshesOnlyPastItsLead(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	stub := newIssuerStub(t, 200, tokenBody("tok-1", 3600))
	src := newSource(t, stub, clock)

	if _, err := src.BearerToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	clock.advance(3600*time.Second - tokenRefreshLead - 1)
	if token, err := src.BearerToken(context.Background()); err != nil || token != "tok-1" {
		t.Fatalf("before the refresh point: token %q, err %v; want the cached token", token, err)
	}
	if got := stub.grants.Load(); got != 1 {
		t.Fatalf("%d grants before the refresh point, want 1", got)
	}

	clock.advance(1)
	if token, err := src.BearerToken(context.Background()); err != nil || token != "tok-1" {
		t.Fatalf("at the refresh point: token %q, err %v; want a refreshed token", token, err)
	}
	if got := stub.grants.Load(); got != 2 {
		t.Fatalf("%d grants after the refresh point, want 2", got)
	}
}

// A short-lived token refreshes after half its lifetime, not after the full
// lead that would overshoot the expiry.
func TestClientCredentialsRefreshesShortLivedTokensEarly(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	stub := newIssuerStub(t, 200, tokenBody("tok-1", 4))
	src := newSource(t, stub, clock)

	if _, err := src.BearerToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	clock.advance(2 * time.Second)
	if _, err := src.BearerToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got := stub.grants.Load(); got != 2 {
		t.Fatalf("%d grants, want 2: half of a 4s lifetime is the refresh point", got)
	}
}

func TestClientCredentialsConcurrentCallersShareOneFetch(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	stub := newIssuerStub(t, 200, tokenBody("tok-1", 3600))
	src := newSource(t, stub, clock)

	var wg sync.WaitGroup
	tokens := make([]string, 10)
	for i := range tokens {
		wg.Add(1)
		go func() {
			defer wg.Done()
			token, err := src.BearerToken(context.Background())
			if err != nil {
				t.Error(err)
				return
			}
			tokens[i] = token
		}()
	}
	wg.Wait()
	for i, token := range tokens {
		if token != "tok-1" {
			t.Fatalf("caller %d got token %q, want the shared one", i, token)
		}
	}
	if got := stub.grants.Load(); got != 1 {
		t.Fatalf("%d grants for 10 concurrent callers, want 1", got)
	}
}

func TestClientCredentialsFailuresAreTypedAndFailClosed(t *testing.T) {
	tests := []struct {
		name       string
		status     int
		body       string
		wantReason string
	}{
		{name: "bad request with OAuth error", status: 400, body: `{"error":"invalid_client"}`, wantReason: "issuer refused the grant"},
		{name: "server error", status: 500, body: "boom", wantReason: "issuer refused the grant"},
		{name: "malformed body", status: 200, body: `{"access_token":`, wantReason: "malformed token response"},
		{name: "missing access_token", status: 200, body: `{"token_type":"Bearer","expires_in":3600}`, wantReason: "without access_token"},
		{name: "empty access_token", status: 200, body: `{"access_token":"","token_type":"Bearer","expires_in":3600}`, wantReason: "without access_token"},
		{name: "non-bearer token_type", status: 200, body: `{"access_token":"t","token_type":"mac","expires_in":3600}`, wantReason: "non-bearer token_type"},
		{name: "token_type echoing the secret", status: 200, body: `{"access_token":"t","token_type":"` + testClientSecret + `","expires_in":3600}`, wantReason: "non-bearer token_type"},
		{name: "missing token_type", status: 200, body: `{"access_token":"t","expires_in":3600}`, wantReason: "non-bearer token_type"},
		{name: "missing expires_in", status: 200, body: `{"access_token":"t","token_type":"Bearer"}`, wantReason: "without expires_in"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
			stub := newIssuerStub(t, tt.status, tt.body)
			src := newSource(t, stub, clock)

			token, err := src.BearerToken(context.Background())
			if err == nil {
				t.Fatalf("got token %q, want a typed error", token)
			}
			var fetchErr *TokenFetchError
			if !errors.As(err, &fetchErr) {
				t.Fatalf("error %v is not a *TokenFetchError", err)
			}
			if !strings.Contains(fetchErr.Reason, tt.wantReason) {
				t.Fatalf("reason %q does not name %q", fetchErr.Reason, tt.wantReason)
			}
			if strings.Contains(err.Error(), testClientSecret) {
				t.Fatalf("error %q leaks the client secret", err)
			}
			if token := fetchErr.Unwrap(); token != nil && strings.Contains(token.Error(), testClientSecret) {
				t.Fatalf("wrapped error %q leaks the client secret", token)
			}
			if stub.grants.Load() != 1 {
				t.Fatal("a failed grant must not be cached as a token")
			}
		})
	}
}

// A canceled context stops the grant instead of hanging or guessing.
func TestClientCredentialsHonorsCancellation(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-release
		_, _ = w.Write([]byte(tokenBody("tok-1", 3600)))
	}))
	defer server.Close()
	defer close(release)
	src, err := NewClientCredentialsTokenSource(ClientCredentialsTokenSource{
		TokenURL:     server.URL,
		ClientID:     "svc-test",
		ClientSecret: testClientSecret,
		Now:          clock.Now,
	})
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := src.BearerToken(ctx); err == nil {
		t.Fatal("expected an error for a canceled context")
	}
}

// A grant the transport cannot deliver is an error, not an empty token.
func TestClientCredentialsTransportFailure(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	server.Close() // the issuer is gone
	src, err := NewClientCredentialsTokenSource(ClientCredentialsTokenSource{
		TokenURL:     server.URL,
		ClientID:     "svc-test",
		ClientSecret: testClientSecret,
		Now:          clock.Now,
		HTTPClient:   &http.Client{Timeout: 5 * time.Second},
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = src.BearerToken(context.Background())
	var fetchErr *TokenFetchError
	if !errors.As(err, &fetchErr) {
		t.Fatalf("error %v is not a *TokenFetchError", err)
	}
	if strings.Contains(err.Error(), testClientSecret) {
		t.Fatalf("error %q leaks the client secret", err)
	}
}

// A failed refresh drops the cache: the stale token is never handed out.
func TestClientCredentialsNeverReturnsStaleTokenAfterFailedRefresh(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	stub := newIssuerStub(t, 200, tokenBody("tok-1", 3600))
	src := newSource(t, stub, clock)

	if _, err := src.BearerToken(context.Background()); err != nil {
		t.Fatal(err)
	}
	stub.server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	})
	clock.advance(3600 * time.Second)
	token, err := src.BearerToken(context.Background())
	if err == nil {
		t.Fatalf("got token %q after a failed refresh, want an error", token)
	}
	if token != "" {
		t.Fatalf("got stale token %q, want none", token)
	}
}

// ClientCredentialsTokenSource satisfies the client's TokenSource.
func TestClientCredentialsIsUsableAsATokenSource(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	stub := newIssuerStub(t, 200, tokenBody("tok-1", 3600))
	src := newSource(t, stub, clock)

	var source TokenSource = src
	client, err := NewClient(Config{BaseURL: stub.server.URL, Tokens: source})
	if err != nil {
		t.Fatal(err)
	}
	var seen atomic.Value
	stub.server.Config.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost { // still the token endpoint
			stub.grants.Add(1)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(tokenBody("tok-1", 3600)))
			return
		}
		seen.Store(r.Header.Get("Authorization"))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"items":[],"next":null}`))
	})
	if _, err := client.Whoami(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got := seen.Load().(string); got != "Bearer tok-1" {
		t.Fatalf("Authorization %q, want Bearer tok-1", got)
	}
}

// The issuer's response body never reaches the error strings: a body that
// echoes the credentials cannot leak them through a refusal.
func TestClientCredentialsRefusalBodyStaysOutOfErrors(t *testing.T) {
	clock := &fakeClock{now: time.Unix(1_700_000_000, 0)}
	stub := newIssuerStub(t, 500, `{"error":"invalid_client","secret":"`+testClientSecret+`"}`)
	src := newSource(t, stub, clock)

	_, err := src.BearerToken(context.Background())
	if err == nil {
		t.Fatal("expected a refusal")
	}
	if strings.Contains(err.Error(), testClientSecret) {
		t.Fatalf("error %q leaks the client secret", err)
	}
	var fetchErr *TokenFetchError
	if !errors.As(err, &fetchErr) {
		t.Fatalf("error %v is not a *TokenFetchError", err)
	}
	if fetchErr.StatusCode != 500 {
		t.Fatalf("status %d, want 500", fetchErr.StatusCode)
	}
}
