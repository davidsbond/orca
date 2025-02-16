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
	}

	WorkflowRepository interface {
		Save(ctx context.Context, run workflow.Run) error
		Get(ctx context.Context, id string) (workflow.Run, error)
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

func (h *Handler) ScheduleTask(ctx context.Context, id string, name string, input json.RawMessage) (string, error) {
	run := task.Run{
		ID:            uuid.NewString(),
		WorkflowRunID: id,
		TaskName:      name,
		CreatedAt:     time.Now(),
		Status:        task.StatusPending,
	}

	if len(input) > 0 {
		run.Input.Valid = true
		run.Input.V = pgtype.JSONB{
			Bytes:  input,
			Status: pgtype.Present,
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

func (h *Handler) ScheduleWorkflow(ctx context.Context, parentWorkflowRunID string, name string, input json.RawMessage) (string, error) {
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

	if parentWorkflowRunID != "" {
		run.ParentWorkflowRunID.V = parentWorkflowRunID
		run.ParentWorkflowRunID.Valid = true
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
