// Code generated from traust-contracts 2e521561f022fe839dd828ddf2185ba3a0a1045f. DO NOT EDIT.

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
