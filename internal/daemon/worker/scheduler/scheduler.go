package scheduler

import (
	"context"
	"encoding/json"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/task"
	"github.com/davidsbond/orca/internal/workflow"
)

type (
	Scheduler struct {
		workflows  chan *scheduledWorkflow
		tasks      chan *scheduledTask
		controller ControllerClient
	}

	scheduledWorkflow struct {
		runID    string
		workflow workflow.Workflow
		input    json.RawMessage
	}

	scheduledTask struct {
		runID string
		task  task.Task
		input json.RawMessage
	}

	ControllerClient interface {
		SetWorkflowRunStatus(ctx context.Context, runID string, status workflow.Status, output json.RawMessage) error
		SetTaskRunStatus(ctx context.Context, runID string, status task.Status, output json.RawMessage) error
		ScheduleTask(ctx context.Context, runID string, name string, params json.RawMessage) (string, error)
		GetTaskRun(ctx context.Context, runID string) (task.Run, error)
		GetWorkflowRun(ctx context.Context, runID string) (workflow.Run, error)
	}
)

func New(controller ControllerClient) *Scheduler {
	return &Scheduler{
		controller: controller,
		workflows:  make(chan *scheduledWorkflow, 1024),
		tasks:      make(chan *scheduledTask, 1024),
	}
}

func (s *Scheduler) Run(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return s.runWorkflows(ctx)
	})
	group.Go(func() error {
		return s.runTasks(ctx)
	})

	return group.Wait()
}

func (s *Scheduler) ScheduleWorkflow(ctx context.Context, id string, wf workflow.Workflow, input json.RawMessage) error {
	sw := &scheduledWorkflow{
		runID:    id,
		workflow: wf,
		input:    input,
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.workflows <- sw:
		return nil
	}
}

func (s *Scheduler) ScheduleTask(ctx context.Context, id string, t task.Task, input json.RawMessage) error {
	st := &scheduledTask{
		runID: id,
		task:  t,
		input: input,
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.tasks <- st:
		return nil
	}
}

func (s *Scheduler) runWorkflows(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sw := <-s.workflows:
			if err := s.runWorkflow(ctx, sw); err != nil {
				return err
			}
		}
	}
}

func (s *Scheduler) runTasks(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case st := <-s.tasks:
			if err := s.runTask(ctx, st); err != nil {
				return err
			}
		}
	}
}

func (s *Scheduler) runWorkflow(ctx context.Context, sw *scheduledWorkflow) error {
	run, err := s.controller.GetWorkflowRun(ctx, sw.runID)
	if err != nil {
		return err
	}

	if err = s.controller.SetWorkflowRunStatus(ctx, run.ID, workflow.StatusRunning, nil); err != nil {
		return err
	}

	ctx = workflow.RunToContext(ctx, run)
	ctx = task.ClientToContext(ctx, s.controller)

	output, err := sw.workflow.Run(ctx, sw.input)
	if err != nil {
		if err = s.controller.SetWorkflowRunStatus(ctx, sw.runID, workflow.StatusFailed, json.RawMessage(err.Error())); err != nil {
			return err
		}

		return nil
	}

	if err = s.controller.SetWorkflowRunStatus(ctx, sw.runID, workflow.StatusComplete, output); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) runTask(ctx context.Context, st *scheduledTask) error {
	if err := s.controller.SetTaskRunStatus(ctx, st.runID, task.StatusRunning, nil); err != nil {
		return err
	}

	output, err := st.task.Run(ctx, st.input)
	if err != nil {
		if err = s.controller.SetTaskRunStatus(ctx, st.runID, task.StatusFailed, json.RawMessage(err.Error())); err != nil {
			return err
		}

		return nil
	}

	if err = s.controller.SetTaskRunStatus(ctx, st.runID, task.StatusComplete, output); err != nil {
		return err
	}

	return nil
}
