// Code generated from traust-contracts 05e1b0f8f7866d1a12e18770cc79229731c1df1d. DO NOT EDIT.

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
