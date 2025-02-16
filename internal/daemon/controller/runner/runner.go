package runner

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/controller/database/workflow"
	worker_api "github.com/davidsbond/orca/internal/daemon/worker/api/worker"
	"github.com/davidsbond/orca/internal/log"
)

type (
	Runner struct {
		workers   WorkerRepository
		tasks     TaskRepository
		workflows WorkflowRepository
	}

	WorkerRepository interface {
		GetWorkerForTask(ctx context.Context, name string) (worker.Worker, error)
		GetWorkerForWorkflow(ctx context.Context, name string) (worker.Worker, error)
	}

	TaskRepository interface {
		Save(ctx context.Context, run task.Run) error
		GetPendingTaskRuns(ctx context.Context) ([]task.Run, error)
	}

	WorkflowRepository interface {
		Save(ctx context.Context, run workflow.Run) error
		GetPendingWorkflowRuns(ctx context.Context) ([]workflow.Run, error)
	}
)

func New(workers WorkerRepository, workflows WorkflowRepository, tasks TaskRepository) *Runner {
	return &Runner{
		workers:   workers,
		tasks:     tasks,
		workflows: workflows,
	}
}

func (s *Runner) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return s.runWorkflows(ctx)
	})
	group.Go(func() error {
		return s.runTasks(ctx)
	})

	return group.Wait()
}

func (s *Runner) runWorkflows(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			runs, err := s.workflows.GetPendingWorkflowRuns(ctx)
			if err != nil {
				return err
			}

			for _, run := range runs {
				if err = s.runWorkflow(ctx, run); err != nil {
					return err
				}
			}
		}
	}
}

func (s *Runner) runTasks(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			runs, err := s.tasks.GetPendingTaskRuns(ctx)
			if err != nil {
				return err
			}

			for _, run := range runs {
				if err = s.runTask(ctx, run); err != nil {
					return err
				}
			}
		}
	}
}

func (s *Runner) runTask(ctx context.Context, run task.Run) error {
	logger := log.FromContext(ctx).With(
		slog.String("workflow_run_id", run.WorkflowRunID),
		slog.String("task_run_id", run.ID),
		slog.String("task_name", run.TaskName),
	)

	w, err := s.workers.GetWorkerForTask(ctx, run.TaskName)
	switch {
	case errors.Is(err, worker.ErrNotFound):
		return nil
	case err != nil:
		return err
	}

	wrk, err := worker_api.Dial(ctx, w.AdvertiseAddress)
	if err != nil {
		return err
	}

	run.WorkerID.Valid = true
	run.WorkerID.V = w.ID
	run.Status = task.StatusScheduled
	run.ScheduledAt.Valid = true
	run.ScheduledAt.V = time.Now()

	logger.With(slog.String("worker_id", w.ID)).InfoContext(ctx, "assigned task to worker")

	if err = s.tasks.Save(ctx, run); err != nil {
		return err
	}

	if err = wrk.RunTask(ctx, run.ID, run.TaskName, run.Input.V.Bytes); err != nil {
		return err
	}

	return wrk.Close()
}

func (s *Runner) runWorkflow(ctx context.Context, run workflow.Run) error {
	logger := log.FromContext(ctx).With(
		slog.String("workflow_run_id", run.ID),
		slog.String("parent_workflow_run_id", run.ParentWorkflowRunID.V),
		slog.String("workflow_name", run.WorkflowName),
	)

	w, err := s.workers.GetWorkerForWorkflow(ctx, run.WorkflowName)
	switch {
	case errors.Is(err, worker.ErrNotFound):
		return nil
	case err != nil:
		return err
	}

	wrk, err := worker_api.Dial(ctx, w.AdvertiseAddress)
	if err != nil {
		return err
	}

	run.WorkerID.Valid = true
	run.WorkerID.V = w.ID
	run.Status = workflow.StatusScheduled
	run.ScheduledAt.Valid = true
	run.ScheduledAt.V = time.Now()

	logger.With(slog.String("worker_id", w.ID)).DebugContext(ctx, "assigned workflow to worker")

	if err = s.workflows.Save(ctx, run); err != nil {
		return err
	}

	if err = wrk.RunWorkflow(ctx, run.ID, run.WorkflowName, run.Input.V.Bytes); err != nil {
		return err
	}

	return wrk.Close()
}
