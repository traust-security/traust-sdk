package storage

import (
	"context"
	"testing"
)

func TestThreatReadersRetainPinnedContractColumns(t *testing.T) {
	eachDialect(t, func(t *testing.T, client *Client) {
		ctx := context.Background()
		payload := replaceJSONField(t, ratedThreatModel(t), func(document map[string]any) {
			document["threats"].([]any)[1].(map[string]any)["review_surfaces"] = []string{"scope-enforcement"}
		})
		if _, err := client.saveNamed(ctx, "threat-model", payload, bindingFor("threat-model")); err != nil {
			t.Fatal(err)
		}
		rows, err := client.QueryThreatCurrent(ctx, []string{"local"})
		if err != nil || len(rows) != 2 {
			t.Fatalf("threat query = %+v, %v; want two threats", rows, err)
		}
		for _, row := range rows {
			switch row.ThreatID {
			case "T1":
				if row.AttackRefs == nil || !sameJSONList(*row.AttackRefs, "T1190", "T1078") ||
					row.Severity != nil || row.RiskRating != nil || row.ReviewSurfaces != nil ||
					row.Impact == nil || *row.Impact != "critical" || row.Likelihood == nil || *row.Likelihood != "almost_certain" {
					t.Fatalf("legacy threat lost its pinned columns: %+v", row)
				}
			case "T2":
				if row.RiskRating == nil || row.Severity == nil || *row.Severity != "high" ||
					row.LikelihoodScore == nil || *row.LikelihoodScore != 4.375 ||
					row.LikelihoodLevel == nil || *row.LikelihoodLevel != "medium" ||
					row.ImpactScore == nil || *row.ImpactScore != 7.25 ||
					row.ImpactLevel == nil || *row.ImpactLevel != "high" ||
					row.ImpactBasis == nil || *row.ImpactBasis != "technical" ||
					row.ReviewSurfaces == nil || !sameJSONList(*row.ReviewSurfaces, "scope-enforcement") ||
					row.Impact != nil || row.Likelihood != nil || row.Score != nil {
					t.Fatalf("rated threat lost its pinned columns: %+v", row)
				}
				assertProjectionJSON(t, []byte(*row.RiskRating), []byte(owaspWorkedExample))
			default:
				t.Fatalf("unexpected threat: %+v", row)
			}
		}
		exposure, err := client.QueryThreatExposure(ctx, []string{"local"})
		if err != nil || len(exposure) != 2 {
			t.Fatalf("threat exposure = %+v, %v; want two distinct groups", exposure, err)
		}
		for _, row := range exposure {
			if row.Threats != 1 || row.Subjects != 1 {
				t.Fatalf("threat exposure changed grain: %+v", row)
			}
			if row.Severity != nil {
				if *row.Severity != "high" || row.Impact != nil || row.Likelihood != nil || row.TopScore != nil {
					t.Fatalf("rated threat exposure lost severity: %+v", row)
				}
			} else if row.Impact == nil || *row.Impact != "critical" || row.Likelihood == nil || *row.Likelihood != "almost_certain" {
				t.Fatalf("legacy threat exposure lost labels: %+v", row)
			}
		}
	})
}
