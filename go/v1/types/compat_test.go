package types_test

import (
	"encoding/json"
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/enums"
	"github.com/traust-security/traust-sdk/go/v1/types"
)

// Published shapes consumers already build on keep compiling and keep their wire values.
func TestPublishedShapesStillBuild(t *testing.T) {
	_ = types.AdapterResult{Findings: []types.ReportFinding{{}}}
	_ = types.ReportFinding{Locations: []types.Location{{}}}
	_ = types.IsolationReview{Gaps: []types.Gap{{}}}

	encoded, err := json.Marshal([]enums.RoadmapPriority{enums.RoadmapPriorityP0, enums.RoadmapPriorityP4})
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `["p0","p4"]` {
		t.Fatalf("published roadmap priorities changed their wire values: %s", encoded)
	}
}

func TestNullableSectionsDecodeToTheirTypes(t *testing.T) {
	var facts types.CodeFacts
	if err := json.Unmarshal([]byte(`{"symbols":[{"kind":"func","language":"go","line":3,"name":"main"}],"languages":null}`), &facts); err != nil {
		t.Fatal(err)
	}
	if len(facts.Symbols) != 1 || facts.Symbols[0].Name != "main" || facts.Languages != nil {
		t.Fatalf("code facts sections lost their types: %+v", facts)
	}
	var report types.FuzzReport
	if err := json.Unmarshal([]byte(`{"coverage":{"lines_covered":4,"lines_total":10}}`), &report); err != nil {
		t.Fatal(err)
	}
	if report.Coverage == nil || report.Coverage.LinesCovered != 4 {
		t.Fatalf("fuzz coverage lost its type: %+v", report.Coverage)
	}
	if err := json.Unmarshal([]byte(`{"coverage":null}`), &report); err != nil || report.Coverage != nil {
		t.Fatalf("null coverage: %v %+v", err, report.Coverage)
	}
}
