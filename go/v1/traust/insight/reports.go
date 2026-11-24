package insight

import (
	"context"
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
