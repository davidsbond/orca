package client

import (
	"context"
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

var (
	ErrNotFound = errors.New("not found")
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

type (
	ctxKey struct{}
)

func ToContext(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, ctxKey{}, client)
}

func FromContext(ctx context.Context) (*Client, bool) {
	client, ok := ctx.Value(ctxKey{}).(*Client)
	return client, ok
}
