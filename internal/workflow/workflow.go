package workflow

import (
	"context"
	"encoding/json"
	"time"
)

type (
	Workflow interface {
		Run(ctx context.Context, runID string, input json.RawMessage) (json.RawMessage, error)
		Name() string
	}

	Status int

	Run struct {
		ID           string
		ParentID     string
		WorkflowName string
		CreatedAt    time.Time
		ScheduledAt  time.Time
		StartedAt    time.Time
		CompletedAt  time.Time
		Status       Status
		Input        json.RawMessage
		Output       json.RawMessage
	}

	Client interface {
		ScheduleWorkflow(ctx context.Context, runID string, name string, params json.RawMessage) (string, error)
		GetWorkflowRun(ctx context.Context, runID string) (Run, error)
	}

	runCtxKey    struct{}
	clientCtxKey struct{}

	Error struct {
		Message      string `json:"message"`
		WorkflowName string `json:"workflowName"`
		RunID        string `json:"runId"`
		TaskRunID    string `json:"taskRunId,omitempty"`
		TaskName     string `json:"taskName,omitempty"`
		Panic        bool   `json:"panic,omitempty"`
	}
)

func (e Error) Error() string {
	return e.Message
}

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusScheduled
	StatusRunning
	StatusComplete
	StatusFailed
)

func RunToContext(ctx context.Context, r Run) context.Context {
	return context.WithValue(ctx, runCtxKey{}, r)
}

func RunFromContext(ctx context.Context) (Run, bool) {
	run, ok := ctx.Value(runCtxKey{}).(Run)
	return run, ok
}

func ClientFromContext(ctx context.Context) (Client, bool) {
	client, ok := ctx.Value(clientCtxKey{}).(Client)

	return client, ok
}

func ClientToContext(ctx context.Context, client Client) context.Context {
	return context.WithValue(ctx, clientCtxKey{}, client)
}
