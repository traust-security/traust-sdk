package traust

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// The canonical config: a default, two profiles, one with a ~/ token file.
const profileConfig = `default: prod
profiles:
  prod:
    url: https://core.example.com
    token_file: ~/prod-token
  staging:
    url: https://staging.example.com
    token_file: /var/run/staging-token
`

// A machine identity: the client-credentials trio, optional scopes, the secret in a file.
const identityConfig = `profiles:
  ci:
    url: https://core.example.com
    token_url: https://issuer.example.com/token
    client_id: ci-runner
    client_secret_file: ~/client-secret
    scopes:
      - findings:read
      - findings:write
`

func TestLoadConfigConnectsFromAProfile(t *testing.T) {
	tests := []struct {
		name      string
		opts      []ProfileOption
		env       map[string]string
		file      string // config.yaml contents; "" writes none
		secret    string // contents of ~/client-secret; "" writes none
		wantURL   string
		wantToken string // wantTokenFile, absolute
		// wantCC: the wanted ClientCredentialsTokenSource fields.
		wantTokenURL     string
		wantClientID     string
		wantClientSecret string
		wantScopes       []string
		wantErr          error
		wantText         string // substring of the error message
		wantPath         string // path (under home) the error message must name
		wantKey          *ProfileKeyError
		wantCred         *ProfileCredentialError
	}{
		{
			name:      "both env vars set need no config file",
			env:       map[string]string{"TRAUST_URL": "https://env.example.com", "TRAUST_TOKEN_FILE": "/var/run/env-token"},
			wantURL:   "https://env.example.com",
			wantToken: "/var/run/env-token",
		},
		{
			name:      "Profile selects the named profile",
			opts:      []ProfileOption{Profile("staging")},
			file:      profileConfig,
			wantURL:   "https://staging.example.com",
			wantToken: "/var/run/staging-token",
		},
		{
			name:      "Profile expands a leading ~ in token_file",
			opts:      []ProfileOption{Profile("prod")},
			file:      profileConfig,
			wantURL:   "https://core.example.com",
			wantToken: "HOME/prod-token",
		},
		{
			name:      "TRAUST_PROFILE selects the profile",
			env:       map[string]string{"TRAUST_PROFILE": "staging"},
			file:      profileConfig,
			wantURL:   "https://staging.example.com",
			wantToken: "/var/run/staging-token",
		},
		{
			name:      "the file's default selects the profile",
			file:      profileConfig,
			wantURL:   "https://core.example.com",
			wantToken: "HOME/prod-token",
		},
		{
			name:      "TRAUST_URL overrides the profile's url",
			opts:      []ProfileOption{Profile("staging")},
			env:       map[string]string{"TRAUST_URL": "https://env.example.com"},
			file:      profileConfig,
			wantURL:   "https://env.example.com",
			wantToken: "/var/run/staging-token",
		},
		{
			name:      "TRAUST_TOKEN_FILE overrides the profile's token_file",
			opts:      []ProfileOption{Profile("staging")},
			env:       map[string]string{"TRAUST_TOKEN_FILE": "/var/run/env-token"},
			file:      profileConfig,
			wantURL:   "https://staging.example.com",
			wantToken: "/var/run/env-token",
		},
		{
			name:      "TRAUST_TOKEN_FILE overrides the profile's client identity",
			opts:      []ProfileOption{Profile("ci")},
			env:       map[string]string{"TRAUST_TOKEN_FILE": "/var/run/env-token"},
			file:      identityConfig,
			wantURL:   "https://core.example.com",
			wantToken: "/var/run/env-token",
		},
		{
			name:             "a machine-identity profile builds a ClientCredentialsTokenSource",
			opts:             []ProfileOption{Profile("ci")},
			file:             identityConfig,
			secret:           "ci-secret\n",
			wantURL:          "https://core.example.com",
			wantTokenURL:     "https://issuer.example.com/token",
			wantClientID:     "ci-runner",
			wantClientSecret: "ci-secret", // the trailing newline is trimmed
			wantScopes:       []string{"findings:read", "findings:write"},
		},
		{
			name:             "scopes are optional on a machine identity",
			opts:             []ProfileOption{Profile("ci")},
			file:             "profiles:\n  ci:\n    url: https://core.example.com\n    token_url: https://issuer.example.com/token\n    client_id: ci-runner\n    client_secret_file: ~/client-secret\n",
			secret:           "ci-secret",
			wantURL:          "https://core.example.com",
			wantTokenURL:     "https://issuer.example.com/token",
			wantClientID:     "ci-runner",
			wantClientSecret: "ci-secret",
		},
		{
			name:    "an unknown profile fails closed",
			opts:    []ProfileOption{Profile("nope")},
			file:    profileConfig,
			wantErr: ErrUnknownProfile,
		},
		{
			name:     "a profile with no credential names the profile",
			opts:     []ProfileOption{Profile("broken")},
			file:     "profiles:\n  broken:\n    url: https://core.example.com\n",
			wantCred: &ProfileCredentialError{Profile: "broken", Err: ErrProfileNoCredential},
		},
		{
			name:     "a profile with both credentials names the profile",
			opts:     []ProfileOption{Profile("both")},
			file:     "profiles:\n  both:\n    url: https://core.example.com\n    token_file: /var/run/token\n    token_url: https://issuer.example.com/token\n    client_id: ci-runner\n    client_secret_file: ~/client-secret\n",
			secret:   "ci-secret",
			wantCred: &ProfileCredentialError{Profile: "both", Err: ErrProfileBothCredentials},
		},
		{
			name:    "an incomplete client-credentials trio names the missing key",
			opts:    []ProfileOption{Profile("ci")},
			file:    "profiles:\n  ci:\n    url: https://core.example.com\n    token_url: https://issuer.example.com/token\n    client_id: ci-runner\n",
			wantKey: &ProfileKeyError{Profile: "ci", Key: "client_secret_file"},
		},
		{
			name:     "a missing client secret file names the path, never a secret",
			opts:     []ProfileOption{Profile("ci")},
			file:     "profiles:\n  ci:\n    url: https://core.example.com\n    token_url: https://issuer.example.com/token\n    client_id: ci-runner\n    client_secret_file: ~/client-secret\n",
			wantErr:  os.ErrNotExist,
			wantPath: "client-secret",
		},
		{
			name:     "an http token_url off loopback is refused",
			opts:     []ProfileOption{Profile("ci")},
			file:     "profiles:\n  ci:\n    url: https://core.example.com\n    token_url: http://issuer.example.com/token\n    client_id: ci-runner\n    client_secret_file: ~/client-secret\n",
			secret:   "ci-secret",
			wantText: "must be https",
		},
		{
			name: "an unknown key in the file fails closed",
			file: "default: prod\nprofiles:\n  prod:\n    url: https://core.example.com\n    token_file: /var/run/token\n    region: eu\n",
		},
		{
			name: "an unknown top-level key in the file fails closed",
			file: "default: prod\nendpoints: {}\nprofiles:\n  prod:\n    url: https://core.example.com\n    token_file: /var/run/token\n",
		},
		{
			name:    "a missing config file fails closed when env is half set",
			env:     map[string]string{"TRAUST_URL": "https://env.example.com"},
			wantErr: ErrNoConfigFile,
		},
		{
			name:    "a missing config file fails closed when nothing is set",
			wantErr: ErrNoConfigFile,
		},
		{
			name:    "no default and no name fails closed",
			file:    "profiles:\n  prod:\n    url: https://core.example.com\n    token_file: /var/run/token\n",
			wantErr: ErrNoProfile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home := t.TempDir()
			t.Setenv("HOME", home)
			for _, key := range []string{"TRAUST_URL", "TRAUST_TOKEN_FILE", "TRAUST_PROFILE"} {
				t.Setenv(key, "")
			}
			for key, value := range tt.env {
				t.Setenv(key, value)
			}
			if tt.file != "" {
				dir := filepath.Join(home, ".traust")
				if err := os.MkdirAll(dir, 0o700); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(dir, "config.yaml"), []byte(tt.file), 0o600); err != nil {
					t.Fatal(err)
				}
			}
			if tt.secret != "" {
				if err := os.WriteFile(filepath.Join(home, "client-secret"), []byte(tt.secret), 0o600); err != nil {
					t.Fatal(err)
				}
			}
			cfg, err := LoadConfig(tt.opts...)
			switch {
			case tt.wantErr != nil:
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("LoadConfig() error = %v, want %v", err, tt.wantErr)
				}
			case tt.wantKey != nil:
				var keyErr *ProfileKeyError
				if !errors.As(err, &keyErr) {
					t.Fatalf("LoadConfig() error = %v, want %v", err, tt.wantKey)
				}
				if *keyErr != *tt.wantKey {
					t.Errorf("LoadConfig() key error = %+v, want %+v", *keyErr, *tt.wantKey)
				}
			case tt.wantCred != nil:
				var credErr *ProfileCredentialError
				if !errors.As(err, &credErr) {
					t.Fatalf("LoadConfig() error = %v, want %v", err, tt.wantCred)
				}
				if *credErr != *tt.wantCred {
					t.Errorf("LoadConfig() credential error = %+v, want %+v", *credErr, *tt.wantCred)
				}
			case tt.wantText != "":
				if err == nil || !strings.Contains(err.Error(), tt.wantText) {
					t.Fatalf("LoadConfig() error = %v, want it to mention %q", err, tt.wantText)
				}
			case tt.wantTokenURL != "":
				if err != nil {
					t.Fatalf("LoadConfig() error = %v", err)
				}
				if cfg.BaseURL != tt.wantURL {
					t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, tt.wantURL)
				}
				cc, ok := cfg.Tokens.(*ClientCredentialsTokenSource)
				if !ok {
					t.Fatalf("Tokens = %T, want *ClientCredentialsTokenSource", cfg.Tokens)
				}
				if cc.TokenURL != tt.wantTokenURL || cc.ClientID != tt.wantClientID || cc.ClientSecret != tt.wantClientSecret {
					t.Errorf("token source = {%s %s %q}, want {%s %s %q}",
						cc.TokenURL, cc.ClientID, cc.ClientSecret, tt.wantTokenURL, tt.wantClientID, tt.wantClientSecret)
				}
				if !slices.Equal(cc.Scopes, tt.wantScopes) {
					t.Errorf("scopes = %v, want %v", cc.Scopes, tt.wantScopes)
				}
				if cc.cache == nil {
					t.Error("token source was not built with NewClientCredentialsTokenSource")
				}
			case tt.wantURL == "" && tt.wantToken == "":
				if err == nil {
					t.Fatalf("LoadConfig() = %+v, want an error", cfg)
				}
			default:
				if err != nil {
					t.Fatalf("LoadConfig() error = %v", err)
				}
				if cfg.BaseURL != tt.wantURL {
					t.Errorf("BaseURL = %q, want %q", cfg.BaseURL, tt.wantURL)
				}
				tokens, ok := cfg.Tokens.(FileTokenSource)
				if !ok {
					t.Fatalf("Tokens = %T, want FileTokenSource", cfg.Tokens)
				}
				want := tt.wantToken
				if strings.HasPrefix(want, "HOME/") {
					want = filepath.Join(home, strings.TrimPrefix(want, "HOME/"))
				}
				if tokens.Path != want {
					t.Errorf("token file = %q, want %q", tokens.Path, want)
				}
			}
			if tt.wantPath != "" && (err == nil || !strings.Contains(err.Error(), filepath.Join(home, tt.wantPath))) {
				t.Errorf("LoadConfig() error = %v, want it to name %s", err, filepath.Join(home, tt.wantPath))
			}
		})
	}
}
