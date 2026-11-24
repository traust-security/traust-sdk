// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
