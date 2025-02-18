package orca

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/davidsbond/orca/internal/log"
	"github.com/davidsbond/orca/internal/task"
)

type (
	Task[Input, Output any] struct {
		TaskName string
		Action   Action[Input, Output]
		KeyFunc  KeyFunc[Input]
		Options  ActionOptions
	}
)

func (t *Task[Input, Output]) Run(ctx context.Context, runID string, input json.RawMessage) (output json.RawMessage, err error) {
	var inp Input
	if len(input) > 0 {
		if err = json.Unmarshal(input, &inp); err != nil {
			return nil, err
		}
	}

	logger := log.FromContext(ctx).With(
		slog.String("task_name", t.TaskName),
		slog.String("task_run_id", runID),
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

		err = task.Error{
			Message:  message,
			TaskName: t.TaskName,
			RunID:    runID,
			Panic:    true,
		}
	}()

	tCtx := ctx
	var cancel context.CancelFunc

	// If the task has specified a timeout, we modify the special context passed
	// through the task action to use it.
	if t.Options.Timeout > 0 {
		tCtx, cancel = context.WithTimeout(ctx, t.Options.Timeout)
		defer cancel()
	}

	var out Output

	// Perform the action up to the retry count, if the error given by the action
	// is nil we break from the loop.
	for range t.Options.RetryCount + 1 {
		out, err = t.Action(tCtx, inp)
		if err == nil {
			break
		}
	}

	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return nil, task.ErrTimeout
	case err != nil:
		return nil, task.Error{
			Message:  err.Error(),
			TaskName: t.TaskName,
			RunID:    runID,
		}
	}

	output, err = json.Marshal(out)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (t *Task[Input, Output]) Name() string {
	return t.TaskName
}
