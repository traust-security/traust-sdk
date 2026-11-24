// Code generated from traust-contracts 3ec3df824c7c0a225af78030111f84068d3d058e. DO NOT EDIT.

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
