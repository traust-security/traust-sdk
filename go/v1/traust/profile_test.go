package traust

import (
	"errors"
	"os"
	"path/filepath"
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

func TestLoadConfigConnectsFromAProfile(t *testing.T) {
	tests := []struct {
		name      string
		opts      []ProfileOption
		env       map[string]string
		file      string // config.yaml contents; "" writes none
		wantURL   string
		wantToken string // wantTokenFile, absolute
		wantErr   error
		wantKey   *ProfileKeyError
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
			name:    "an unknown profile fails closed",
			opts:    []ProfileOption{Profile("nope")},
			file:    profileConfig,
			wantErr: ErrUnknownProfile,
		},
		{
			name:    "a missing token_file names the profile and the key",
			opts:    []ProfileOption{Profile("broken")},
			file:    "profiles:\n  broken:\n    url: https://core.example.com\n",
			wantKey: &ProfileKeyError{Profile: "broken", Key: "token_file"},
		},
		{
			name:    "a missing url names the profile and the key",
			opts:    []ProfileOption{Profile("broken")},
			file:    "profiles:\n  broken:\n    token_file: /var/run/token\n",
			wantKey: &ProfileKeyError{Profile: "broken", Key: "url"},
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
		})
	}
}
