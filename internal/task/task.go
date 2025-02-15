package task

import (
	"context"
	"encoding/json"
	"time"
)

type (
	Task interface {
		Run(ctx context.Context, input []byte) ([]byte, error)
		Name() string
	}

	Status int

	Client interface {
		ScheduleTask(ctx context.Context, runID string, name string, params []byte) (string, error)
		GetTaskRun(ctx context.Context, runID string) (Run, error)
	}

	Run struct {
		ID            string
		WorkflowRunID string
		TaskName      string
		CreatedAt     time.Time
		ScheduledAt   time.Time
		StartedAt     time.Time
		CompletedAt   time.Time
		Status        Status
		Input         json.RawMessage
		Output        json.RawMessage
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
