package worker

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"

	"github.com/davidsbond/orca/internal/daemon"
	workersvcv1 "github.com/davidsbond/orca/internal/proto/orca/private/worker/service/v1"
)

type (
	Client struct {
		conn       *grpc.ClientConn
		controller workersvcv1.WorkerServiceClient
	}
)

func Dial(ctx context.Context, addr string) (*Client, error) {
	conn, err := daemon.Dial(ctx, addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:       conn,
		controller: workersvcv1.NewWorkerServiceClient(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) RunTask(ctx context.Context, runID string, name string, input json.RawMessage) error {
	request := &workersvcv1.RunTaskRequest{
		TaskRunId: runID,
		TaskName:  name,
		Input:     input,
	}

	if _, err := c.controller.RunTask(ctx, request); err != nil {
		return err
	}

	return nil
}

func (c *Client) RunWorkflow(ctx context.Context, runID string, name string, input json.RawMessage) error {
	request := &workersvcv1.RunWorkflowRequest{
		WorkflowRunId: runID,
		WorkflowName:  name,
		Input:         input,
	}

	if _, err := c.controller.RunWorkflow(ctx, request); err != nil {
		return err
	}

	return nil
}

func (c *Client) CancelWorkflowRun(ctx context.Context, runID string) error {
	request := &workersvcv1.CancelWorkflowRunRequest{
		WorkflowRunId: runID,
	}

	if _, err := c.controller.CancelWorkflowRun(ctx, request); err != nil {
		return err
	}

	return nil
}
