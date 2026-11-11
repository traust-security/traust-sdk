package traust_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

func ExampleNewClient() {
	client, err := traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_CORE_URL"),
		Tokens:  traust.FileTokenSource{Path: "/var/run/secrets/traust-core/token"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = client
}

func ExampleClient_ListRepositories() {
	core := traustest.Start() // a fake core; a real client points at core's URL
	defer core.Close()
	core.Handle("GET /v1/repositories", traustest.JSON(http.StatusOK, traust.PageRepository{
		Items: []traust.Repository{{Url: "https://github.com/example/api", Public: true}},
	}))
	url := "https://github.com/example/api"
	found, err := core.Client().ListRepositories(context.Background(), &traust.ListRepositoriesParams{Url: &url})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(found.Items), found.Items[0].Url)
	// Output: 1 https://github.com/example/api
}

func ExampleClient_ListRepositoriesAll() {
	core := traustest.Start() // a fake core; a real client points at core's URL
	defer core.Close()
	core.Handle("GET /v1/repositories", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := traust.PageRepository{Items: []traust.Repository{{Url: "https://github.com/example/b"}}}
		if r.URL.Query().Get("cursor") == "" {
			next := "page-2"
			page = traust.PageRepository{Items: []traust.Repository{{Url: "https://github.com/example/a"}}, Next: &next}
		}
		traustest.JSON(http.StatusOK, page).ServeHTTP(w, r)
	}))
	for repo, err := range core.Client().ListRepositoriesAll(context.Background(), nil) {
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(repo.Url)
	}
	// Output:
	// https://github.com/example/a
	// https://github.com/example/b
}

func ExampleClient_SummarizeFindings() {
	core := traustest.Start() // a fake core; a real client points at core's URL
	defer core.Close()
	repository := uuid.MustParse("6f1f8e8e-3c9f-4f53-9a43-3d1f2b9b0c11")
	high := "high"
	core.Handle("GET /v1/repositories/{id}/findings/summary", traustest.JSON(http.StatusOK, []traust.FindingRollup{
		{Key: &high, Findings: 3, Open: 2},
	}))
	counts, err := core.Client().SummarizeFindings(context.Background(), repository,
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

func ExampleProblemError() {
	core := traustest.Start() // a fake core; a real client points at core's URL
	defer core.Close()
	core.Handle("GET /v1/repositories/{id}", traustest.Problem(http.StatusNotFound, "repository not registered"))
	_, err := core.Client().GetRepository(context.Background(), uuid.New())
	var problem *traust.ProblemError
	if errors.As(err, &problem) {
		fmt.Println(problem.StatusCode, problem.Problem.Detail)
	}
	// Output: 404 repository not registered
}
