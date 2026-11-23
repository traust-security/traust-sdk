package insight

import "github.com/traust-security/traust-sdk/go/v1/traust"

// Group is the row's value for a dimension the query grouped by; nil when the row has none.
func Group(row traust.DashboardRow, d traust.Dimension) *string { return row.Groups[string(d)] }

// Measure is the row's value for a measure the query asked for; nil when unmeasured.
func Measure(row traust.DashboardRow, m traust.Measure) *float32 { return row.Measures[string(m)] }
