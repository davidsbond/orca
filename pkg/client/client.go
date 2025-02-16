package client

import (
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	workflowsvcv1 "github.com/davidsbond/orca/internal/proto/orca/workflow/service/v1"
)

type (
	Client struct {
		conn      *grpc.ClientConn
		workflows workflowsvcv1.WorkflowServiceClient
	}
)

func Dial(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:      conn,
		workflows: workflowsvcv1.NewWorkflowServiceClient(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

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

func formatError(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	return errors.New(st.Message())
}
