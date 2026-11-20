// Code generated from traust-contracts 15bded2e3e4b6c70bd1e20ef8df50c7b66b5ebc1. DO NOT EDIT.

package enums

// ReviewSurface — The seven code surfaces every review must answer for: where audits most often find nothing because no step of the review was pointed there.
type ReviewSurface string

const (
	ReviewSurfaceCredentialAcceptance   ReviewSurface = "credential-acceptance"
	ReviewSurfaceOwnershipResolution    ReviewSurface = "ownership-resolution"
	ReviewSurfaceSecurityFlags          ReviewSurface = "security-flags"
	ReviewSurfaceOutboundFetch          ReviewSurface = "outbound-fetch"
	ReviewSurfaceStoredContentIntegrity ReviewSurface = "stored-content-integrity"
	ReviewSurfaceScopeEnforcement       ReviewSurface = "scope-enforcement"
	ReviewSurfaceLlmIntegration         ReviewSurface = "llm-integration"
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
		ReviewSurfaceLlmIntegration,
	}
}
