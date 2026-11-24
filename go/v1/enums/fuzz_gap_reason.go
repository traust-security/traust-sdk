// Code generated from traust-contracts 369a2ab2c9b65225c6878eca75bada6a65df548f. DO NOT EDIT.

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
