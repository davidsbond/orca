package workflow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
	worker_api "github.com/davidsbond/orca/internal/daemon/worker/api/worker"
)

type (
	Handler struct {
		workflows Repository
		tasks     TaskRepository
		workers   WorkerRepository
	}

	Repository interface {
		Get(ctx context.Context, id string) (workflow.Run, error)
		Save(ctx context.Context, run workflow.Run) error
		ListForWorkflowRun(ctx context.Context, id string) ([]workflow.Run, error)
	}

	TaskRepository interface {
		ListForWorkflowRun(ctx context.Context, runID string) ([]task.Run, error)
	}

	WorkerRepository interface {
		Get(ctx context.Context, id string) (worker.Worker, error)
	}
)

var (
	ErrNotFound     = errors.New("not found")
	ErrCannotCancel = errors.New("cannot cancel")
)

func NewHandler(workflows Repository, tasks TaskRepository, workers WorkerRepository) *Handler {
	return &Handler{
		workflows: workflows,
		tasks:     tasks,
		workers:   workers,
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

func (h *Handler) CancelRun(ctx context.Context, runID string) error {
	run, err := h.workflows.Get(ctx, runID)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return ErrNotFound
	case err != nil:
		return err
	}

	if run.Status == workflow.StatusSkipped ||
		run.Status == workflow.StatusComplete ||
		run.Status == workflow.StatusFailed ||
		run.Status == workflow.StatusTimeout ||
		run.Status == workflow.StatusCancelled {
		return ErrCannotCancel
	}

	// If this workflow has been assigned and sent to a worker, we'll call the worker to cancel the
	// workflow if its in flight.
	if run.WorkerID.Valid {
		wrk, err := h.workers.Get(ctx, run.WorkerID.V)
		if err != nil {
			return err
		}

		client, err := worker_api.Dial(ctx, wrk.AdvertiseAddress)
		if err != nil {
			return fmt.Errorf("failed to dial worker: %w", err)
		}

		return errors.Join(
			client.CancelWorkflowRun(ctx, runID),
			client.Close(),
		)
	}

	// If this workflow has yet to be sent to a worker, we can just update its status
	// locally, so it won't be picked up by a controller
	run.Status = workflow.StatusCancelled
	return h.workflows.Save(ctx, run)
}
