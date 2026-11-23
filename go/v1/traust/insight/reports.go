package insight

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// BeginRequestReport asks core to build a report and follows the build; a ready or failed
// report has finished.
func (c *Client) BeginRequestReport(ctx context.Context, body traust.ReportRequest) (*traust.Poller[traust.Report], error) {
	report, err := c.RequestReport(ctx, body)
	if err != nil {
		return nil, err
	}
	return c.ResumeReport(report.OrderId.String())
}

// ResumeReport follows the report build a resume token names.
func (c *Client) ResumeReport(token string) (*traust.Poller[traust.Report], error) {
	id, err := uuid.Parse(token)
	if err != nil {
		return nil, fmt.Errorf("insight: resume token: %w", err)
	}
	return traust.NewPoller(token, func(ctx context.Context) (traust.Report, bool, error) {
		report, err := c.GetReport(ctx, id)
		if err != nil {
			return traust.Report{}, false, err
		}
		return *report, report.Status == traust.ReportStatusReady || report.Status == traust.ReportStatusFailed, nil
	}), nil
}

// ReportView decodes a ready json report as its contracted view T (a generated …Insight type,
// such as FindingsRollupInsight). A field outside the contract is an error, never dropped.
func ReportView[T any](ctx context.Context, c *Client, report traust.Report) (*T, error) {
	name := viewName(*new(T))
	if name == "" {
		return nil, fmt.Errorf("insight: %T is not a contracted view type", *new(T))
	}
	if report.Report != name || report.Format != traust.ReportFormatJson || report.Status != traust.ReportStatusReady {
		return nil, fmt.Errorf("insight: report %s is a %s %s %s, not a ready json %s", report.OrderId, report.Status, report.Format, report.Report, name)
	}
	body, err := c.GetReportFile(ctx, report.OrderId)
	if err != nil {
		return nil, fmt.Errorf("getting the %s report: %w", name, err)
	}
	if bytes.Equal(bytes.TrimSpace(body), []byte("null")) {
		return nil, fmt.Errorf("traust-core: the %s report is empty", name)
	}
	var view T
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&view); err != nil {
		return nil, fmt.Errorf("decoding the %s report: %w", name, err)
	}
	return &view, nil
}
