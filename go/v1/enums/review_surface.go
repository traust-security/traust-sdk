// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// ReviewSurface — The six code surfaces every review must answer for: where audits most often find nothing because no step of the review was pointed there.
type ReviewSurface string

const (
	ReviewSurfaceCredentialAcceptance   ReviewSurface = "credential-acceptance"
	ReviewSurfaceOwnershipResolution    ReviewSurface = "ownership-resolution"
	ReviewSurfaceSecurityFlags          ReviewSurface = "security-flags"
	ReviewSurfaceOutboundFetch          ReviewSurface = "outbound-fetch"
	ReviewSurfaceStoredContentIntegrity ReviewSurface = "stored-content-integrity"
	ReviewSurfaceScopeEnforcement       ReviewSurface = "scope-enforcement"
)

// ReviewSurfaceValues returns all valid ReviewSurface values.
func ReviewSurfaceValues() []ReviewSurface {
	return []ReviewSurface{
		ReviewSurfaceCredentialAcceptance,
		ReviewSurfaceOwnershipResolution,
		ReviewSurfaceSecurityFlags,
		ReviewSurfaceOutboundFetch,
		ReviewSurfaceStoredContentIntegrity,
		ReviewSurfaceScopeEnforcement,
	}
}
