// Code generated from traust-contracts 8f684d27d978e53c0ffc2b43af379fc1284c26eb. DO NOT EDIT.

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
