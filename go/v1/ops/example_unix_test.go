//go:build unix

package ops_test

import (
	"context"
	"fmt"

	"github.com/traust-security/traust-sdk/go/v1/ops"
)

// The command prints its JobResult as JSON; that output is also the report.
func ExampleProcessExecutor() {
	produced, err := ops.ProcessExecutor{}.Execute(context.Background(), ops.Execution{
		Argv: []string{"sh", "-c", `echo '{"outcome":"succeeded"}'`},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(produced.Result.Outcome)
	// Output: succeeded
}
