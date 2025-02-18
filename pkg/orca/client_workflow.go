package orca

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/disiqueira/gotree"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	taskv1 "github.com/davidsbond/orca/internal/proto/orca/task/v1"
	workflowsvcv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/service/v1"
	workflowv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/v1"
	"github.com/davidsbond/orca/internal/task"
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

	return workflowRunFromProto(run), nil
}

type (
	Description struct {
		Run     workflow.Run     `json:"run"`
		Actions []WorkflowAction `json:"actions,omitempty"`
	}

	WorkflowAction struct {
		WorkflowRun *Description `json:"workflowRun,omitempty"`
		TaskRun     *task.Run    `json:"taskRun,omitempty"`
	}
)

func (c *Client) DescribeWorkflowRun(ctx context.Context, runID string) (Description, error) {
	request := &workflowsvcv1.DescribeRunRequest{
		WorkflowRunId: runID,
	}

	response, err := c.workflows.DescribeRun(ctx, request)
	if err != nil {
		return Description{}, formatError(err)
	}

	description := response.GetDescription()

	result := Description{
		Run:     workflowRunFromProto(description.GetRun()),
		Actions: workflowActionsFromProto(description.GetActions()),
	}

	return result, nil
}

func workflowRunFromProto(run *workflowv1.Run) workflow.Run {
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
	}
}

func workflowActionsFromProto(actions []*workflowv1.WorkflowAction) []WorkflowAction {
	out := make([]WorkflowAction, len(actions))
	for i, action := range actions {
		if wf := action.GetWorkflowRun(); wf != nil {
			out[i] = WorkflowAction{
				WorkflowRun: &Description{
					Run:     workflowRunFromProto(wf.GetRun()),
					Actions: workflowActionsFromProto(wf.GetActions()),
				},
			}

			continue
		}

		if tsk := action.GetTaskRun(); tsk != nil {
			out[i] = WorkflowAction{
				TaskRun: taskRunFromProto(tsk.GetRun()),
			}

			continue
		}
	}

	return out
}

func taskRunFromProto(run *taskv1.Run) *task.Run {
	return &task.Run{
		ID:            run.GetRunId(),
		WorkflowRunID: run.GetWorkflowRunId(),
		TaskName:      run.GetTaskName(),
		CreatedAt:     run.GetCreatedAt().AsTime(),
		ScheduledAt:   run.GetScheduledAt().AsTime(),
		StartedAt:     run.GetStartedAt().AsTime(),
		CompletedAt:   run.GetCompletedAt().AsTime(),
		Status:        task.Status(run.GetStatus()),
		Input:         run.GetInput(),
		Output:        run.GetOutput(),
	}
}

func (d Description) String() string {
	return d.tree().Print()
}

func (d Description) tree() gotree.Tree {
	wf := fmt.Sprintf("Workflow: %s (%s) [%s]",
		d.Run.WorkflowName,
		d.Run.CompletedAt.Sub(d.Run.CreatedAt).Truncate(time.Millisecond),
		d.Run.Status,
	)

	root := gotree.New(wf)

	for _, action := range d.Actions {
		if action.TaskRun != nil {
			tsk := fmt.Sprintf("Task: %s (%s) [%s]",
				action.TaskRun.TaskName,
				action.TaskRun.CompletedAt.Sub(action.TaskRun.CreatedAt).Truncate(time.Millisecond),
				action.TaskRun.Status,
			)

			root.Add(tsk)
			continue
		}

		if action.WorkflowRun != nil {
			root.AddTree(action.WorkflowRun.tree())
		}
	}

	return root
}
