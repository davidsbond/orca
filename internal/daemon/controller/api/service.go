package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
)

type (
	ControllerService struct {
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

func NewControllerService(workers WorkerRepository, tasks TaskRepository, workflows WorkflowRepository) *ControllerService {
	return &ControllerService{
		workers:   workers,
		tasks:     tasks,
		workflows: workflows,
	}
}

func (svc *ControllerService) RegisterWorker(ctx context.Context, w worker.Worker) error {
	err := svc.workers.Save(ctx, w)
	switch {
	case err != nil:
		return err
	default:
		return nil
	}
}

func (svc *ControllerService) DeregisterWorker(ctx context.Context, id string) error {
	err := svc.workers.Delete(ctx, id)
	switch {
	case errors.Is(err, worker.ErrNotFound):
		return ErrWorkerNotFound
	case err != nil:
		return err
	default:
		return nil
	}
}

func (svc *ControllerService) ScheduleTask(ctx context.Context, id string, name string, input []byte) (string, error) {
	run := task.Run{
		ID:            uuid.NewString(),
		WorkflowRunID: id,
		TaskName:      name,
		CreatedAt:     time.Now(),
		Status:        task.StatusPending,
		Input:         input,
	}

	err := svc.tasks.Save(ctx, run)
	switch {
	case errors.Is(err, task.ErrWorkflowRunNotFound):
		return "", ErrWorkflowRunNotFound
	case err != nil:
		return "", fmt.Errorf("failed to save task run: %w", err)
	default:
		return run.ID, nil
	}
}

func (svc *ControllerService) SetTaskRunStatus(ctx context.Context, id string, status task.Status, output []byte) error {
	run, err := svc.tasks.Get(ctx, id)
	switch {
	case errors.Is(err, task.ErrNotFound):
		return ErrTaskRunNotFound
	case err != nil:
		return fmt.Errorf("failed to get task run: %w", err)
	}

	run.Status = status

	if output != nil {
		run.Output = output
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

	err = svc.tasks.Save(ctx, run)
	switch {
	case errors.Is(err, task.ErrWorkflowRunNotFound):
		return ErrWorkflowRunNotFound
	case err != nil:
		return fmt.Errorf("failed to save task run: %w", err)
	default:
		return nil
	}
}

func (svc *ControllerService) GetTaskRun(ctx context.Context, id string) (task.Run, error) {
	run, err := svc.tasks.Get(ctx, id)
	switch {
	case errors.Is(err, task.ErrNotFound):
		return task.Run{}, ErrTaskRunNotFound
	case err != nil:
		return task.Run{}, fmt.Errorf("failed to get task run: %w", err)
	default:
		return run, nil
	}
}

func (svc *ControllerService) ScheduleWorkflow(ctx context.Context, parentWorkflowRunID string, name string, input []byte) (string, error) {
	run := workflow.Run{
		ID:           uuid.NewString(),
		WorkflowName: name,
		CreatedAt:    time.Now(),
		Status:       workflow.StatusPending,
		Input:        input,
	}

	if parentWorkflowRunID != "" {
		run.ParentWorkflowRunID.V = parentWorkflowRunID
		run.ParentWorkflowRunID.Valid = true
	}

	err := svc.workflows.Save(ctx, run)
	switch {
	case errors.Is(err, workflow.ErrParentWorkflowNotFound):
		return "", ErrWorkflowRunNotFound
	case err != nil:
		return "", fmt.Errorf("failed to save workflow run: %w", err)
	default:
		return run.ID, nil
	}
}

func (svc *ControllerService) SetWorkflowRunStatus(ctx context.Context, id string, status workflow.Status, output []byte) error {
	run, err := svc.workflows.Get(ctx, id)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return ErrWorkerNotFound
	case err != nil:
		return fmt.Errorf("failed to get workflow run: %w", err)
	}

	run.Status = status

	if output != nil {
		run.Output = output
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

	err = svc.workflows.Save(ctx, run)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return ErrWorkflowRunNotFound
	case err != nil:
		return fmt.Errorf("failed to save workflow run: %w", err)
	default:
		return nil
	}
}

func (svc *ControllerService) GetWorkflowRun(ctx context.Context, id string) (workflow.Run, error) {
	run, err := svc.workflows.Get(ctx, id)
	switch {
	case errors.Is(err, workflow.ErrNotFound):
		return workflow.Run{}, ErrWorkflowRunNotFound
	case err != nil:
		return workflow.Run{}, fmt.Errorf("failed to get workflow run: %w", err)
	default:
		return run, nil
	}
}
