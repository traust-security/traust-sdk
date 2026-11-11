// Package ledger provides historical reads and raw and human-action writes for
// the peer ledger service. It is transport-agnostic through Provider;
// NewHTTPClient supplies the built-in HTTP transport.
//
// Machine reports are recorded through the generated v1/traust native recording
// endpoints or v1/ops.RecordingRun, using native IDs, owned run orders and the
// original report bytes. Producers must pass their canonical final gate before
// recording; this package does not convert reports into disposition events.
//
//	client := ledger.NewHTTPClient("https://ledger.example.com",
//		ledger.WithBearerToken(os.Getenv("LEDGER_TOKEN")),
//	)
//	findings, err := client.GetFindings(ctx, "repo-a")
package ledger
