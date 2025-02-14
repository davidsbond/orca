package task

import "context"

type (
	Task interface {
		Run(ctx context.Context, input []byte) ([]byte, error)
		Name() string
	}

	Status int

	Client interface {
		ScheduleTask(ctx context.Context, runID string, name string, params []byte) (string, error)
		GetTaskStatus(ctx context.Context, runID string) (Status, []byte, error)
	}

	ctxKey struct{}
)

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusScheduled
	StatusRunning
	StatusComplete
	StatusFailed
)

func ClientFromContext(ctx context.Context) (Client, bool) {
	client, ok := ctx.Value(ctxKey{}).(Client)

	return client, ok
}

func ClientToContext(ctx context.Context, client Client) context.Context {
	return context.WithValue(ctx, ctxKey{}, client)
}
