package traust

// LoadConfig is the SDK's one read of the environment and the user's home directory: the explicit
// loader a consumer calls to connect from a profile name (AGENTS.md "One config").

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ProfileOption tunes which profile LoadConfig selects.
type ProfileOption func(*profileChoice)

type profileChoice struct{ name string }

type profileEntry struct {
	URL              string   `yaml:"url"`
	TokenFile        string   `yaml:"token_file"`
	TokenURL         string   `yaml:"token_url"`
	ClientID         string   `yaml:"client_id"`
	ClientSecretFile string   `yaml:"client_secret_file"`
	Scopes           []string `yaml:"scopes"`
}

type profileFile struct {
	Default  string                  `yaml:"default"`
	Profiles map[string]profileEntry `yaml:"profiles"`
}

// Profile selects the profile named in ~/.traust/config.yaml.
func Profile(name string) ProfileOption {
	return func(c *profileChoice) { c.name = name }
}

// LoadConfig builds a Config from a profile in ~/.traust/config.yaml: the one opts name (Profile),
// else TRAUST_PROFILE, else the file's default. TRAUST_URL and TRAUST_TOKEN_FILE override the
// profile, and with both set no file is needed; TRAUST_TOKEN_FILE selects the token-file
// credential. The profile names exactly one credential: a token_file (a leading ~/ expanded,
// served by FileTokenSource so rotated tokens just work), or a machine identity — token_url,
// client_id and client_secret_file, with optional scopes — served by ClientCredentialsTokenSource.
// The client secret is read from its file here (a leading ~/ expanded, a trailing newline
// trimmed), never from the config file itself.
func LoadConfig(opts ...ProfileOption) (Config, error) {
	var choice profileChoice
	for _, opt := range opts {
		opt(&choice)
	}
	url, tokenFile := os.Getenv("TRAUST_URL"), os.Getenv("TRAUST_TOKEN_FILE")
	if url == "" || tokenFile == "" {
		file, path, err := readProfileFile()
		if err != nil {
			return Config{}, err
		}
		name := choice.name
		if name == "" {
			name = os.Getenv("TRAUST_PROFILE")
		}
		if name == "" {
			name = file.Default
		}
		if name == "" {
			return Config{}, fmt.Errorf("%w: %s has no default", ErrNoProfile, path)
		}
		entry, known := file.Profiles[name]
		if !known {
			return Config{}, fmt.Errorf("%w: %q in %s", ErrUnknownProfile, name, path)
		}
		if url == "" {
			if entry.URL == "" {
				return Config{}, &ProfileKeyError{Profile: name, Key: "url"}
			}
			url = entry.URL
		}
		if tokenFile == "" {
			tokens, err := profileTokens(name, entry)
			if err != nil {
				return Config{}, err
			}
			return Config{BaseURL: url, Tokens: tokens}, nil
		}
	}
	tokenFile, err := expandHome(tokenFile)
	if err != nil {
		return Config{}, err
	}
	return Config{BaseURL: url, Tokens: FileTokenSource{Path: tokenFile}}, nil
}

// profileTokens builds the profile's one credential: token_file, or the client-credentials trio
// (token_url, client_id, client_secret_file) built with NewClientCredentialsTokenSource so its
// validation applies.
func profileTokens(name string, entry profileEntry) (TokenSource, error) {
	hasTrio := entry.TokenURL != "" || entry.ClientID != "" || entry.ClientSecretFile != ""
	switch {
	case entry.TokenFile != "" && hasTrio:
		return nil, &ProfileCredentialError{Profile: name, Err: ErrProfileBothCredentials}
	case entry.TokenFile != "":
		path, err := expandHome(entry.TokenFile)
		if err != nil {
			return nil, err
		}
		return FileTokenSource{Path: path}, nil
	case !hasTrio:
		return nil, &ProfileCredentialError{Profile: name, Err: ErrProfileNoCredential}
	case entry.TokenURL == "":
		return nil, &ProfileKeyError{Profile: name, Key: "token_url"}
	case entry.ClientID == "":
		return nil, &ProfileKeyError{Profile: name, Key: "client_id"}
	case entry.ClientSecretFile == "":
		return nil, &ProfileKeyError{Profile: name, Key: "client_secret_file"}
	}
	secretPath, err := expandHome(entry.ClientSecretFile)
	if err != nil {
		return nil, err
	}
	raw, err := os.ReadFile(secretPath) // #nosec G304 -- the path is the profile's client_secret_file
	if err != nil {
		return nil, fmt.Errorf("traust: reading the client secret: %w", err)
	}
	return NewClientCredentialsTokenSource(ClientCredentialsTokenSource{
		TokenURL:     entry.TokenURL,
		ClientID:     entry.ClientID,
		ClientSecret: strings.TrimRight(string(raw), "\r\n"),
		Scopes:       entry.Scopes,
	})
}

// readProfileFile decodes ~/.traust/config.yaml strictly: a key it does not know is an error.
func readProfileFile() (profileFile, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return profileFile{}, "", fmt.Errorf("traust: locating the home directory: %w", err)
	}
	path := filepath.Join(home, ".traust", "config.yaml")
	raw, err := os.ReadFile(path) // #nosec G304 -- the path is the user's own ~/.traust/config.yaml
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return profileFile{}, path, fmt.Errorf("%w: %s", ErrNoConfigFile, path)
		}
		return profileFile{}, path, fmt.Errorf("traust: reading %s: %w", path, err)
	}
	var file profileFile
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true)
	if err := dec.Decode(&file); err != nil && !errors.Is(err, io.EOF) {
		return profileFile{}, path, fmt.Errorf("traust: parsing %s: %w", path, err)
	}
	return file, path, nil
}

// expandHome resolves a leading ~/ against the home directory.
func expandHome(path string) (string, error) {
	if !strings.HasPrefix(path, "~/") {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("traust: expanding ~ in %s: %w", path, err)
	}
	return filepath.Join(home, path[2:]), nil
}
