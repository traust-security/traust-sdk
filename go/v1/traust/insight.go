package traust

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// insightView is a generated …Insight type: one Insight view's contracts ViewModel.
type insightView[T any] interface {
	*T
	insightName() ViewName
}

// GetInsightAs is GetInsight decoded as view T's contracts ViewModel (T is a generated …Insight
// type, such as FindingsRollupInsight). A field outside that shape is an error, never dropped.
func GetInsightAs[T any, PT insightView[T]](ctx context.Context, c *Client, params *GetInsightParams) (*T, error) {
	name := PT(nil).insightName()
	raw, err := c.ClientInterface.GetInsight(ctx, name, params)
	if err != nil {
		return nil, err
	}
	if err := checkRefusal(raw); err != nil {
		return nil, err
	}
	response, err := ParseGetInsightResponse(raw)
	if err != nil {
		return nil, err
	}
	if _, err := Decode(response.HTTPResponse, response.Body, response.JSON200); err != nil {
		return nil, err
	}
	var view T
	decoder := json.NewDecoder(bytes.NewReader(response.Body))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&view); err != nil {
		return nil, fmt.Errorf("decoding the %s insight: %w", name, err)
	}
	return &view, nil
}
