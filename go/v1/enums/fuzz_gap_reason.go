// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

package enums

// FuzzGapReason — Why a surface of a fuzz target was not fuzzed. A gap is never clean.
type FuzzGapReason string

const (
	FuzzGapReasonNoHarness            FuzzGapReason = "no_harness"
	FuzzGapReasonFuzzedUpstream       FuzzGapReason = "fuzzed_upstream"
	FuzzGapReasonUnsupportedLanguage  FuzzGapReason = "unsupported_language"
	FuzzGapReasonNeedsLiveEnvironment FuzzGapReason = "needs_live_environment"
	FuzzGapReasonTrustedInput         FuzzGapReason = "trusted_input"
	FuzzGapReasonHeavyDependencies    FuzzGapReason = "heavy_dependencies"
	FuzzGapReasonHarnessDefect        FuzzGapReason = "harness_defect"
	FuzzGapReasonBuildFailed          FuzzGapReason = "build_failed"
)

// FuzzGapReasonValues returns all valid FuzzGapReason values.
func FuzzGapReasonValues() []FuzzGapReason {
	return []FuzzGapReason{
		FuzzGapReasonNoHarness,
		FuzzGapReasonFuzzedUpstream,
		FuzzGapReasonUnsupportedLanguage,
		FuzzGapReasonNeedsLiveEnvironment,
		FuzzGapReasonTrustedInput,
		FuzzGapReasonHeavyDependencies,
		FuzzGapReasonHarnessDefect,
		FuzzGapReasonBuildFailed,
	}
}
