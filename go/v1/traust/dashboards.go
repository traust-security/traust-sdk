package traust

// Group is the row's value for a dimension the query grouped by; nil when the row has none.
func (r DashboardRow) Group(d Dimension) *string { return r.Groups[string(d)] }

// Measure is the row's value for a measure the query asked for; nil when unmeasured.
func (r DashboardRow) Measure(m Measure) *float32 { return r.Measures[string(m)] }
