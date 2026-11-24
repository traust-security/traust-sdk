package traust

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// insightView is a generated …Insight type: one view's contracts ViewModel.
type insightView[T any] interface {
	*T
	insightName() ViewName
}

// ViewNameOf returns the core ViewName associated with a generated …Insight type T.
func ViewNameOf[T any, PT insightView[T]]() ViewName {
	return PT(nil).insightName()
}

// DecodeReportView decodes a JSON report file body strictly as view T's contracts ViewModel
// (such as FindingsRollupInsight). Any field outside the contract is an error, never dropped.
func DecodeReportView[T any, PT insightView[T]](body []byte) (*T, error) {
	name := PT(nil).insightName()
	var view T
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&view); err != nil {
		return nil, fmt.Errorf("decoding the %s report: %w", name, err)
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("decoding the %s report: trailing data after JSON document", name)
	}
	return &view, nil
}

// GetReportFileAs fetches a ready JSON report's file and strictly decodes it as view T's
// contracts ViewModel.
func GetReportFileAs[T any, PT insightView[T]](ctx context.Context, c *Client, orderID openapi_types.UUID) (*T, error) {
	body, err := c.GetReportFile(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return DecodeReportView[T, PT](body)
}
