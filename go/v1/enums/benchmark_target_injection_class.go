// Code generated from traust-contracts 83d90e75a562532845eaf4d13636e95f51c69134. DO NOT EDIT.

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
