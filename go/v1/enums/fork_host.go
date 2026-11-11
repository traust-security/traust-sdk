// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
