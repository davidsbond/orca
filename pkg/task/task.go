package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/davidsbond/orca/internal/daemon/worker/task"
	"github.com/davidsbond/orca/internal/daemon/worker/workflow"
)

type (
	Implementation[Input, Output any] struct {
		TaskName string
		Action   Action[Input, Output]
	}

	Action[Input any, Output any] func(ctx context.Context, input Input) (Output, error)
)

func (t *Implementation[Input, Output]) Run(ctx context.Context, input []byte) ([]byte, error) {
	var inp Input
	if err := json.Unmarshal(input, &input); err != nil {
		return nil, err
	}

	output, err := t.Action(ctx, inp)
	if err != nil {
		return nil, err
	}

	return json.Marshal(output)
}

func (t *Implementation[Input, Output]) Name() string {
	return t.TaskName
}

func Execute[Output any](ctx context.Context, t task.Task, input any) (Output, error) {
	var output Output

	client, ok := task.ClientFromContext(ctx)
	if !ok {
		return output, errors.New("client not present in context")
	}

	params, err := json.Marshal(input)
	if err != nil {
		return output, fmt.Errorf("failed to marshal input: %w", err)
	}

	run, ok := workflow.RunFromContext(ctx)
	if !ok {
		return output, errors.New("workflow not present in context")
	}

	runID, err := client.ScheduleTask(ctx, run.ID, t.Name(), params)
	if err != nil {
		return output, fmt.Errorf("could not schedule task: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return output, ctx.Err()
		default:
			status, result, err := client.GetTaskStatus(ctx, runID)
			if err != nil {
				return output, err
			}

			switch status {
			case task.StatusRunning, task.StatusPending, task.StatusUnspecified:
				continue
			case task.StatusComplete:
				if err = json.Unmarshal(result, &output); err != nil {
					return output, fmt.Errorf("failed to unmarshal output: %w", err)
				}

				return output, nil
			case task.StatusFailed:
				return output, errors.New(string(result))
			default:
				continue
			}
		}
	}
}
