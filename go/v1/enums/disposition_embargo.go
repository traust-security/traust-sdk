// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

package enums

// DispositionEmbargo — Embargo handling assertion for a finding, under security team criteria (remote code execution, authentication bypass, privilege escalation, sensitive data exposure, network-accessible vulnerability). Orthogonal to the validity and resolution axes: an embargoed finding is still tracked on both, and an embargo assertion alone never changes finding state. Human-only, on the same identity-provider-verified footing as disposition.severity.
type DispositionEmbargo string

const (
	DispositionEmbargoRequired    DispositionEmbargo = "required"
	DispositionEmbargoActive      DispositionEmbargo = "active"
	DispositionEmbargoNotRequired DispositionEmbargo = "not_required"
	DispositionEmbargoUncertain   DispositionEmbargo = "uncertain"
)

// DispositionEmbargoValues returns all valid DispositionEmbargo values.
func DispositionEmbargoValues() []DispositionEmbargo {
	return []DispositionEmbargo{
		DispositionEmbargoRequired,
		DispositionEmbargoActive,
		DispositionEmbargoNotRequired,
		DispositionEmbargoUncertain,
	}
}
