package workflow

import (
	"encoding/json"
	"time"
)

type (
	Workflow interface {
		Run(ctx *Context, runID string, input json.RawMessage) (json.RawMessage, error)
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
		CancelledAt  time.Time       `json:"cancelledAt"`
		Status       Status          `json:"status"`
		Input        json.RawMessage `json:"input"`
		Output       json.RawMessage `json:"output"`
	}

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
	StatusSkipped
	StatusTimeout
	StatusCancelled
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
	case StatusCancelled:
		return "CANCELLED"
	default:
		return "UNKNOWN"
	}
}
