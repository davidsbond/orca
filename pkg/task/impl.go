package task

import (
	"context"
	"encoding/json"

	"github.com/davidsbond/orca/internal/task"
)

type (
	Implementation[Input, Output any] struct {
		TaskName string
		Action   Action[Input, Output]
	}

	Action[Input any, Output any] func(ctx context.Context, input Input) (Output, error)
)

func (t *Implementation[Input, Output]) Run(ctx context.Context, runID string, input json.RawMessage) (json.RawMessage, error) {
	var inp Input
	if len(input) > 0 {
		if err := json.Unmarshal(input, &inp); err != nil {
			return nil, err
		}
	}

	output, err := t.Action(ctx, inp)
	if err != nil {
		return nil, task.Error{
			Message:  err.Error(),
			TaskName: t.TaskName,
			RunID:    runID,
		}
	}

	return json.Marshal(output)
}

func (t *Implementation[Input, Output]) Name() string {
	return t.TaskName
}
