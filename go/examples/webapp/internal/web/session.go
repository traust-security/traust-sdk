package web

import (
	"net/http"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// The session cookie holds the signed-in person's traust token; a real app's OIDC login sets it.
const SessionCookie = "traust_token"

// SignedIn sends the handler's traust calls as the signed-in person.
func SignedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(SessionCookie)
		if err != nil {
			Render(w, r, http.StatusUnauthorized, Failure("Sign in first."))
			return
		}
		next(w, r.WithContext(traust.OnBehalfOf(r.Context(), cookie.Value)))
	}
}
