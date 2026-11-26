package findings

import (
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// restatements are why a person may restate a decision.
var restatements = []traust.RestatementReason{traust.OperatorError, traust.DataError, traust.BaselineRewrite}

// validities are the answers a person may give about a finding.
var validities = []traust.Validity{traust.ValidityConfirmed, traust.ValidityFalsePositive, traust.ValidityHardening, traust.ValidityNotAffected}

// decisionValidity words a recorded decision's validity.
func decisionValidity(d traust.Decision) string {
	if d.Validity == nil {
		return "no validity"
	}
	if v, err := d.Validity.AsValidity(); err == nil {
		return string(v)
	}
	return "restated validity"
}

// caseTitle is a review case's finding, or its id when core sent no view.
func caseTitle(c traust.ReviewCase) string {
	if c.View != nil {
		return c.View.Finding.Title
	}
	return c.FindingId.String()
}
