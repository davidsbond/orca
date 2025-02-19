package workflow

import (
	"context"
	"encoding/json"

	"github.com/davidsbond/orca/internal/task"
)

type (
	Client interface {
		ScheduleWorkflow(ctx context.Context, params ScheduleWorkflowParams) (string, error)
		GetWorkflowRun(ctx context.Context, runID string) (Run, error)
		ScheduleTask(ctx context.Context, params ScheduleTaskParams) (string, error)
		GetTaskRun(ctx context.Context, runID string) (task.Run, error)
	}

	ScheduleWorkflowParams struct {
		WorkflowRunID string
		WorkflowName  string
		IdempotentKey string
		Input         json.RawMessage
	}

	ScheduleTaskParams struct {
		WorkflowRunID string
		TaskName      string
		IdempotentKey string
		Input         json.RawMessage
	}
)
