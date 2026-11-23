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

// Named failures of LoadConfig, for errors.Is.
var (
	// ErrNoConfigFile: the config file is needed to fill in the Config but missing.
	ErrNoConfigFile = errors.New("traust: no config file")
	// ErrNoProfile: nothing selected a profile — no Profile option, no TRAUST_PROFILE, no default.
	ErrNoProfile = errors.New("traust: no profile selected")
	// ErrUnknownProfile: the selected profile is not in the config file.
	ErrUnknownProfile = errors.New("traust: unknown profile")
)

// ProfileKeyError reports a profile in the config file that is missing its url or token_file.
type ProfileKeyError struct {
	Profile string
	Key     string
}

func (e *ProfileKeyError) Error() string {
	return fmt.Sprintf("traust: profile %q has no %s", e.Profile, e.Key)
}

// ProfileOption tunes which profile LoadConfig selects.
type ProfileOption func(*profileChoice)

// Profile selects the profile named in ~/.traust/config.yaml.
func Profile(name string) ProfileOption {
	return func(c *profileChoice) { c.name = name }
}

type profileChoice struct{ name string }

type profileEntry struct {
	URL       string `yaml:"url"`
	TokenFile string `yaml:"token_file"`
}

type profileFile struct {
	Default  string                  `yaml:"default"`
	Profiles map[string]profileEntry `yaml:"profiles"`
}

// LoadConfig builds a Config from a profile in ~/.traust/config.yaml: the one opts name (Profile),
// else TRAUST_PROFILE, else the file's default. TRAUST_URL and TRAUST_TOKEN_FILE override the
// profile, and with both set no file is needed. The token comes from FileTokenSource on the
// profile's token_file (a leading ~/ expanded against the home directory).
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
			if entry.TokenFile == "" {
				return Config{}, &ProfileKeyError{Profile: name, Key: "token_file"}
			}
			tokenFile = entry.TokenFile
		}
	}
	tokenFile, err := expandHome(tokenFile)
	if err != nil {
		return Config{}, err
	}
	return Config{BaseURL: url, Tokens: FileTokenSource{Path: tokenFile}}, nil
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
