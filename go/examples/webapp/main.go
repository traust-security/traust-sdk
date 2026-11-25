// Command webapp is a small web app on traust: onboard repositories, scan them, work their
// findings, watch the dashboard and build reports. One package per feature under internal/
// (identity, inventory, runs, findings, exposure, remediation, ledger, insight, web), wired
// by main.go, each feature importing its SDK domain package as core<domain>. It only serves
// the pages; the box's claim worker (go/Makefile's webapp target) runs every order the
// shipped pack declares.
//
//	go run . -core-url https://core.example.com
//
// It holds no identity of its own: every page calls core as the signed-in person (OnBehalfOf).
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/exposure"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/findings"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/identity"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/insight"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/inventory"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/ledger"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/remediation"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/runs"
	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

func main() {
	addr := flag.String("addr", "localhost:3000", "listen address")
	coreURL := flag.String("core-url", "", "traust-core's URL, e.g. https://core.example.com (required)")
	oidcIssuer := flag.String("oidc-issuer", "", "the OIDC issuer people sign in with; empty offers pasted-token sign-in only")
	oidcClientID := flag.String("oidc-client-id", "", "the app's client id at the issuer")
	oidcSecretFile := flag.String("oidc-client-secret-file", "", "a file holding the app's client secret; empty for a public client")
	oidcRedirect := flag.String("oidc-redirect-url", "", "where the issuer sends the person back, e.g. http://localhost:3000/sso/callback")
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	secret, err := clientSecret(*oidcSecretFile)
	if err != nil {
		log.Fatal(err)
	}
	sso, err := identity.NewSSO(*oidcIssuer, *oidcClientID, secret, *oidcRedirect)
	if err != nil {
		log.Fatal(err)
	}
	if err := run(ctx, *addr, *coreURL, sso); err != nil {
		log.Fatal(err)
	}
}

// clientSecret reads the OIDC client secret file, whole and trimmed; no file is a public client.
func clientSecret(path string) (string, error) {
	if path == "" {
		return "", nil
	}
	raw, err := os.ReadFile(path) // #nosec G304 -- the operator names this file with -oidc-client-secret-file
	if err != nil {
		return "", fmt.Errorf("reading the OIDC client secret: %w", err)
	}
	secret := strings.TrimSpace(string(raw))
	if secret == "" {
		return "", fmt.Errorf("the OIDC client secret file %s is empty", path)
	}
	return secret, nil
}

func run(ctx context.Context, addr, coreURL string, sso *identity.SSO) error {
	core, err := traust.NewClient(traust.Config{BaseURL: coreURL})
	if err != nil {
		return err
	}
	server := &http.Server{Addr: addr, Handler: routes(core, sso), ReadHeaderTimeout: 10 * time.Second}
	return serve(ctx, server)
}

// routes wires the web app: the shared server, then each feature's pages.
func routes(core *traust.Client, sso *identity.SSO) http.Handler {
	return web.Handler(
		identity.New(core, sso).Routes,
		inventory.New(core, sso.Link()).Routes,
		runs.New(core).Routes,
		findings.New(core).Routes,
		exposure.New(core).Routes,
		ledger.New(core).Routes,
		remediation.New(core).Routes,
		insight.New(core).Routes,
	)
}

// serve runs the pages until ctx ends or the server fails, then stops gracefully.
func serve(ctx context.Context, server *http.Server) error {
	go func() { <-ctx.Done(); _ = server.Shutdown(context.WithoutCancel(ctx)) }()
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
