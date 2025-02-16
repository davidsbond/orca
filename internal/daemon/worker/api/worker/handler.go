package worker

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/davidsbond/orca/internal/workflow"

	"github.com/davidsbond/orca/internal/task"
)

type (
	Handler struct {
		registry  Registry
		scheduler Scheduler
	}

	Registry interface {
		GetWorkflow(name string) (workflow.Workflow, bool)
		GetTask(name string) (task.Task, bool)
	}

	Scheduler interface {
		ScheduleWorkflow(ctx context.Context, runID string, workflow workflow.Workflow, input json.RawMessage) error
		ScheduleTask(ctx context.Context, runID string, task task.Task, input json.RawMessage) error
	}
)

var (
	ErrNotFound = errors.New("not found")
)

func NewHandler(registry Registry, scheduler Scheduler) *Handler {
	return &Handler{
		registry:  registry,
		scheduler: scheduler,
	}
}

func (svc *Handler) ScheduleWorkflow(ctx context.Context, name string, runID string, input json.RawMessage) error {
	wf, ok := svc.registry.GetWorkflow(name)
	if !ok {
		return ErrNotFound
	}

	return svc.scheduler.ScheduleWorkflow(ctx, runID, wf, input)
}

func (svc *Handler) ScheduleTask(ctx context.Context, name string, runID string, input json.RawMessage) error {
	tsk, ok := svc.registry.GetTask(name)
	if !ok {
		return ErrNotFound
	}

	return svc.scheduler.ScheduleTask(ctx, runID, tsk, input)
}
