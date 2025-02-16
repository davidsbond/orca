package client

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	workflowsvcv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/service/v1"
	"github.com/davidsbond/orca/internal/workflow"
)

func (c *Client) ScheduleWorkflow(ctx context.Context, name string, input json.RawMessage) (string, error) {
	request := &workflowsvcv1.ScheduleRequest{
		WorkflowName: name,
		Input:        input,
	}

	response, err := c.workflows.Schedule(ctx, request)
	if err != nil {
		return "", formatError(err)
	}

	return response.GetWorkflowRunId(), nil
}

func (c *Client) GetWorkflowRun(ctx context.Context, id string) (workflow.Run, error) {
	request := &workflowsvcv1.GetRunRequest{
		WorkflowRunId: id,
	}

	response, err := c.workflows.GetRun(ctx, request)
	switch {
	case status.Code(err) == codes.NotFound:
		return workflow.Run{}, ErrNotFound
	case err != nil:
		return workflow.Run{}, formatError(err)
	}

	run := response.GetWorkflowRun()

	return workflow.Run{
		ID:           run.GetRunId(),
		ParentID:     run.GetParentWorkflowRunId(),
		WorkflowName: run.GetWorkflowName(),
		CreatedAt:    run.GetCreatedAt().AsTime(),
		ScheduledAt:  run.GetScheduledAt().AsTime(),
		StartedAt:    run.GetStartedAt().AsTime(),
		CompletedAt:  run.GetCompletedAt().AsTime(),
		Status:       workflow.Status(run.GetStatus()),
		Input:        run.GetInput(),
		Output:       run.GetOutput(),
	}, nil
}
