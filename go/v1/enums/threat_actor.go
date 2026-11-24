// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// ThreatActor — Weakest-position actor FIRST: scoring reads the first entry.
type ThreatActor string

const (
	ThreatActorRemoteUnauth    ThreatActor = "remote_unauth"
	ThreatActorRemoteAuth      ThreatActor = "remote_auth"
	ThreatActorAdjacentNetwork ThreatActor = "adjacent_network"
	ThreatActorLocalUser       ThreatActor = "local_user"
	ThreatActorLocalAdmin      ThreatActor = "local_admin"
	ThreatActorSupplyChain     ThreatActor = "supply_chain"
	ThreatActorInsider         ThreatActor = "insider"
)

// ThreatActorValues returns all valid ThreatActor values.
func ThreatActorValues() []ThreatActor {
	return []ThreatActor{
		ThreatActorRemoteUnauth,
		ThreatActorRemoteAuth,
		ThreatActorAdjacentNetwork,
		ThreatActorLocalUser,
		ThreatActorLocalAdmin,
		ThreatActorSupplyChain,
		ThreatActorInsider,
	}
}
