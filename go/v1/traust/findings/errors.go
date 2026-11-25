package findings

import "errors"

// Artifact refusals, for errors.Is: an artifact the SDK cannot type stops its consumer.
var (
	// ErrUnknownKind: artifact content carried no Traust-Artifact-Kind header, or one naming
	// no kind core publishes.
	ErrUnknownKind = errors.New("findings: artifact content without a kind core publishes")
	// ErrKindMismatch: the artifact's kind is not the contract type it was decoded into.
	ErrKindMismatch = errors.New("findings: artifact kind does not match the decoded type")
	// ErrNotJSON: the artifact's media type carries no document to decode.
	ErrNotJSON = errors.New("findings: artifact media type is not JSON")
)
