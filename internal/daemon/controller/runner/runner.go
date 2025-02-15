package runner

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/daemon/controller/database/task"
	"github.com/davidsbond/orca/internal/daemon/controller/database/worker"
	"github.com/davidsbond/orca/internal/daemon/worker/client"
	"github.com/davidsbond/orca/internal/workflow"
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
	w, err := s.workers.GetWorkerForTask(ctx, run.TaskName)
	switch {
	case errors.Is(err, worker.ErrNotFound):
		return nil
	case err != nil:
		return err
	}

	workerClient, err := client.Dial(w.AdvertiseAddress)
	if err != nil {
		return err
	}

	return errors.Join(
		workerClient.RunTask(ctx, run.ID, run.TaskName, run.Input),
		workerClient.Close(),
	)
}

func (s *Runner) runWorkflow(ctx context.Context, run workflow.Run) error {
	w, err := s.workers.GetWorkerForWorkflow(ctx, run.WorkflowName)
	switch {
	case errors.Is(err, worker.ErrNotFound):
		return nil
	case err != nil:
		return err
	}

	workerClient, err := client.Dial(w.AdvertiseAddress)
	if err != nil {
		return err
	}

	return errors.Join(
		workerClient.RunWorkflow(ctx, run.ID, run.WorkflowName, run.Input),
		workerClient.Close(),
	)
}
