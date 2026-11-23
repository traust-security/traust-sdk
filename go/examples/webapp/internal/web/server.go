// Package web is what the app's features share: the server, the sign-in session, the helpers its
// pages answer with, and the layout they render in.
package web

import (
	"embed"
	"net/http"
)

//go:embed static
var static embed.FS

// Handler serves the app's static files and each feature's routes on one mux.
func Handler(routes ...func(*http.ServeMux)) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(static))
	for _, register := range routes {
		register(mux)
	}
	return mux
}
