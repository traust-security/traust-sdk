// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// Ecosystem — Package ecosystem of a dependency or affected module, shared by impact-analysis and report findings. actions/docker/helm are manifest-level surfaces and rpm/deb/apk are distro ecosystems for C/C++; neither has a reachability tier.
type Ecosystem string

const (
	EcosystemGo      Ecosystem = "go"
	EcosystemNpm     Ecosystem = "npm"
	EcosystemPypi    Ecosystem = "pypi"
	EcosystemMaven   Ecosystem = "maven"
	EcosystemCargo   Ecosystem = "cargo"
	EcosystemRuby    Ecosystem = "ruby"
	EcosystemNuget   Ecosystem = "nuget"
	EcosystemActions Ecosystem = "actions"
	EcosystemDocker  Ecosystem = "docker"
	EcosystemHelm    Ecosystem = "helm"
	EcosystemRpm     Ecosystem = "rpm"
	EcosystemDeb     Ecosystem = "deb"
	EcosystemApk     Ecosystem = "apk"
)

// EcosystemValues returns all valid Ecosystem values.
func EcosystemValues() []Ecosystem {
	return []Ecosystem{
		EcosystemGo,
		EcosystemNpm,
		EcosystemPypi,
		EcosystemMaven,
		EcosystemCargo,
		EcosystemRuby,
		EcosystemNuget,
		EcosystemActions,
		EcosystemDocker,
		EcosystemHelm,
		EcosystemRpm,
		EcosystemDeb,
		EcosystemApk,
	}
}
