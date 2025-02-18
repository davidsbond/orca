package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"

	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
)

type (
	Handler struct {
		workers   WorkerRepository
		tasks     TaskRepository
		workflows WorkflowRepository
	}

	WorkerRepository interface {
		Save(ctx context.Context, worker worker.Worker) error
		Delete(ctx context.Context, id string) error
	}

	TaskRepository interface {
		Save(ctx context.Context, run task.Run) error
		Get(ctx context.Context, id string) (task.Run, error)
		FindByIdempotentKey(ctx context.Context, name string, key string) (task.Run, error)
	}

	WorkflowRepository interface {
		Save(ctx context.Context, run workflow.Run) error
		Get(ctx context.Context, id string) (workflow.Run, error)
		FindByIdempotentKey(ctx context.Context, name string, key string) (workflow.Run, error)
	}
)

var (
	ErrWorkerNotFound         = errors.New("worker not found")
	ErrWorkflowRunNotFound    = errors.New("workflow run not found")
	ErrTaskRunNotFound        = errors.New("task run not found")
	ErrParentWorkflowNotFound = errors.New("parent workflow not found")
)

func NewHandler(workers WorkerRepository, tasks TaskRepository, workflows WorkflowRepository) *Handler {
	return &Handler{
		workers:   workers,
		tasks:     tasks,
		workflows: workflows,
	}
}

func (h *Handler) RegisterWorker(ctx context.Context, w worker.Worker) error {
	err := h.workers.Save(ctx, w)
	switch {
	case err != nil:
		return err
	default:
		return nil
	}
}

func (h *Handler) DeregisterWorker(ctx context.Context, id string) error {
	err := h.workers.Delete(ctx, id)
	switch {
	case errors.Is(err, worker.ErrNotFound):
		return ErrWorkerNotFound
	case err != nil:
		return err
	default:
		return nil
	}
}

type (
	ScheduleTaskParams struct {
		WorkflowRunID string
		TaskName      string
		Input         json.RawMessage
		IdempotentKey string
	}
)

func (h *Handler) ScheduleTask(ctx context.Context, params ScheduleTaskParams) (string, error) {
	run := task.Run{
		ID:            uuid.NewString(),
		WorkflowRunID: params.WorkflowRunID,
		TaskName:      params.TaskName,
		CreatedAt:     time.Now(),
		Status:        task.StatusPending,
	}

	if len(params.Input) > 0 {
		run.Input.Valid = true
		run.Input.V = pgtype.JSONB{
			Bytes:  params.Input,
			Status: pgtype.Present,
		}
	}

	if params.IdempotentKey != "" {
		run.IdempotentKey.Valid = true
		run.IdempotentKey.V = params.IdempotentKey

		existing, err := h.tasks.FindByIdempotentKey(ctx, params.TaskName, params.IdempotentKey)
		switch {
		case errors.Is(err, task.ErrNotFound):
			break
		case err != nil:
			return "", fmt.Errorf("failed to lookup idempotent key: %w", err)
		default:
			// If the key has been used before, set the status to skipped and fill out the output.
			run.Status = task.StatusSkipped
			run.Output = existing.Output
			run.CompletedAt.Valid = true
			run.CompletedAt.V = run.CreatedAt
		}
	}

	err := h.tasks.Save(ctx, run)
	switch {
	case errors.Is(err, task.ErrWorkflowRunNotFound):
		return "", ErrWorkflowRunNotFound
	case err != nil:
		return "", fmt.Errorf("failed to save task run: %w", err)
	default:
		return run.ID, nil
	}
}

func (h *Handler) SetTaskRunStatus(ctx context.Context, id string, status task.Status, output json.RawMessage) error {
	run, err := h.tasks.Get(ctx, id)
	switch {
	case errors.Is(err, task.ErrNotFound):
		return ErrTaskRunNotFound
	case err != nil:
		return err
	}

	// Don't allow further updates to task runs in one of the final stages.
	if run.Status == task.StatusSkipped || run.Status == task.StatusComplete || run.Status == task.StatusFailed {
		return nil
	}

	run.Status = status

	if len(output) > 0 {
		run.Output.Valid = true
		run.Output.V = pgtype.JSONB{
			Bytes:  output,
			Status: pgtype.Present,
		}
	}

	if run.Status == task.StatusScheduled && !run.ScheduledAt.Valid {
		run.ScheduledAt.Valid = true
		run.ScheduledAt.V = time.Now()
	}

	if run.Status == task.StatusRunning && !run.StartedAt.Valid {
		run.StartedAt.Valid = true
		run.StartedAt.V = time.Now()
	}

	if run.Status == task.StatusComplete && !run.CompletedAt.Valid {
		run.CompletedAt.Valid = true
		run.CompletedAt.V = time.Now()
	}

	if run.Status == task.StatusFailed && !run.CompletedAt.Valid {
		run.CompletedAt.Valid = true
		run.CompletedAt.V = time.Now()
	}

	err = h.tasks.Save(ctx, run)
	switch {
	case errors.Is(err, task.ErrWorkflowRunNotFound):
		return ErrWorkflowRunNotFound
	case err != nil:
		return fmt.Errorf("failed to save task run: %w", err)
	default:
		return nil
	}
}

func (h *Handler) GetTaskRun(ctx context.Context, id string) (task.Run, error) {
	run, err := h.tasks.Get(ctx, id)
	switch {
	case errors.Is(err, task.ErrNotFound):
		return task.Run{}, ErrTaskRunNotFound
	case err != nil:
		return task.Run{}, err
	default:
		return run, nil
	}
}

type (
	ScheduleWorkflowParams struct {
		ParentWorkflowRunID string
		WorkflowName        string
		Input               json.RawMessage
		IdempotentKey       string
	}
)

func (h *Handler) ScheduleWorkflow(ctx context.Context, params ScheduleWorkflowParams) (string, error) {
	run := workflow.Run{
		ID:           uuid.NewString(),
		WorkflowName: params.WorkflowName,
		CreatedAt:    time.Now(),
		Status:       workflow.StatusPending,
	}

	if len(params.Input) > 0 {
		run.Input.Valid = true
		run.Input.V = pgtype.JSONB{
			Bytes:  params.Input,
			Status: pgtype.Present,
		}
	}

	if params.ParentWorkflowRunID != "" {
		run.ParentWorkflowRunID.V = params.ParentWorkflowRunID
		run.ParentWorkflowRunID.Valid = true
	}

	if params.IdempotentKey != "" {
		run.IdempotentKey.Valid = true
		run.IdempotentKey.V = params.IdempotentKey

		existing, err := h.workflows.FindByIdempotentKey(ctx, params.WorkflowName, params.IdempotentKey)
		switch {
		case errors.Is(err, workflow.ErrNotFound):
			break
		case err != nil:
			return "", fmt.Errorf("failed to lookup idempotent key: %w", err)
		default:
			// If the key has been used before, set the status to skipped and fill out the output.
			run.Status = workflow.StatusSkipped
			run.Output = existing.Output
			run.CompletedAt.Valid = true
			run.CompletedAt.V = run.CreatedAt
		}
	}

	err := h.workflows.Save(ctx, run)
	switch {
	case errors.Is(err, workflow.ErrParentWorkflowNotFound):
		return "", ErrWorkflowRunNotFound
	case err != nil:
		return "", err
	default:
		return run.ID, nil
	}
}

func (h *Handler) SetWorkflowRunStatus(ctx context.Context, id string, status workflow.Status, output json.RawMessage) error {
	run, err := h.workflows.Get(ctx, id)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return ErrWorkerNotFound
	case err != nil:
		return err
	}

	// Don't allow further updates to workflow runs in one of the final stages.
	if run.Status == workflow.StatusSkipped || run.Status == workflow.StatusComplete || run.Status == workflow.StatusFailed {
		return nil
	}

	run.Status = status

	if len(output) > 0 {
		run.Output.Valid = true
		run.Output.V = pgtype.JSONB{
			Bytes:  output,
			Status: pgtype.Present,
		}
	}

	if run.Status == workflow.StatusScheduled && !run.ScheduledAt.Valid {
		run.ScheduledAt.Valid = true
		run.ScheduledAt.V = time.Now()
	}

	if run.Status == workflow.StatusRunning && !run.StartedAt.Valid {
		run.StartedAt.Valid = true
		run.StartedAt.V = time.Now()
	}

	if run.Status == workflow.StatusComplete && !run.CompletedAt.Valid {
		run.CompletedAt.Valid = true
		run.CompletedAt.V = time.Now()
	}

	if run.Status == workflow.StatusFailed && !run.CompletedAt.Valid {
		run.CompletedAt.Valid = true
		run.CompletedAt.V = time.Now()
	}

	err = h.workflows.Save(ctx, run)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return ErrWorkflowRunNotFound
	case err != nil:
		return err
	default:
		return nil
	}
}

func (h *Handler) GetWorkflowRun(ctx context.Context, id string) (workflow.Run, error) {
	run, err := h.workflows.Get(ctx, id)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return workflow.Run{}, ErrWorkflowRunNotFound
	case err != nil:
		return workflow.Run{}, err
	default:
		return run, nil
	}
}
