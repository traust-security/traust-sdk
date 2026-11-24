// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ForkHost — Values of `host` in remediation.schema.json.
type ForkHost string

const (
	ForkHostGithub ForkHost = "github"
	ForkHostGitlab ForkHost = "gitlab"
	ForkHostOther  ForkHost = "other"
)

// ForkHostValues returns all valid ForkHost values.
func ForkHostValues() []ForkHost {
	return []ForkHost{
		ForkHostGithub,
		ForkHostGitlab,
		ForkHostOther,
	}
}
