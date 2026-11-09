package ops

import (
	"context"
	"encoding/json"
	"errors"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// record posts the executor's report through the run's one template. The remediation
// template reconciles its source_finding refs to the door's bare finding ids first; every
// other template posts the executor's exact original bytes.
func (run RecordingRun) record(ctx context.Context, l *lease, raw []byte, preparedFindings []uuid.UUID) (RecordingReceipt, error) {
	var receipt RecordingReceipt
	var err error
	// raw passed utf8.Valid, so the string carries the exact original bytes.
	payload := string(raw)
	request := run.Request
	switch {
	case request.Assessment != nil:
		body := *request.Assessment
		if err := l.stamp(&body.RunOrderId, &body.Provenance, &body.Payload, payload); err != nil {
			return receipt, err
		}
		receipt.Assessment, err = run.Client.RecordAssessment(ctx, body)
		return receipt, err
	case request.Triage != nil:
		body := *request.Triage
		if err := l.stamp(&body.RunOrderId, &body.Provenance, &body.Payload, payload); err != nil {
			return receipt, err
		}
		receipt.Triage, err = run.Client.RecordTriage(ctx, request.AssessmentID, body)
		return receipt, err
	case request.Validation != nil:
		body := *request.Validation
		if err := l.stamp(&body.RunOrderId, &body.Provenance, &body.Payload, payload); err != nil {
			return receipt, err
		}
		receipt.Validation, err = run.Client.RecordValidation(ctx, body)
		return receipt, err
	case request.Verification != nil:
		body := *request.Verification
		if err := l.stamp(&body.RunOrderId, &body.Provenance, &body.Payload, payload); err != nil {
			return receipt, err
		}
		receipt.Verification, err = run.Client.RecordVerification(ctx, body)
		return receipt, err
	case request.Remediation != nil:
		body := *request.Remediation
		payload := string(reconcileFindingRefs(raw, preparedFindings))
		if err := l.stamp(&body.RunOrderId, &body.Provenance, &body.Payload, payload); err != nil {
			return receipt, err
		}
		receipt.Remediation, err = run.Client.RecordRemediation(ctx, body)
		return receipt, err
	}
	return receipt, errors.New("ops: recording run has no request template")
}

// reconcileFindingRefs rewrites a remediation report's qualified source_finding refs to the
// bare finding ids core's record door validates. Contracts document `source_findings[].
// finding_ref` as the source report's stable reference — '<report>#<finding-id>' — while
// core's door accepts only the bare finding id its prepare served, so a worker that follows
// the documented spelling is refused with 422 "finding_ref is not a finding id from prepare"
// (found live on the dev cluster's first full record). The SDK owns the record seam, so it
// owns the reconciliation, and it binds to what prepare served: the id after the last '#'
// is rewritten only when it names one of the prepared work item's findings. Every other byte
// rides through untouched (json.RawMessage surgery, not a re-encoding), and a report whose
// refs reconcile to nothing — no source_findings, a non-string ref, a tail that names no
// prepared finding — is returned exactly as the executor produced it: core, not the SDK,
// validates the report, and its door's refusal classifies the output.
func reconcileFindingRefs(report []byte, prepared []uuid.UUID) []byte {
	if len(prepared) == 0 {
		return report
	}
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(report, &doc); err != nil {
		return report
	}
	rawFindings, ok := doc["source_findings"]
	if !ok {
		return report
	}
	var findings []map[string]json.RawMessage
	if err := json.Unmarshal(rawFindings, &findings); err != nil || len(findings) == 0 {
		return report
	}
	rewritten := false
	for _, finding := range findings {
		refRaw, ok := finding["finding_ref"]
		if !ok {
			return report
		}
		var ref string
		if err := json.Unmarshal(refRaw, &ref); err != nil {
			return report
		}
		cut := strings.LastIndex(ref, "#")
		if cut < 0 {
			continue // already bare: the door validates it against prepare
		}
		id, err := uuid.Parse(ref[cut+1:])
		if err != nil || !slices.Contains(prepared, id) {
			continue // names no prepared finding: the door's refusal is the answer
		}
		encoded, err := json.Marshal(id.String())
		if err != nil {
			return report
		}
		finding["finding_ref"] = encoded
		rewritten = true
	}
	if !rewritten {
		return report
	}
	encodedFindings, err := json.Marshal(findings)
	if err != nil {
		return report
	}
	doc["source_findings"] = encodedFindings
	encoded, err := json.Marshal(doc)
	if err != nil {
		return report
	}
	return encoded
}

// stamp binds a template to the held order: its id and start time are the record's run identity.
func (l *lease) stamp(runOrderID *uuid.UUID, provenance *traust.Provenance, payload *string, raw string) error {
	id, started := l.id(), *l.origin.StartedAt
	switch {
	case *runOrderID != uuid.Nil && *runOrderID != id:
		return errors.New("ops: template run_order_id is not the server's order")
	case provenance.RunId != uuid.Nil && provenance.RunId != id:
		return errors.New("ops: template provenance run_id is not the server's order")
	case !provenance.OccurredAt.IsZero() && !provenance.OccurredAt.Equal(started):
		return errors.New("ops: template provenance occurred_at is not the server's start time")
	case provenance.Source == "":
		return errors.New("ops: template needs an explicit provenance source")
	}
	*runOrderID, provenance.RunId, provenance.OccurredAt, *payload = id, id, started, raw
	return nil
}

func (request RecordingRequest) template() (uuid.UUID, string, traust.Provenance) {
	switch {
	case request.Assessment != nil:
		return request.Assessment.RunOrderId, request.Assessment.Payload, request.Assessment.Provenance
	case request.Triage != nil:
		return request.Triage.RunOrderId, request.Triage.Payload, request.Triage.Provenance
	case request.Validation != nil:
		return request.Validation.RunOrderId, request.Validation.Payload, request.Validation.Provenance
	case request.Verification != nil:
		return request.Verification.RunOrderId, request.Verification.Payload, request.Verification.Provenance
	case request.Remediation != nil:
		return request.Remediation.RunOrderId, request.Remediation.Payload, request.Remediation.Provenance
	}
	return uuid.Nil, "", traust.Provenance{}
}
