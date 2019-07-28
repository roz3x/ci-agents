package data

import (
	"context"

	"github.com/tinyci/ci-agents/ci-gen/grpc/services/data"
	"github.com/tinyci/ci-agents/ci-gen/grpc/types"
	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/model"
)

// PutSubmission puts a submission into the datasvc
func (c *Client) PutSubmission(sub *model.Submission) (*model.Submission, *errors.Error) {
	s, err := c.client.PutSubmission(context.Background(), sub.ToProto())
	if err != nil {
		return nil, errors.New(err)
	}

	return model.NewSubmissionFromProto(s)
}

// GetSubmissionByID returns the submission for the given ID.
func (c *Client) GetSubmissionByID(id int64) (*model.Submission, *errors.Error) {
	s, err := c.client.GetSubmission(context.Background(), &types.IntID{ID: id})
	if err != nil {
		return nil, errors.New(err)
	}

	return model.NewSubmissionFromProto(s)
}

// GetTasksForSubmission returns the tasks for the given submission; with pagination
func (c *Client) GetTasksForSubmission(sub *model.Submission, page, perPage int64) ([]*model.Task, *errors.Error) {
	tasks, err := c.client.GetSubmissionTasks(context.Background(), &data.SubmissionQuery{Submission: sub.ToProto(), Page: page, PerPage: perPage})
	if err != nil {
		return nil, errors.New(err)
	}

	list := []*model.Task{}

	for _, task := range tasks.Tasks {
		t, err := model.NewTaskFromProto(task)
		if err != nil {
			return nil, err
		}
		list = append(list, t)
	}

	return list, nil
}

// ListSubmissions lists the submissions with pagination, and an optional (just
// pass empty strings if undesired) repository and sha filter.
func (c *Client) ListSubmissions(page, perPage int64, repository, sha string) ([]*model.Submission, *errors.Error) {
	list, err := c.client.ListSubmissions(context.Background(), &data.RepositoryFilterRequest{Page: page, PerPage: perPage, Repository: repository, Sha: sha})
	if err != nil {
		return nil, errors.New(err)
	}

	subs := []*model.Submission{}

	for _, sub := range list.Submissions {
		newSub, err := model.NewSubmissionFromProto(sub)
		if err != nil {
			return nil, err
		}

		subs = append(subs, newSub)
	}

	return subs, nil
}
