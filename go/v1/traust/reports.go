package traust

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// insightView is a generated …Insight type: one view's contracts ViewModel.
type insightView[T any] interface {
	*T
	insightName() ViewName
}

// BeginRequestReport asks core to build a report and follows the build; a ready or failed
// report has finished.
func (c *Client) BeginRequestReport(ctx context.Context, body ReportRequest) (*Poller[Report], error) {
	report, err := c.RequestReport(ctx, body)
	if err != nil {
		return nil, err
	}
	return c.ResumeReport(report.OrderId.String())
}

// ResumeReport follows the report build a resume token names.
func (c *Client) ResumeReport(token string) (*Poller[Report], error) {
	id, err := uuid.Parse(token)
	if err != nil {
		return nil, fmt.Errorf("traust: resume token: %w", err)
	}
	return &Poller[Report]{token: token, poll: func(ctx context.Context) (Report, bool, error) {
		report, err := c.GetReport(ctx, id)
		if err != nil {
			return Report{}, false, err
		}
		return *report, report.Status == ReportStatusReady || report.Status == ReportStatusFailed, nil
	}}, nil
}

// ReportView decodes a ready json report as its contracted view T (a generated …Insight type,
// such as FindingsRollupInsight). A field outside the contract is an error, never dropped.
func ReportView[T any, PT insightView[T]](ctx context.Context, c *Client, report Report) (*T, error) {
	name := PT(nil).insightName()
	if report.Report != name || report.Format != ReportFormatJson || report.Status != ReportStatusReady {
		return nil, fmt.Errorf("traust: report %s is a %s %s %s, not a ready json %s", report.OrderId, report.Status, report.Format, report.Report, name)
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
