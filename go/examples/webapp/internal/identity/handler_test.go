package identity

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

// app is identity's routes alone: the layout's two asks, no feature pages.
func app(core *traustest.Server) http.Handler {
	mux := http.NewServeMux()
	New(core.Client()).Routes(mux)
	return mux
}

// ask makes one request, signed in when token is set.
func ask(h http.Handler, method, path, token string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, nil)
	if token != "" {
		r.AddCookie(&http.Cookie{Name: web.SessionCookie, Value: token})
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

// forToken answers requests carrying token, and refuses every other caller, as core would.
func forToken(token string, answer http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+token {
			traustest.Problem(http.StatusUnauthorized, "unknown caller").ServeHTTP(w, r)
			return
		}
		answer.ServeHTTP(w, r)
	})
}

func TestTheHeaderNamesWhoCoreSaysIsSignedIn(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", forToken("ada", traustest.JSON(http.StatusOK, traust.Principal{
		Display: "Ada", Roles: &[]traust.Role{traust.Admin, traust.Restate},
	})))

	if w := ask(app(core), http.MethodGet, "/whoami", "ada"); w.Code != http.StatusOK ||
		!strings.Contains(w.Body.String(), "Signed in as Ada") || !strings.Contains(w.Body.String(), "(admin, restate)") {
		t.Fatalf("whoami: %d %s, want the person named with her roles", w.Code, w.Body)
	}
	if w := ask(app(core), http.MethodGet, "/whoami", ""); w.Code != http.StatusUnauthorized {
		t.Fatalf("without a session: %d, want 401", w.Code)
	}
}

func TestARefusedWhoamiKeepsItsStatus(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", traustest.Problem(http.StatusForbidden, "no such caller"))
	if w := ask(app(core), http.MethodGet, "/whoami", "ada"); w.Code != http.StatusForbidden {
		t.Fatalf("refused whoami: %d, want core's 403 kept", w.Code)
	}
}

func TestTheFooterWordsCoreHealth(t *testing.T) {
	healthy := traustest.New(t)
	healthy.Handle("GET /v1/healthz", traustest.JSON(http.StatusOK, traust.Health{Status: "ok"}))
	if w := ask(app(healthy), http.MethodGet, "/health", ""); w.Code != http.StatusOK ||
		!strings.Contains(w.Body.String(), "Core: ok") {
		t.Fatalf("a healthy core: %d %s", w.Code, w.Body)
	}
	refused := traustest.New(t) // no healthz handler: core refuses every ask
	if w := ask(app(refused), http.MethodGet, "/health", ""); w.Code != http.StatusOK ||
		!strings.Contains(w.Body.String(), "Core: unreachable") {
		t.Fatalf("a refusing core: %d %s, want it said unreachable, not healthy", w.Code, w.Body)
	}
}
