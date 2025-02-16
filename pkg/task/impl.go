package task

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/davidsbond/orca/internal/log"
	"github.com/davidsbond/orca/internal/task"
)

type (
	Implementation[Input, Output any] struct {
		TaskName string
		Action   Action[Input, Output]
	}

	Action[Input any, Output any] func(ctx context.Context, input Input) (Output, error)
)

func (t *Implementation[Input, Output]) Run(ctx context.Context, runID string, input json.RawMessage) (output json.RawMessage, err error) {
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
			slog.String("task_name", t.TaskName),
			slog.String("task_run_id", runID),
			slog.String("error", message),
		).ErrorContext(ctx, "panic during task execution")

		err = task.Error{
			Message:  message,
			TaskName: t.TaskName,
			RunID:    runID,
			Panic:    true,
		}
	}()

	out, err := t.Action(ctx, inp)
	if err != nil {
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

func (t *Implementation[Input, Output]) Name() string {
	return t.TaskName
}
