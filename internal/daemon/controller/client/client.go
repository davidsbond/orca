package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	controllersvcv1 "github.com/davidsbond/orca/internal/proto/orca/controller/service/v1"
	taskv1 "github.com/davidsbond/orca/internal/proto/orca/task/v1"
	workflowv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/v1"
	"github.com/davidsbond/orca/internal/task"
	"github.com/davidsbond/orca/internal/workflow"
)

type (
	Client struct {
		conn       *grpc.ClientConn
		controller controllersvcv1.ControllerServiceClient
	}
)

func Dial(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:       conn,
		controller: controllersvcv1.NewControllerServiceClient(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) SetWorkflowRunStatus(ctx context.Context, runID string, status workflow.Status, output []byte) error {
	request := &controllersvcv1.SetWorkflowRunStatusRequest{
		WorkflowRunId: runID,
		Status:        workflowv1.Status(status),
		Output:        output,
	}

	if _, err := c.controller.SetWorkflowRunStatus(ctx, request); err != nil {
		return fmt.Errorf("failed to set workflow status: %w", err)
	}

	return nil
}

func (c *Client) SetTaskRunStatus(ctx context.Context, runID string, status task.Status, output []byte) error {
	request := &controllersvcv1.SetTaskRunStatusRequest{
		TaskRunId: runID,
		Status:    taskv1.Status(status),
		Output:    output,
	}

	if _, err := c.controller.SetTaskRunStatus(ctx, request); err != nil {
		return fmt.Errorf("failed to set task status: %w", err)
	}

	return nil
}

func (c *Client) ScheduleTask(ctx context.Context, runID string, name string, input []byte) (string, error) {
	request := &controllersvcv1.ScheduleTaskRequest{
		WorkflowRunId: runID,
		TaskName:      name,
		Input:         input,
	}

	response, err := c.controller.ScheduleTask(ctx, request)
	if err != nil {
		return "", fmt.Errorf("failed to schedule task: %w", err)
	}

	return response.GetTaskRunId(), nil
}

func (c *Client) GetTaskRun(ctx context.Context, runID string) (task.Run, error) {
	request := &controllersvcv1.GetTaskRunRequest{
		TaskRunId: runID,
	}

	response, err := c.controller.GetTaskRun(ctx, request)
	if err != nil {
		return task.Run{}, fmt.Errorf("failed to get task status: %w", err)
	}

	t := response.GetTaskRun()
	return task.Run{
		ID:            t.GetRunId(),
		WorkflowRunID: t.GetWorkflowRunId(),
		TaskName:      t.GetTaskName(),
		CreatedAt:     t.GetCreatedAt().AsTime(),
		ScheduledAt:   t.GetScheduledAt().AsTime(),
		StartedAt:     t.GetStartedAt().AsTime(),
		CompletedAt:   t.GetCompletedAt().AsTime(),
		Status:        task.Status(t.GetStatus()),
		Input:         t.GetInput(),
		Output:        t.GetOutput(),
	}, nil
}

func (c *Client) GetWorkflowRun(ctx context.Context, runID string) (workflow.Run, error) {
	request := &controllersvcv1.GetWorkflowRunRequest{
		WorkflowRunId: runID,
	}

	response, err := c.controller.GetWorkflowRun(ctx, request)
	if err != nil {
		return workflow.Run{}, fmt.Errorf("failed to get workflow status: %w", err)
	}

	t := response.GetWorkflowRun()
	return workflow.Run{
		ID:           t.GetRunId(),
		ParentID:     t.GetParentWorkflowRunId(),
		WorkflowName: t.GetWorkflowName(),
		CreatedAt:    t.GetCreatedAt().AsTime(),
		ScheduledAt:  t.GetScheduledAt().AsTime(),
		StartedAt:    t.GetStartedAt().AsTime(),
		CompletedAt:  t.GetCompletedAt().AsTime(),
		Status:       workflow.Status(t.GetStatus()),
		Input:        t.GetInput(),
		Output:       t.GetOutput(),
	}, nil
}
