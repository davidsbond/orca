package task

import (
	"context"
	"encoding/json"
	"time"
)

type (
	Task interface {
		Run(ctx context.Context, runID string, input json.RawMessage) (json.RawMessage, error)
		Name() string
	}

	Status int

	Client interface {
		ScheduleTask(ctx context.Context, runID string, name string, params json.RawMessage) (string, error)
		GetTaskRun(ctx context.Context, runID string) (Run, error)
	}

	Run struct {
		ID            string          `json:"id"`
		WorkflowRunID string          `json:"workflowRunId"`
		TaskName      string          `json:"taskName"`
		CreatedAt     time.Time       `json:"createdAt"`
		ScheduledAt   time.Time       `json:"scheduledAt"`
		StartedAt     time.Time       `json:"startedAt"`
		CompletedAt   time.Time       `json:"completedAt"`
		Status        Status          `json:"status"`
		Input         json.RawMessage `json:"input"`
		Output        json.RawMessage `json:"output"`
	}

	ctxKey struct{}

	Error struct {
		Message  string `json:"message"`
		TaskName string `json:"taskName"`
		RunID    string `json:"runId"`
		Panic    bool   `json:"panic,omitempty"`
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

func ClientFromContext(ctx context.Context) (Client, bool) {
	client, ok := ctx.Value(ctxKey{}).(Client)

	return client, ok
}

func ClientToContext(ctx context.Context, client Client) context.Context {
	return context.WithValue(ctx, ctxKey{}, client)
}
