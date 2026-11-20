package ops

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

type defaultRefAssessment struct {
	client          *traust.Client
	source          traust.SourceType
	assessorVersion string
}

// AssessDefaultRef records each order's report as an assessment of its repository's current
// default ref: the commit core says the repository is at, never one the runner picked. The
// executor's Input is that subject (a traust.SubjectRef).
func AssessDefaultRef(client *traust.Client, source traust.SourceType, assessorVersion string) (RecordingPreparer, error) {
	if client == nil {
		return nil, errors.New("ops: AssessDefaultRef needs a client")
	}
	if source == "" || assessorVersion == "" {
		return nil, errors.New("ops: AssessDefaultRef needs a source and an assessor version")
	}
	return defaultRefAssessment{client, source, assessorVersion}, nil
}

func (a defaultRefAssessment) PrepareRecording(ctx context.Context, order traust.RunOrder) (RecordingRequest, error) {
	if order.RepositoryId == nil {
		return RecordingRequest{}, &PreparationRefusedError{Reason: traust.RunOrderReasonBadRequest, Err: errors.New("the order names no repository")}
	}
	id := *order.RepositoryId
	repository, err := a.client.GetRepository(ctx, id)
	if err != nil {
		return RecordingRequest{}, fmt.Errorf("reading repository %s: %w", id, err)
	}
	ref, err := a.client.GetCurrentDefaultRef(ctx, id)
	if errors.Is(err, traust.ErrNotFound) {
		return RecordingRequest{}, &PreparationRefusedError{Reason: traust.RunOrderReasonNotFound, Err: fmt.Errorf("repository %s has no default ref", id)}
	}
	if err != nil {
		return RecordingRequest{}, fmt.Errorf("reading repository %s's default ref: %w", id, err)
	}
	if ref.Id == nil || ref.RepositoryId != id || ref.Sha == "" {
		return RecordingRequest{}, fmt.Errorf("core answered an unusable default ref for repository %s", id)
	}
	input, err := json.Marshal(traust.SubjectRef{RepositoryId: id, RepositoryRefId: *ref.Id, RepositoryUrl: repository.Url, Sha: ref.Sha})
	if err != nil {
		return RecordingRequest{}, err
	}
	return RecordingRequest{
		Assessment: &traust.AssessmentRecordRequest{
			RepositoryRefId: *ref.Id, AssessorVersion: a.assessorVersion, Provenance: traust.Provenance{Source: a.source},
		},
		Input: input,
	}, nil
}
