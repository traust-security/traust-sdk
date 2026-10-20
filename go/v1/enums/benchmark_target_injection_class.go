// Code generated from traust-contracts ac517047c6566b019e63b911f6b7522000826b48. DO NOT EDIT.

package enums

// BenchmarkTargetInjectionClass — Values of `class` in benchmark-target.schema.json.
type BenchmarkTargetInjectionClass string

const (
	BenchmarkTargetInjectionClassReadmeBlatant    BenchmarkTargetInjectionClass = "readme_blatant"
	BenchmarkTargetInjectionClassCommentAuthority BenchmarkTargetInjectionClass = "comment_authority"
	BenchmarkTargetInjectionClassFileSuppression  BenchmarkTargetInjectionClass = "file_suppression"
	BenchmarkTargetInjectionClassSoftMisdirection BenchmarkTargetInjectionClass = "soft_misdirection"
	BenchmarkTargetInjectionClassReportShape      BenchmarkTargetInjectionClass = "report_shape"
	BenchmarkTargetInjectionClassHiddenText       BenchmarkTargetInjectionClass = "hidden_text"
)

// BenchmarkTargetInjectionClassValues returns all valid BenchmarkTargetInjectionClass values.
func BenchmarkTargetInjectionClassValues() []BenchmarkTargetInjectionClass {
	return []BenchmarkTargetInjectionClass{
		BenchmarkTargetInjectionClassReadmeBlatant,
		BenchmarkTargetInjectionClassCommentAuthority,
		BenchmarkTargetInjectionClassFileSuppression,
		BenchmarkTargetInjectionClassSoftMisdirection,
		BenchmarkTargetInjectionClassReportShape,
		BenchmarkTargetInjectionClassHiddenText,
	}
}
