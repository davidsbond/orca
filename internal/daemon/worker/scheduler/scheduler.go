package scheduler

import (
	"context"

	"github.com/davidsbond/orca/internal/daemon/worker/task"
	"github.com/davidsbond/orca/internal/daemon/worker/workflow"

	"golang.org/x/sync/errgroup"
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
		input    []byte
	}

	scheduledTask struct {
		runID string
		task  task.Task
		input []byte
	}

	ControllerClient interface {
		SetWorkflowStatus(ctx context.Context, runID string, status workflow.Status, output []byte) error
		SetTaskStatus(ctx context.Context, runID string, status task.Status, output []byte) error
		ScheduleTask(ctx context.Context, runID string, name string, params []byte) (string, error)
		GetTaskStatus(ctx context.Context, runID string) (task.Status, []byte, error)
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

func (s *Scheduler) ScheduleWorkflow(ctx context.Context, id string, wf workflow.Workflow, input []byte) error {
	sw := &scheduledWorkflow{
		runID:    id,
		workflow: wf,
		input:    input,
	}

	if err := s.controller.SetWorkflowStatus(ctx, sw.runID, workflow.StatusScheduled, nil); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.workflows <- sw:
		return nil
	}
}

func (s *Scheduler) ScheduleTask(ctx context.Context, id string, t task.Task, input []byte) error {
	st := &scheduledTask{
		runID: id,
		task:  t,
		input: input,
	}

	if err := s.controller.SetTaskStatus(ctx, st.runID, task.StatusScheduled, nil); err != nil {
		return err
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
	if err := s.controller.SetWorkflowStatus(ctx, sw.runID, workflow.StatusRunning, nil); err != nil {
		return err
	}

	ctx = workflow.RunToContext(ctx, workflow.Run{ID: sw.runID})

	output, err := sw.workflow.Run(ctx, sw.input)
	if err != nil {
		if err = s.controller.SetWorkflowStatus(ctx, sw.runID, workflow.StatusFailed, []byte(err.Error())); err != nil {
			return err
		}

		return nil
	}

	if err = s.controller.SetWorkflowStatus(ctx, sw.runID, workflow.StatusComplete, output); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) runTask(ctx context.Context, st *scheduledTask) error {
	if err := s.controller.SetTaskStatus(ctx, st.runID, task.StatusRunning, nil); err != nil {
		return err
	}

	ctx = task.ClientToContext(ctx, s.controller)

	output, err := st.task.Run(ctx, st.input)
	if err != nil {
		if err = s.controller.SetTaskStatus(ctx, st.runID, task.StatusFailed, []byte(err.Error())); err != nil {
			return err
		}

		return nil
	}

	if err = s.controller.SetTaskStatus(ctx, st.runID, task.StatusComplete, output); err != nil {
		return err
	}

	return nil
}
