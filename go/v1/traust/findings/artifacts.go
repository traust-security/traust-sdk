package findings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"reflect"
	"strings"

	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// Object is one artifact as core serves it: what it is — kind, media type, digest, size — and
// its exact bytes, never where core keeps them.
type Object struct {
	Kind      traust.ArtifactKind
	Digest    string
	MediaType string
	Size      int64
	Body      []byte
}

// Get reads one artifact in a single call, the way a cloud SDK gets an object: the kind, digest
// and media type from the content response's headers, the bytes — and their size — from its
// body, which a proxy may serve chunked and the transport may decompress, so no length header
// outlives the hop. It refuses a response whose Traust-Artifact-Kind header is missing or
// names no kind core publishes: an artifact without a kind is never handed on as bytes.
func (c *Client) Get(ctx context.Context, id openapi_types.UUID) (*Object, error) {
	raw, err := c.core.Raw().GetArtifactContent(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Get(%s): %w", id, err)
	}
	if err := traust.Refusal(raw); err != nil {
		return nil, fmt.Errorf("Get(%s): %w", id, err)
	}
	defer func() { _ = raw.Body.Close() }()
	// the bytes come only with 200: another success is no artifact, never empty bytes
	if raw.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Get(%s): traust-core: HTTP %d, want 200 with the artifact", id, raw.StatusCode)
	}
	kind := traust.ArtifactKind(raw.Header.Get("Traust-Artifact-Kind"))
	if kind == "" {
		return nil, fmt.Errorf("Get(%s): the content carries no Traust-Artifact-Kind header: %w", id, ErrUnknownKind)
	}
	if !artifactKinds[kind] {
		return nil, fmt.Errorf("Get(%s): Traust-Artifact-Kind %q names no kind core publishes: %w", id, kind, ErrUnknownKind)
	}
	etag := raw.Header.Get("ETag")
	if etag == "" {
		return nil, fmt.Errorf("Get(%s): the content carries no ETag header for the digest core documents", id)
	}
	mediaType := raw.Header.Get("Content-Type")
	if mediaType == "" {
		return nil, fmt.Errorf("Get(%s): the content carries no Content-Type header", id)
	}
	body, err := io.ReadAll(raw.Body)
	if err != nil {
		return nil, fmt.Errorf("Get(%s): %w", id, err)
	}
	return &Object{Kind: kind, Digest: strings.Trim(etag, `"`), MediaType: mediaType, Size: int64(len(body)), Body: body}, nil
}

// Decode reads an artifact's document: the JSON body of an object whose kind is T's, so a
// report arrives as the contract type core publishes instead of bytes to parse by hand. A kind
// mismatch is ErrKindMismatch; a media type that is not JSON cannot be decoded, and neither can
// a null body. The kinds
// without a contract — a diff, an sbom, a feed — are read from Object's Body.
func Decode[T any](obj *Object) (*T, error) {
	typ := reflect.TypeFor[T]()
	kind, contract := kindsByType[typ]
	if !contract {
		return nil, fmt.Errorf("findings: decoding an artifact into %s: core publishes no artifact contract of that type", typ)
	}
	if obj.Kind != kind {
		return nil, fmt.Errorf("findings: decoding a %s artifact as a %s: %w", obj.Kind, kind, ErrKindMismatch)
	}
	media, _, err := mime.ParseMediaType(obj.MediaType)
	if err != nil || media != "application/json" {
		return nil, fmt.Errorf("findings: decoding the %s artifact served as %s: %w", obj.Kind, obj.MediaType, ErrNotJSON)
	}
	if bytes.Equal(bytes.TrimSpace(obj.Body), []byte("null")) {
		return nil, fmt.Errorf("findings: decoding the %s artifact's body: null is no document", obj.Kind)
	}
	var doc T
	if err := json.Unmarshal(obj.Body, &doc); err != nil {
		return nil, fmt.Errorf("findings: decoding the %s artifact's body: %w", obj.Kind, err)
	}
	return &doc, nil
}
