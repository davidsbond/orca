package workflow

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	"github.com/davidsbond/orca/internal/log"
	"github.com/davidsbond/orca/internal/task"
	"github.com/davidsbond/orca/internal/workflow"
)

type (
	Implementation[Input, Output any] struct {
		WorkflowName string
		Options      Options
		Action       Action[Input, Output]
		KeyFunc      KeyFunc[Input]
	}

	KeyFunc[T any] func(T) string

	Action[Input any, Output any] func(ctx context.Context, input Input) (Output, error)

	Options struct {
		Timeout time.Duration
	}
)

func (w *Implementation[Input, Output]) Run(ctx context.Context, runID string, input json.RawMessage) (output json.RawMessage, err error) {
	var inp Input

	if len(input) > 0 {
		if err = json.Unmarshal(input, &inp); err != nil {
			return nil, err
		}
	}

	defer func() {
		r := recover()
		if r == nil {
			return
		}

		var message string
		switch e := r.(type) {
		case error:
			message = e.Error()
		case string:
			message = e
		}

		log.FromContext(ctx).With(
			slog.String("workflow_name", w.WorkflowName),
			slog.String("workflow_run_id", runID),
			slog.String("error", message),
		).ErrorContext(ctx, "panic during task execution")

		err = workflow.Error{
			Message:      message,
			WorkflowName: w.WorkflowName,
			RunID:        runID,
			Panic:        true,
		}
	}()

	out, err := w.Action(ctx, inp)
	if err != nil {
		workflowErr := workflow.Error{
			Message:      err.Error(),
			WorkflowName: w.WorkflowName,
			RunID:        runID,
		}

		// If we're getting an error from a child workflow, do some mappings here to preserve the original
		// error message and whether it was a panic.
		var we workflow.Error
		if errors.As(err, &we) {
			workflowErr.Message = we.Message
			workflowErr.Panic = we.Panic
		}

		// If we're getting an error from a task this workflow runs, we also do some mappings so that we have
		// information on the task that failed and if that was a panic.
		var te task.Error
		if errors.As(err, &te) {
			we.Message = te.Message
			we.TaskRunID = te.RunID
			we.TaskName = te.TaskName
			we.Panic = te.Panic
		}

		return nil, workflowErr
	}

	output, err = json.Marshal(out)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (w *Implementation[Input, Output]) Name() string {
	return w.WorkflowName
}
