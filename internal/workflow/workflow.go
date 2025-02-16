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

	ctxKey struct{}

	Error struct {
		Message      string `json:"message"`
		WorkflowName string `json:"workflowName"`
		RunID        string `json:"runId"`
		TaskRunID    string `json:"taskRunId,omitempty"`
		TaskName     string `json:"taskName,omitempty"`
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
	return context.WithValue(ctx, ctxKey{}, r)
}

func RunFromContext(ctx context.Context) (Run, bool) {
	run, ok := ctx.Value(ctxKey{}).(Run)
	return run, ok
}
