package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	workersvcv1 "github.com/davidsbond/orca/internal/proto/orca/worker/service/v1"
)

type (
	Client struct {
		conn       *grpc.ClientConn
		controller workersvcv1.WorkerServiceClient
	}
)

func Dial(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func (c *Client) RunTask(ctx context.Context, runID string, name string, input []byte) error {
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

func (c *Client) RunWorkflow(ctx context.Context, runID string, name string, input []byte) error {
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
