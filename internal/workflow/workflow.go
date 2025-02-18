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
		ID           string          `json:"id"`
		ParentID     string          `json:"parentId"`
		WorkflowName string          `json:"workflowName"`
		CreatedAt    time.Time       `json:"createdAt"`
		ScheduledAt  time.Time       `json:"scheduledAt"`
		StartedAt    time.Time       `json:"startedAt"`
		CompletedAt  time.Time       `json:"completedAt"`
		Status       Status          `json:"status"`
		Input        json.RawMessage `json:"input"`
		Output       json.RawMessage `json:"output"`
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

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "PENDING"
	case StatusScheduled:
		return "SCHEDULED"
	case StatusRunning:
		return "RUNNING"
	case StatusComplete:
		return "COMPLETE"
	case StatusFailed:
		return "FAILED"
	default:
		return "UNKNOWN"
	}
}

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
