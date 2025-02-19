package task

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

type (
	Task interface {
		Run(ctx context.Context, runID string, input json.RawMessage) (json.RawMessage, error)
		Name() string
	}

	Status int

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

var (
	ErrTimeout = errors.New("task timeout")
)

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusScheduled
	StatusRunning
	StatusComplete
	StatusFailed
	StatusSkipped
	StatusTimeout
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
	case StatusSkipped:
		return "SKIPPED"
	case StatusTimeout:
		return "TIMEOUT"
	default:
		return "UNKNOWN"
	}
}
