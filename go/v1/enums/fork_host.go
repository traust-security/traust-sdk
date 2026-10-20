// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

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
