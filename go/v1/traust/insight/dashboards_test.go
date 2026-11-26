package insight_test

import (
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/insight"
)

// The typed keys read a dashboard row's group and measure; nil when the row has none.
func TestGroupAndMeasureReadARowByTypedKey(t *testing.T) {
	critical := "critical"
	count := float32(3)
	row := traust.DashboardRow{
		Groups:   map[string]*string{string(traust.DimensionSeverity): &critical},
		Measures: map[string]*float32{string(traust.MeasureFindings): &count},
	}
	if group := insight.Group(row, traust.DimensionSeverity); group == nil || *group != critical {
		t.Fatalf("group = %v, want %s", group, critical)
	}
	if group := insight.Group(row, traust.DimensionProduct); group != nil {
		t.Fatalf("an ungrouped dimension read %v, want nil", *group)
	}
	if measure := insight.Measure(row, traust.MeasureFindings); measure == nil || *measure != count {
		t.Fatalf("measure = %v, want %v", measure, count)
	}
	if measure := insight.Measure(row, traust.MeasureRepositories); measure != nil {
		t.Fatalf("an unmeasured measure read %v, want nil", *measure)
	}
}
