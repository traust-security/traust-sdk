package findings_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/findings"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

func ExampleClient_SummarizeFindings() {
	core := traustest.Start() // a fake core; a real client points at core's URL
	defer core.Close()
	repository := uuid.MustParse("6f1f8e8e-3c9f-4f53-9a43-3d1f2b9b0c11")
	high := "high"
	core.Handle("GET /v1/repositories/{id}/findings/summary", traustest.JSON(http.StatusOK, []traust.FindingRollup{
		{Key: &high, Findings: 3, Open: 2},
	}))
	counts, err := findings.New(core.Client()).SummarizeFindings(context.Background(), repository,
		&traust.SummarizeFindingsParams{GroupBy: traust.FindingDimensionSeverity})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, group := range counts {
		fmt.Printf("%s: %d open of %d\n", *group.Key, group.Open, group.Findings)
	}
	// Output: high: 2 open of 3
}
