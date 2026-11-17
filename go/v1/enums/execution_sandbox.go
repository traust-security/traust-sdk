// Code generated from traust-contracts 31d0d884672149e606934c2844ffcbf3bdd4b925. DO NOT EDIT.

package enums

// ExecutionSandbox — How lanes run a target repository's own build and test code: none (scrubbed environment on a disposable copy) or podman (rootless nested container).
type ExecutionSandbox string

const (
	ExecutionSandboxNone   ExecutionSandbox = "none"
	ExecutionSandboxPodman ExecutionSandbox = "podman"
)

// ExecutionSandboxValues returns all valid ExecutionSandbox values.
func ExecutionSandboxValues() []ExecutionSandbox {
	return []ExecutionSandbox{
		ExecutionSandboxNone,
		ExecutionSandboxPodman,
	}
}
