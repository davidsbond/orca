package workflow

import (
	"context"
	"encoding/json"
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
)

type (
	Handler struct {
		workflows Repository
		tasks     TaskRepository
	}

	Repository interface {
		Get(ctx context.Context, id string) (workflow.Run, error)
		Save(ctx context.Context, run workflow.Run) error
		ListForWorkflowRun(ctx context.Context, id string) ([]workflow.Run, error)
	}

	TaskRepository interface {
		ListForWorkflowRun(ctx context.Context, runID string) ([]task.Run, error)
	}
)

var (
	ErrNotFound = errors.New("not found")
)

func NewHandler(workflows Repository, tasks TaskRepository) *Handler {
	return &Handler{
		workflows: workflows,
		tasks:     tasks,
	}
}

func (h *Handler) ScheduleWorkflow(ctx context.Context, name string, input json.RawMessage) (string, error) {
	run := workflow.Run{
		ID:           uuid.NewString(),
		WorkflowName: name,
		CreatedAt:    time.Now(),
		Status:       workflow.StatusPending,
	}

	if len(input) > 0 {
		run.Input.Valid = true
		run.Input.V = pgtype.JSONB{
			Bytes:  input,
			Status: pgtype.Present,
		}
	}

	if err := h.workflows.Save(ctx, run); err != nil {
		return "", err
	}

	return run.ID, nil
}

func (h *Handler) GetRun(ctx context.Context, runID string) (workflow.Run, error) {
	run, err := h.workflows.Get(ctx, runID)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return workflow.Run{}, ErrNotFound
	case err != nil:
		return workflow.Run{}, err
	default:
		return run, nil
	}
}

func (h *Handler) DescribeRun(ctx context.Context, runID string) (Description, error) {
	root, err := h.workflows.Get(ctx, runID)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return Description{}, ErrNotFound
	case err != nil:
		return Description{}, err
	}

	description := Description{
		Run: root,
	}

	taskRuns, err := h.tasks.ListForWorkflowRun(ctx, runID)
	if err != nil {
		return Description{}, err
	}

	for _, taskRun := range taskRuns {
		description.Actions = append(description.Actions, Action{
			TaskRun: &taskRun,
		})
	}

	workflowRuns, err := h.workflows.ListForWorkflowRun(ctx, runID)
	if err != nil {
		return Description{}, err
	}

	for _, workflowRun := range workflowRuns {
		d, err := h.DescribeRun(ctx, workflowRun.ID)
		if err != nil {
			return Description{}, err
		}

		description.Actions = append(description.Actions, Action{
			WorkflowRun: &d,
		})
	}

	slices.SortFunc(description.Actions, func(a, b Action) int {
		var (
			aCreated time.Time
			bCreated time.Time
		)

		if a.WorkflowRun != nil {
			aCreated = a.WorkflowRun.Run.CreatedAt
		}

		if b.WorkflowRun != nil {
			bCreated = b.WorkflowRun.Run.CreatedAt
		}

		if a.TaskRun != nil {
			aCreated = a.TaskRun.CreatedAt
		}

		if b.TaskRun != nil {
			bCreated = b.TaskRun.CreatedAt
		}

		if aCreated.Before(bCreated) {
			return -1
		}

		if bCreated.Before(aCreated) {
			return 1
		}

		return 0
	})

	return description, nil
}
