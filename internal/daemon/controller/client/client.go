package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/davidsbond/orca/internal/daemon/worker/task"
	"github.com/davidsbond/orca/internal/daemon/worker/workflow"
	controllersvcv1 "github.com/davidsbond/orca/internal/proto/orca/controller/service/v1"
	taskv1 "github.com/davidsbond/orca/internal/proto/orca/task/v1"
	workflowv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/v1"
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

func (c *Client) SetWorkflowStatus(ctx context.Context, runID string, status workflow.Status, output []byte) error {
	request := &controllersvcv1.SetWorkflowStatusRequest{
		WorkflowRunId: runID,
		Status:        workflowv1.Status(status),
		Output:        output,
	}

	if _, err := c.controller.SetWorkflowStatus(ctx, request); err != nil {
		return fmt.Errorf("failed to set workflow status: %w", err)
	}

	return nil
}

func (c *Client) SetTaskStatus(ctx context.Context, runID string, status task.Status, output []byte) error {
	request := &controllersvcv1.SetTaskStatusRequest{
		TaskRunId: runID,
		Status:    taskv1.Status(status),
		Output:    output,
	}

	if _, err := c.controller.SetTaskStatus(ctx, request); err != nil {
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

func (c *Client) GetTaskStatus(ctx context.Context, runID string) (task.Status, []byte, error) {
	request := &controllersvcv1.GetTaskStatusRequest{
		TaskRunId: runID,
	}

	response, err := c.controller.GetTaskStatus(ctx, request)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to get task status: %w", err)
	}

	t := response.GetTask()
	return task.Status(t.GetStatus()), t.GetOutput(), nil
}
