package ops

import (
	"context"
	"encoding/json"
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// record posts the executor's report through the run's one template. The remediation
// template reconciles its source_finding refs to the door's bare finding ids first; every
// other template posts the executor's exact original bytes.
func (run RecordingRun) record(ctx context.Context, l *lease, raw []byte, capturedAt time.Time, preparedFindings []uuid.UUID) (RecordingReceipt, error) {
	var receipt RecordingReceipt
	// raw passed utf8.Valid, so the string carries the exact original bytes.
	payload := string(raw)
	request := run.Request
	id, generation, _, provenance := request.template()
	if err := l.validateBinding(id, generation, provenance); err != nil {
		return receipt, err
	}
	id, generation = l.id(), l.origin.LeaseGeneration
	provenance.RunId = id
	if provenance.OccurredAt.IsZero() {
		provenance.OccurredAt = capturedAt
	}
	var err error
	switch {
	case request.Assessment != nil:
		body := *request.Assessment
		body.RunOrderId, body.LeaseGeneration, body.Provenance, body.Payload = id, generation, provenance, payload
		receipt.Assessment, err = run.Client.RecordAssessment(ctx, body)
		return receipt, err
	case request.Triage != nil:
		body := *request.Triage
		body.RunOrderId, body.LeaseGeneration, body.Provenance, body.Payload = id, generation, provenance, payload
		receipt.Triage, err = run.Client.RecordTriage(ctx, request.AssessmentID, body)
		return receipt, err
	case request.Validation != nil:
		body := *request.Validation
		body.RunOrderId, body.LeaseGeneration, body.Provenance, body.Payload = id, generation, provenance, payload
		receipt.Validation, err = run.Client.RecordValidation(ctx, body)
		return receipt, err
	case request.Verification != nil:
		body := *request.Verification
		body.RunOrderId, body.LeaseGeneration, body.Provenance, body.Payload = id, generation, provenance, payload
		receipt.Verification, err = run.Client.RecordVerification(ctx, body)
		return receipt, err
	case request.Remediation != nil:
		body := *request.Remediation
		payload := string(reconcileFindingRefs(raw, preparedFindings))
		body.RunOrderId, body.LeaseGeneration, body.Provenance, body.Payload = id, generation, provenance, payload
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
// is rewritten only when it names one of the prepared work item's findings. Every value
// rides through as raw JSON (RawMessage surgery — only the rewritten refs re-encode), the
// document itself re-marshals with sorted keys and normalized whitespace, exactly the
// encoding the deleted executor-side hook produced; a report whose refs reconcile to
// nothing — no source_findings, a non-string ref, a tail that names no prepared finding —
// is returned exactly as the executor produced it, bytes untouched: core, not the SDK,
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

// A consumer may pre-bind a template, but never to a different order or attempt.
func (l *lease) validateBinding(runOrderID uuid.UUID, generation int, provenance traust.Provenance) error {
	id := l.id()
	switch {
	case runOrderID != uuid.Nil && runOrderID != id:
		return errors.New("ops: template run_order_id is not the server's order")
	case provenance.RunId != uuid.Nil && provenance.RunId != id:
		return errors.New("ops: template provenance run_id is not the server's order")
	case generation != 0 && generation != l.origin.LeaseGeneration:
		return errors.New("ops: template lease_generation is not the held generation")
	case provenance.Source == "":
		return errors.New("ops: template needs an explicit provenance source")
	}
	return nil
}

func (request RecordingRequest) template() (uuid.UUID, int, string, traust.Provenance) {
	switch {
	case request.Assessment != nil:
		return request.Assessment.RunOrderId, request.Assessment.LeaseGeneration, request.Assessment.Payload, request.Assessment.Provenance
	case request.Triage != nil:
		return request.Triage.RunOrderId, request.Triage.LeaseGeneration, request.Triage.Payload, request.Triage.Provenance
	case request.Validation != nil:
		return request.Validation.RunOrderId, request.Validation.LeaseGeneration, request.Validation.Payload, request.Validation.Provenance
	case request.Verification != nil:
		return request.Verification.RunOrderId, request.Verification.LeaseGeneration, request.Verification.Payload, request.Verification.Provenance
	case request.Remediation != nil:
		return request.Remediation.RunOrderId, request.Remediation.LeaseGeneration, request.Remediation.Payload, request.Remediation.Provenance
	}
	return uuid.Nil, 0, "", traust.Provenance{}
}
