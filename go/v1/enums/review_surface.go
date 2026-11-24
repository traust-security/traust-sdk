// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
