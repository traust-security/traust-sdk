package ops

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// record posts the executor's bytes through the run's one template.
func (run RecordingRun) record(ctx context.Context, l *lease, raw []byte) (RecordingReceipt, error) {
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
		if err := l.stamp(&body.RunOrderId, &body.Provenance, &body.Payload, payload); err != nil {
			return receipt, err
		}
		receipt.Remediation, err = run.Client.RecordRemediation(ctx, body)
		return receipt, err
	}
	return receipt, errors.New("ops: recording run has no request template")
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
