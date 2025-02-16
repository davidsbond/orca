package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"golang.org/x/sync/errgroup"

	"github.com/davidsbond/orca/internal/log"
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
		ScheduleTask(ctx context.Context, runID string, name string, input json.RawMessage) (string, error)
		GetTaskRun(ctx context.Context, runID string) (task.Run, error)
		GetWorkflowRun(ctx context.Context, runID string) (workflow.Run, error)
		ScheduleWorkflow(ctx context.Context, runID string, name string, input json.RawMessage) (string, error)
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
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(10)

	for {
		select {
		case <-ctx.Done():
			return group.Wait()
		case sw := <-s.workflows:
			group.Go(func() error {
				return s.runWorkflow(ctx, sw)
			})
		}
	}
}

func (s *Scheduler) runTasks(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(10)

	for {
		select {
		case <-ctx.Done():
			return group.Wait()
		case st := <-s.tasks:
			group.Go(func() error {
				return s.runTask(ctx, st)
			})
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
	ctx = workflow.ClientToContext(ctx, s.controller)
	ctx = task.ClientToContext(ctx, s.controller)

	logger := log.FromContext(ctx).With(
		slog.String("workflow_run_id", sw.runID),
		slog.String("workflow_name", sw.workflow.Name()),
	)

	logger.InfoContext(ctx, "running workflow")
	output, err := sw.workflow.Run(ctx, sw.runID, sw.input)
	if err != nil {
		var e workflow.Error
		if !errors.As(err, &e) {
			e = workflow.Error{
				Message:  err.Error(),
				TaskName: sw.workflow.Name(),
				RunID:    sw.runID,
			}
		}

		output, err = json.Marshal(e)
		if err != nil {
			return err
		}

		logger.With(slog.String("error", e.Message)).ErrorContext(ctx, "workflow run failed")
		if err = s.controller.SetWorkflowRunStatus(ctx, sw.runID, workflow.StatusFailed, output); err != nil {
			return err
		}

		return nil
	}

	logger.InfoContext(ctx, "workflow run complete")
	if err = s.controller.SetWorkflowRunStatus(ctx, sw.runID, workflow.StatusComplete, output); err != nil {
		return err
	}

	return nil
}

func (s *Scheduler) runTask(ctx context.Context, st *scheduledTask) error {
	if err := s.controller.SetTaskRunStatus(ctx, st.runID, task.StatusRunning, nil); err != nil {
		return err
	}

	logger := log.FromContext(ctx).With(
		slog.String("task_run_id", st.runID),
		slog.String("task_name", st.task.Name()),
	)

	logger.InfoContext(ctx, "running task")
	output, err := st.task.Run(ctx, st.runID, st.input)
	if err != nil {
		var e task.Error
		if !errors.As(err, &e) {
			e = task.Error{
				Message:  err.Error(),
				TaskName: st.task.Name(),
				RunID:    st.runID,
			}
		}

		output, err = json.Marshal(e)
		if err != nil {
			return err
		}

		logger.With(slog.String("error", e.Message)).ErrorContext(ctx, "task run failed")
		if err = s.controller.SetTaskRunStatus(ctx, st.runID, task.StatusFailed, output); err != nil {
			return err
		}

		return nil
	}

	logger.InfoContext(ctx, "task run complete")
	if err = s.controller.SetTaskRunStatus(ctx, st.runID, task.StatusComplete, output); err != nil {
		return err
	}

	return nil
}
