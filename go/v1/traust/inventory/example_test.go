package inventory_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
	"github.com/traust-security/traust-sdk/go/v1/traust/inventory"
	"github.com/traust-security/traust-sdk/go/v1/traust/traustest"
)

func ExampleNew() {
	client, err := traust.NewClient(traust.Config{
		BaseURL: os.Getenv("TRAUST_CORE_URL"),
		Tokens:  traust.FileTokenSource{Path: "/var/run/secrets/traust-core/token"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = inventory.New(client)
}

func ExampleClient_ListRepositories() {
	core := traustest.Start() // a fake core; a real client points at core's URL
	defer core.Close()
	core.Handle("GET /v1/repositories", traustest.JSON(http.StatusOK, traust.PageRepository{
		Items: []traust.Repository{{Url: "https://github.com/example/api", Public: true}},
	}))
	url := "https://github.com/example/api"
	found, err := inventory.New(core.Client()).ListRepositories(context.Background(), &traust.ListRepositoriesParams{Url: &url})
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
	for repo, err := range inventory.New(core.Client()).ListRepositoriesAll(context.Background(), nil) {
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

func ExampleProblemError() {
	core := traustest.Start() // a fake core; a real client points at core's URL
	defer core.Close()
	core.Handle("GET /v1/repositories/{id}", traustest.Problem(http.StatusNotFound, "repository not registered"))
	_, err := inventory.New(core.Client()).GetRepository(context.Background(), uuid.New())
	var problem *traust.ProblemError
	if errors.As(err, &problem) {
		fmt.Println(problem.StatusCode, problem.Problem.Detail)
	}
	// Output: 404 repository not registered
}
