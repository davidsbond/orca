package orca

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/davidsbond/orca/internal/log"
	"github.com/davidsbond/orca/internal/task"
	"github.com/davidsbond/orca/internal/workflow"
)

type (
	Workflow[Input, Output any] struct {
		WorkflowName string
		Options      ActionOptions
		Action       Action[Input, Output]
		KeyFunc      KeyFunc[Input]
	}
)

func (w *Workflow[Input, Output]) Run(ctx context.Context, runID string, input json.RawMessage) (output json.RawMessage, err error) {
	var inp Input

	if len(input) > 0 {
		if err = json.Unmarshal(input, &inp); err != nil {
			return nil, err
		}
	}

	logger := log.FromContext(ctx).With(
		slog.String("workflow_name", w.WorkflowName),
		slog.String("workflow_run_id", runID),
	)

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

		logger.With(slog.String("error", message)).ErrorContext(ctx, "panic during task execution")

		err = workflow.Error{
			Message:      message,
			WorkflowName: w.WorkflowName,
			RunID:        runID,
			Panic:        true,
		}
	}()

	wCtx := ctx
	var cancel context.CancelFunc

	// If the workflow has specified a timeout, we modify the special context passed
	// through the workflow action to use it.
	if w.Options.Timeout > 0 {
		wCtx, cancel = context.WithTimeout(ctx, w.Options.Timeout)
		defer cancel()
	}

	var out Output

	// Perform the action up to the retry count, if the error given by the action
	// is nil we break from the loop.
	for range w.Options.RetryCount + 1 {
		out, err = w.Action(wCtx, inp)
		if err == nil {
			break
		}
	}

	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return nil, workflow.ErrTimeout
	case err != nil:
		return nil, workflowError(w.WorkflowName, runID, err)
	}

	output, err = json.Marshal(out)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (w *Workflow[Input, Output]) Name() string {
	return w.WorkflowName
}

func workflowError(name, runID string, err error) error {
	workflowErr := workflow.Error{
		Message:      err.Error(),
		WorkflowName: name,
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

	return workflowErr
}
