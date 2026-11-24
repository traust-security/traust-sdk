// Code generated from traust-contracts 58803ea798ad6431e10da337eb3a28349c1c9992. DO NOT EDIT.

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
