package api

import (
	"context"
	"errors"

	"github.com/davidsbond/orca/internal/workflow"

	"github.com/davidsbond/orca/internal/task"
)

type (
	WorkerService struct {
		registry  Registry
		scheduler Scheduler
	}

	Registry interface {
		GetWorkflow(name string) (workflow.Workflow, bool)
		GetTask(name string) (task.Task, bool)
	}

	Scheduler interface {
		ScheduleWorkflow(ctx context.Context, runID string, workflow workflow.Workflow, input []byte) error
		ScheduleTask(ctx context.Context, runID string, task task.Task, input []byte) error
	}
)

var (
	ErrNotFound = errors.New("not found")
)

func NewWorkerService(registry Registry, scheduler Scheduler) *WorkerService {
	return &WorkerService{
		registry:  registry,
		scheduler: scheduler,
	}
}

func (svc *WorkerService) ScheduleWorkflow(ctx context.Context, name string, runID string, input []byte) error {
	wf, ok := svc.registry.GetWorkflow(name)
	if !ok {
		return ErrNotFound
	}

	return svc.scheduler.ScheduleWorkflow(ctx, runID, wf, input)
}

func (svc *WorkerService) ScheduleTask(ctx context.Context, name string, runID string, input []byte) error {
	tsk, ok := svc.registry.GetTask(name)
	if !ok {
		return ErrNotFound
	}

	return svc.scheduler.ScheduleTask(ctx, runID, tsk, input)
}
