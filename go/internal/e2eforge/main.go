// Command e2eforge is a one-repository GitLab-API stand-in for `make e2e`: the forge
// reads what core's onboarding needs (one project's visibility) and nothing else, so
// the e2e suite never leaves localhost.
//
// Core reaches a gitlab-kind forge at https://{host}/api/v4 — scheme and port are
// fixed, and a loopback bind of :443 needs privileges CI does not have. So the mock
// is instead an HTTP proxy set through HTTPS_PROXY: it answers CONNECT for any host
// with TLS served from an ephemeral self-signed certificate for that name, trusted
// through SSL_CERT_FILE (urllib verifies the origin through the tunnel), and answers
// the one GitLab projects endpoint onboarding reads. Every other request fails
// loudly, so a mock gap can never silently pass.
package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	caFile := flag.String("ca", "", "write the server certificate here for SSL_CERT_FILE")
	flag.Parse()

	cert, err := selfSigned()
	if err != nil {
		log.Fatal(err)
	}
	if *caFile != "" {
		if err := os.WriteFile(*caFile, certPEM(cert), 0o600); err != nil {
			log.Fatal(err)
		}
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", listener.Addr().(*net.TCPAddr).Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serve(conn, cert)
	}
}

// serve answers one proxied connection: CONNECT is always answered with a TLS session
// this mock signs, whose requests hit the forge handler; anything else is refused.
func serve(conn net.Conn, cert *tls.Certificate) {
	defer conn.Close()
	// The buffered reader can read ahead past the CONNECT request into the client's
	// TLS handshake bytes, so the tunnel must drain that buffer before the raw conn.
	buffered := bufio.NewReader(conn)
	request, err := http.ReadRequest(buffered)
	if err != nil {
		return
	}
	if request.Method != http.MethodConnect {
		fmt.Fprint(conn, "HTTP/1.1 403 e2eforge is a CONNECT proxy\r\n\r\n")
		return
	}
	fmt.Fprint(conn, "HTTP/1.1 200 Connected\r\n\r\n")
	tunnel := tls.Server(&drainedConn{Conn: conn, buffered: buffered}, &tls.Config{
		Certificates: []tls.Certificate{*cert},
	})
	if err := tunnel.Handshake(); err != nil {
		log.Printf("handshake: %v", err)
		return
	}
	http.Serve(newOneConnListener(tunnel), forgeHandler())
}

// drainedConn reads the bufio buffer first, then the connection it wrapped.
type drainedConn struct {
	net.Conn
	buffered *bufio.Reader
}

func (c *drainedConn) Read(p []byte) (int, error) {
	if c.buffered.Buffered() > 0 {
		return c.buffered.Read(p)
	}
	return c.Conn.Read(p)
}

// forgeHandler answers the GitLab REST reads core's CommandSourceProvider makes for
// onboarding: one project's info and default branch head.
func forgeHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v4/projects/", func(w http.ResponseWriter, r *http.Request) {
		// GitLab quotes the project path (owner%2Frepo); EscapedPath keeps it quoted.
		// Only the one fixture project exists; anything else is a mock gap.
		raw := strings.TrimPrefix(r.URL.EscapedPath(), "/api/v4/projects/")
		project, err := url.PathUnescape(raw)
		if err != nil || (project != "e2e/fixture" && project != "e2e/fixture/repository/branches/main") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "e2eforge: unknown project %s\n", raw)
			return
		}
		if project == "e2e/fixture/repository/branches/main" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"commit": map[string]string{"id": strings.Repeat("a", 40)},
			})
			return
		}
		// The GitLab project fields CommandSourceProvider.info() reads: last_activity_at,
		// archived, default_branch, visibility, forked_from_project.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"web_url":             "https://e2e-forge.test/e2e/fixture",
			"path_with_namespace": "e2e/fixture",
			"last_activity_at":    "2026-11-03T00:00:00Z",
			"archived":            false,
			"default_branch":      "main",
			"visibility":          "public",
			"forked_from_project": nil,
		})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Any other endpoint is a mock gap: fail loudly, never silently succeed.
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "e2eforge: no handler for %s\n", r.URL.Path)
	})
	return mux
}

// oneConnListener serves exactly the tunnelled connection, once. Accept never returns
// net.ErrClosed — http.Serve reads that as a shutdown and closes the connection it is
// serving — so after the handover it parks until the tunnel is done.
type oneConnListener struct {
	conn net.Conn
	once sync.Once
}

func newOneConnListener(conn net.Conn) *oneConnListener { return &oneConnListener{conn: conn} }

func (l *oneConnListener) Accept() (net.Conn, error) {
	var conn net.Conn
	l.once.Do(func() { conn = l.conn })
	if conn == nil {
		<-make(chan struct{}) // park: Serve re-accepts only after this conn is done
	}
	return conn, nil
}

func (l *oneConnListener) Close() error   { return nil }
func (l *oneConnListener) Addr() net.Addr { return l.conn.LocalAddr() }

// selfSigned makes a throwaway ECDSA P-256 leaf certificate — a real CA-signed shape,
// not a self-signed CA (rustls rejects CaUsedAsEndEntity) — valid for the hosts the
// mock might be asked to prove itself as, for a day. The leaf is its own trust root in
// SSL_CERT_FILE, which OpenSSL and urllib both accept for an end-entity certificate.
func selfSigned() (*tls.Certificate, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "e2eforge"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"e2e-forge.test", "pypi.org", "files.pythonhosted.org"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}
	der, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return nil, err
	}
	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, err
	}
	pair, err := tls.X509KeyPair(
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}),
		pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER}),
	)
	return &pair, err
}

func certPEM(cert *tls.Certificate) []byte {
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Certificate[0]})
}
