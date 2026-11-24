// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
