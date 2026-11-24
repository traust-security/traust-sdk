// Code generated from traust-contracts c2f6e6fb12270f8b3c8703da7bf39bd01776e4b1. DO NOT EDIT.

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
