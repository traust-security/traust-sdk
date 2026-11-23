package traust_test

import (
	"fmt"
	"os"

	"github.com/traust-security/traust-sdk/go/v1/traust"
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
