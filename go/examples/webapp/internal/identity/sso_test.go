package identity

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

// issuerAsks is what the test issuer was asked: the authorize redirect's query, the token
// endpoint's form.
type issuerAsks struct {
	authorize url.Values
	token     url.Values
}

// anIssuer is a small OIDC provider: discovery, an authorize endpoint that answers with a code,
// and a token endpoint whose answer the test words.
func anIssuer(t *testing.T, tokenStatus int, tokenAnswer string) (*httptest.Server, *issuerAsks) {
	t.Helper()
	asks := &issuerAsks{}
	var base string
	mux := http.NewServeMux()
	mux.HandleFunc("GET /.well-known/openid-configuration", func(w http.ResponseWriter, r *http.Request) {
		discovered := map[string]string{
			"issuer": base, "authorization_endpoint": base + "/authorize", "token_endpoint": base + "/token",
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(discovered)
	})
	mux.HandleFunc("GET /authorize", func(w http.ResponseWriter, r *http.Request) {
		asks.authorize = r.URL.Query()
		back := url.Values{"code": {"issuer-code"}, "state": {r.URL.Query().Get("state")}}
		http.Redirect(w, r, "/sso/callback?"+back.Encode(), http.StatusSeeOther)
	})
	mux.HandleFunc("POST /token", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		asks.token = r.PostForm
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(tokenStatus)
		_, _ = w.Write([]byte(tokenAnswer))
	})
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)
	base = server.URL
	return server, asks
}

// ssoApp is identity's routes with an issuer configured.
func ssoApp(core *traustest.Server, sso *SSO) http.Handler {
	mux := http.NewServeMux()
	New(core.Client(), sso).Routes(mux)
	return mux
}

// configured builds the app's SSO against the test issuer.
func configured(issuer *httptest.Server) *SSO {
	sso, err := NewSSO(issuer.URL, "webapp", "test-only-secret", "http://app.test/sso/callback")
	if err != nil {
		panic(err)
	}
	return sso
}

// ssoAsk makes one request carrying the sign-in request's cookie.
func ssoAsk(h http.Handler, path string, cookie *http.Cookie) *httptest.ResponseRecorder {
	r := httptest.NewRequest(http.MethodGet, path, nil)
	if cookie != nil {
		r.AddCookie(cookie)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

// requestCookie pulls one named cookie out of an answer.
func requestCookie(t *testing.T, w *httptest.ResponseRecorder, name string) *http.Cookie {
	t.Helper()
	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

// started sends the person to the issuer and returns their request's cookie with the state the
// redirect carried.
func started(t *testing.T, h http.Handler) (*http.Cookie, string) {
	t.Helper()
	w := ssoAsk(h, "/sso", nil)
	if w.Code != http.StatusSeeOther {
		t.Fatalf("starting SSO: %d %s, want a redirect to the issuer", w.Code, w.Body)
	}
	cookie := requestCookie(t, w, ssoCookie)
	if cookie == nil || !cookie.HttpOnly || cookie.SameSite != http.SameSiteLaxMode || cookie.Path != "/sso" {
		t.Fatalf("starting SSO set %v, want an HttpOnly Lax request cookie on /sso", cookie)
	}
	to, err := url.Parse(w.Header().Get("Location"))
	if err != nil {
		t.Fatalf("the issuer redirect %q is not a URL: %v", w.Header().Get("Location"), err)
	}
	return cookie, to.Query().Get("state")
}

func TestTheSignInPageOffersSSOOnlyWhenAnIssuerIsConfigured(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", forToken("ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	issuer, _ := anIssuer(t, http.StatusOK, `{"access_token":"ada"}`)

	offered := pasteToken(ssoApp(core, configured(issuer)), "mallory")
	if offered.Code != http.StatusUnauthorized || !strings.Contains(offered.Body.String(), `href="/sso"`) ||
		!strings.Contains(offered.Body.String(), "Sign in with SSO") {
		t.Fatalf("a refused token with SSO set: %d %s, want the sign-in page offering SSO", offered.Code, offered.Body)
	}
	plain := pasteToken(app(core), "mallory")
	if plain.Code != http.StatusUnauthorized || strings.Contains(plain.Body.String(), "Sign in with SSO") {
		t.Fatalf("a refused token without SSO: %d %s, want the sign-in page as today", plain.Code, plain.Body)
	}
}

// pasteToken posts a pasted token, as the sign-in form does.
func pasteToken(h http.Handler, token string) *httptest.ResponseRecorder {
	form := url.Values{"token": {token}}
	r := httptest.NewRequest(http.MethodPost, "/sign-in", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}

func TestSSOStartsAPKCEFlowWithState(t *testing.T) {
	core := traustest.New(t)
	issuer, asks := anIssuer(t, http.StatusOK, `{"access_token":"ada"}`)
	h := ssoApp(core, configured(issuer))

	w := ssoAsk(h, "/sso", nil)
	if w.Code != http.StatusSeeOther {
		t.Fatalf("starting SSO: %d %s, want a redirect to the issuer", w.Code, w.Body)
	}
	to, err := url.Parse(w.Header().Get("Location"))
	if err != nil || !strings.HasPrefix(to.String(), issuer.URL+"/authorize?") {
		t.Fatalf("the issuer redirect %q, want the discovered authorize endpoint", to)
	}
	cookie := requestCookie(t, w, ssoCookie)
	if cookie == nil {
		t.Fatal("starting SSO set no request cookie")
	}
	state, verifier, found := strings.Cut(cookie.Value, ".")
	if !found || state == "" || verifier == "" {
		t.Fatalf("the request cookie %q, want a state and a verifier", cookie.Value)
	}
	sum := sha256.Sum256([]byte(verifier))
	want := []string{
		"response_type=code", "client_id=webapp", "code_challenge_method=S256",
		"code_challenge=" + base64.RawURLEncoding.EncodeToString(sum[:]),
		"state=" + url.QueryEscape(state), "scope=" + url.QueryEscape("openid email"),
		"redirect_uri=" + url.QueryEscape("http://app.test/sso/callback"),
	}
	for _, part := range want {
		if !strings.Contains(to.RawQuery, part) {
			t.Fatalf("the authorize redirect %q lacks %q", to.RawQuery, part)
		}
	}
	if asks.token != nil {
		t.Fatal("starting SSO asked the issuer for a token")
	}
}

func TestTheCallbackSignsThePersonInWithTheTokenCoreAccepts(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", forToken("sso-ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	issuer, asks := anIssuer(t, http.StatusOK, `{"access_token":"sso-ada"}`)
	h := ssoApp(core, configured(issuer))

	cookie, state := started(t, h)
	back := url.Values{"code": {"issuer-code"}, "state": {state}}
	w := ssoAsk(h, "/sso/callback?"+back.Encode(), cookie)
	if w.Code != http.StatusSeeOther || w.Header().Get("Location") != "/" {
		t.Fatalf("the callback: %d %s, want a redirect home", w.Code, w.Body)
	}
	session := requestCookie(t, w, web.SessionCookie)
	if session == nil || session.Value != "sso-ada" || !session.HttpOnly || session.SameSite != http.SameSiteStrictMode {
		t.Fatalf("the session cookie %v, want the issuer's token stored like a pasted one", session)
	}
	if requestCookie(t, w, ssoCookie).MaxAge != -1 {
		t.Fatal("the callback kept the sign-in request's cookie")
	}
	_, verifier, _ := strings.Cut(cookie.Value, ".")
	if asks.token.Get("grant_type") != "authorization_code" || asks.token.Get("code") != "issuer-code" ||
		asks.token.Get("code_verifier") != verifier || asks.token.Get("client_id") != "webapp" ||
		asks.token.Get("client_secret") != "test-only-secret" {
		t.Fatalf("the token exchange asked %v, want the code proven by its verifier", asks.token)
	}
}

func TestTheIssuerRefusesThePerson(t *testing.T) {
	core := traustest.New(t)
	issuer, _ := anIssuer(t, http.StatusBadRequest, `{"error":"invalid_grant"}`)
	h := ssoApp(core, configured(issuer))
	cookie, state := started(t, h)

	back := url.Values{"code": {"issuer-code"}, "state": {state}}
	w := ssoAsk(h, "/sso/callback?"+back.Encode(), cookie)
	if w.Code != http.StatusBadGateway || !strings.Contains(w.Body.String(), "The sign-in provider refused that sign-in.") {
		t.Fatalf("a refused exchange: %d %s, want the sign-in page saying so", w.Code, w.Body)
	}
	if requestCookie(t, w, web.SessionCookie) != nil {
		t.Fatal("a refused exchange stored a session")
	}
	if requestCookie(t, w, ssoCookie).MaxAge != -1 {
		t.Fatal("a refused exchange kept the sign-in request's cookie")
	}
}

func TestAStateMismatchIsRefused(t *testing.T) {
	core := traustest.New(t)
	issuer, asks := anIssuer(t, http.StatusOK, `{"access_token":"sso-ada"}`)
	h := ssoApp(core, configured(issuer))
	cookie, _ := started(t, h)

	tests := []struct {
		name   string
		path   string
		cookie *http.Cookie
	}{
		{"another request's state", "/sso/callback?code=issuer-code&state=not-my-state", cookie},
		{"no state at all", "/sso/callback?code=issuer-code", cookie},
		{"no request behind the link", "/sso/callback?code=issuer-code&state=anything", nil},
		{"the issuer refused outright", "/sso/callback?error=access_denied", cookie},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := ssoAsk(h, tt.path, tt.cookie)
			if w.Code != http.StatusBadRequest || !strings.Contains(w.Body.String(), "Sign in") {
				t.Fatalf("%s: %d %s, want the sign-in page", tt.name, w.Code, w.Body)
			}
			if requestCookie(t, w, web.SessionCookie) != nil {
				t.Fatalf("%s: a mismatched request stored a session", tt.name)
			}
		})
	}
	if asks.token != nil {
		t.Fatal("a mismatched request still asked the issuer for a token")
	}
}

func TestCoreRefusalKeepsThePersonSignedOut(t *testing.T) {
	core := traustest.New(t)
	core.Handle("GET /v1/whoami", forToken("sso-ada", traustest.JSON(http.StatusOK, traust.Principal{Display: "Ada"})))
	issuer, _ := anIssuer(t, http.StatusOK, `{"access_token":"someone-else"}`)
	h := ssoApp(core, configured(issuer))

	cookie, state := started(t, h)
	back := url.Values{"code": {"issuer-code"}, "state": {state}}
	w := ssoAsk(h, "/sso/callback?"+back.Encode(), cookie)
	if w.Code != http.StatusUnauthorized || !strings.Contains(w.Body.String(), "traust refused that token.") {
		t.Fatalf("a token core refuses: %d %s, want the sign-in page saying so", w.Code, w.Body)
	}
	if requestCookie(t, w, web.SessionCookie) != nil {
		t.Fatal("a token core refuses stored a session")
	}
}

func TestWithoutAnIssuerTheAppBehavesAsToday(t *testing.T) {
	core := traustest.New(t)
	h := app(core)
	if w := ssoAsk(h, "/sso", nil); w.Code != http.StatusNotFound {
		t.Fatalf("SSO without an issuer: %d, want 404", w.Code)
	}
	if w := ssoAsk(h, "/sso/callback?code=issuer-code&state=anything", nil); w.Code != http.StatusNotFound {
		t.Fatalf("the SSO callback without an issuer: %d, want 404", w.Code)
	}
}

func TestAHalfConfiguredIssuerIsAStartupError(t *testing.T) {
	tests := []struct {
		name                               string
		issuer, clientID, secret, redirect string
		wantErr                            bool
	}{
		{"nothing configured", "", "", "", "", false},
		{"a client id without an issuer", "", "webapp", "", "", true},
		{"an issuer without a client id", "https://issuer.example", "", "", "http://app.test/sso/callback", true},
		{"an issuer without a redirect URL", "https://issuer.example", "webapp", "", "", true},
		{"a redirect URL that is not absolute", "https://issuer.example", "webapp", "", "/sso/callback", true},
		{"whole", "https://issuer.example", "webapp", "test-only-secret", "http://app.test/sso/callback", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sso, err := NewSSO(tt.issuer, tt.clientID, tt.secret, tt.redirect)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NewSSO(%q, %q, %q, %q) built %v, want an error", tt.issuer, tt.clientID, tt.secret, tt.redirect, sso)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewSSO(%q, %q, %q, %q): %v", tt.issuer, tt.clientID, tt.secret, tt.redirect, err)
			}
			if (sso == nil) != (tt.issuer == "") {
				t.Fatalf("NewSSO(%q, %q, %q, %q) built %v", tt.issuer, tt.clientID, tt.secret, tt.redirect, sso)
			}
		})
	}
}
